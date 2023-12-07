package app

import "net/http"

func (a *App) HeathCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("{\"status\": \"OK\"}"))
}

func (a *App) CheckForSpam(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}

func (a *App) MarkAsSpam(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}

func (a *App) MarkAsHam(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}
