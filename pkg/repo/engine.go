package repo

import "github.com/redis/rueidis"

type EngineRepo interface{}

type EngineRedisRepo struct {
	redis any
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
