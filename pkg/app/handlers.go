package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bbalet/stopwords"
)

func (a *App) HeathCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("{\"status\": \"OK\"}"))
}

func (a *App) Check(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("HAM"))
}

type TrainRequest struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Lang    string `json:"lang"`
}

func (a *App) Train(w http.ResponseWriter, r *http.Request) {
	var req TrainRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg := stopwords.CleanString(req.Message, req.Lang, true)
	words := strings.Fields(msg)

	switch req.Type {
	case "spam":
		a.engine.Learn(words, Spam)
	case "ham":
		a.engine.Learn(words, Ham)
	default:
		http.Error(w, "invalid type", http.StatusBadRequest)
		return
	}

	w.Write([]byte("{\"status\": \"OK\"}"))
}
