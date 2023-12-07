package main

import (
	"net/http"
	"os"

	"github.com/ksysoev/spam47/pkg/app"
)

func main() {
	port := os.Getenv("SPAM47_PORT")

	if port == "" {
		port = "80"
	}

	server := app.New()

	http.ListenAndServe(":"+port, server.Mux())
}
