package app

import "net/http"

func (a *App) HeathCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}
