package app

import (
	"log/slog"
	"net/http"

	"github.com/jbrukh/bayesian"
)

const (
	DEFAULT_PORT = "80"
)

const (
	Spam bayesian.Class = "Spam"
	Ham  bayesian.Class = "Ham"
)

type App struct {
	engine *bayesian.Classifier
}

func New() *App {

	return &App{
		engine: bayesian.NewClassifier(Spam, Ham),
	}
}

func (a *App) Run() error {
	listen := ":" + DEFAULT_PORT

	slog.Info("Starting app server on " + listen)
	return http.ListenAndServe(listen, a.Mux())
}

func (a *App) Mux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/livez", a.HeathCheck)
	mux.HandleFunc("/check", a.Check)
	mux.HandleFunc("/train", a.Train)

	return mux
}
