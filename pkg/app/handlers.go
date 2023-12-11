package app

import "net/http"

func (a *App) HeathCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("{\"status\": \"OK\"}"))
}

func (a *App) Check(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("HAM"))
}

func (a *App) Train(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}
