package app

import (
	"log/slog"
	"net/http"
)

const (
	DEFAULT_PORT = "80"
)

type App struct{}

func New() *App {
	return &App{}
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
