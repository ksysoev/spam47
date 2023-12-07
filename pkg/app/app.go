package app

import "net/http"

type App struct{}

func New() *App {
	return &App{}
}

func (a *App) Mux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/livez", a.HeathCheck)

	return mux
}
