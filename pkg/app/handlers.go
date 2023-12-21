package app

import (
	"encoding/json"
	"net/http"
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

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, score := a.engine.Check(req.Message, req.Lang)

	resp := CheckResponse{
		Status: status,
		Score:  score,
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

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = a.engine.Train(req.Message, req.Type, req.Lang)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{\"status\": \"OK\"}"))
}
