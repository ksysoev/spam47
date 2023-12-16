package app

import (
	"log/slog"
	"net/http"

	"github.com/jbrukh/bayesian"
	"github.com/ksysoev/spam47/pkg/repo"
)

const (
	DEFAULT_PORT = "80"
)

const (
	Spam bayesian.Class = "Spam"
	Ham  bayesian.Class = "Ham"
)

type App struct {
	engine     *bayesian.Classifier
	engineRepo repo.EngineRepo
}

func New() *App {
	return &App{
		engine: bayesian.NewClassifier(Spam, Ham),
	}
}

func (a *App) Run() error {
	listen := ":" + DEFAULT_PORT

	slog.Info("Starting app server on " + listen)

	repo, err := repo.NewEngineRedisRepo()
	if err != nil {
		return err
	}

	a.engineRepo = repo

	return http.ListenAndServe(listen, a.Mux())
}

func (a *App) Mux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/livez", a.HeathCheck)
	mux.HandleFunc("/check", a.Check)
	mux.HandleFunc("/train", a.Train)

	return mux
}
