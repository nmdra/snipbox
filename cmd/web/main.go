package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type applog struct {
	Logger *slog.Logger
}

func main() {

	var port = flag.String("port", "4000", "HTTP network address")

	flag.Parse()

	// setup logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	log := &applog{
		Logger: logger,
	}

	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	envPort := os.Getenv("PORT") // for docker
	if envPort != "" {
		*port = envPort
		log.Logger.Info("Using environment variable PORT:" + *port)
	} else {
		log.Logger.Warn("Environment variable PORT is not set, using default or flag value")
	}

	log.Logger.Info("Starting server", slog.Any("PORT", *port))

	err := http.ListenAndServe(":"+*port, log.routes())
	log.Logger.Error(err.Error())
	os.Exit(1)
}
