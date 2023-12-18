package repo

import (
	"context"
	"time"

	"github.com/redis/rueidis"
)

type EngineRepo interface{}

type EngineRedisRepo struct {
	redis rueidis.Client
}

func NewEngineRedisRepo() (*EngineRedisRepo, error) {
	redis, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})

	if err != nil {
		return nil, err
	}

	return &EngineRedisRepo{
		redis: redis,
	}, nil
}

func (e *EngineRedisRepo) Save(class string, data []string) error {
	cmd := e.redis.B().Sadd().Key(class).Member(data...).Build()
	e.redis.Do(context.Background(), cmd)
	return nil
}

func (e *EngineRedisRepo) Load(class string) (data []string, err error) {
	// This provides Redis client-side caching support.
	// If data is updated in redis, cache will be invalidated.
	cmd := e.redis.B().Smembers().Key(class).Cache()
	resp := e.redis.DoCache(context.Background(), cmd, time.Hour)

	if err := resp.Error(); err != nil {
		return nil, err
	}

	return resp.AsStrSlice()
}
