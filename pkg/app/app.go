package app

import (
	"log/slog"
	"net/http"

	"github.com/ksysoev/spam47/pkg/aggr"
	"github.com/ksysoev/spam47/pkg/repo"
)

const (
	DEFAULT_PORT                = "80"
	DEFAULT_CLASSIFIER_FILEPATH = "/tmp/classifier.gob"
)

type App struct {
	engine *aggr.SpamEngine
}

func New() *App {
	return &App{}
}

func (a *App) Run() error {
	listen := ":" + DEFAULT_PORT

	slog.Info("Starting app server on " + listen)

	repo := repo.NewClassifierFileRepo(DEFAULT_CLASSIFIER_FILEPATH)

	engine, err := aggr.NewSpamEngine(repo)

	if err != nil {
		return err
	}

	a.engine = engine

	return http.ListenAndServe(listen, a.Mux())
}

func (a *App) Mux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/livez", a.HeathCheck)
	mux.HandleFunc("/check", a.Check)
	mux.HandleFunc("/train", a.Train)

	return mux
}
