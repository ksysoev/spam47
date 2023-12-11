package main

import (
	"log/slog"
	"os"

	"github.com/ksysoev/spam47/pkg/app"
)

func main() {
	server := app.New()
	err := server.Run()

	if err != nil {
		slog.Error("Fail to start app server", "error", err)
		os.Exit(1)
	}

	os.Exit(0)
}
