package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ksysoev/spam47/pkg/aggr"
	"github.com/ksysoev/spam47/pkg/repo"
)

const (
	DEFAULT_PORT                = "80"
	DEFAULT_CLASSIFIER_FILEPATH = "/tmp/spam47.gob"
)

type App struct {
	engine   *aggr.SpamEngine
	port     string
	datafile string
}

func New() *App {

	datafile := os.Getenv("SPAM47_DATAFILE")

	if datafile == "" {
		datafile = DEFAULT_CLASSIFIER_FILEPATH
	}

	port := os.Getenv("SPAM47_PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	return &App{
		datafile: datafile,
		port:     port,
	}
}

func (a *App) Run() error {
	listen := ":" + a.port

	slog.Info("Starting app server on " + listen)

	repo := repo.NewClassifierFileRepo(a.datafile)

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
