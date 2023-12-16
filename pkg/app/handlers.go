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

type CheckRequest struct {
	Message string `json:"message"`
	Lang    string `json:"lang,omitempty"`
}

type CheckResponse struct {
	Status string  `json:"status"`
	Score  float64 `json:"score"`
}

func (a *App) Check(w http.ResponseWriter, r *http.Request) {
	var req CheckRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg := stopwords.CleanString(req.Message, req.Lang, true)
	words := strings.Fields(msg)

	probs, indx, _ := a.engine.ProbScores(words)

	class := a.engine.Classes[indx]

	resp := CheckResponse{
		Status: string(class),
		Score:  probs[indx],
	}

	respJSON, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(respJSON)
}

type TrainRequest struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Lang    string `json:"lang,omitempty"`
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
