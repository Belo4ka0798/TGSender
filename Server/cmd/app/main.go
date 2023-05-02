package main

import (
	"context"
	"tgsender/config"
	app2 "tgsender/internal/app"
	"tgsender/internal/question/repo"
	"tgsender/internal/server"
	"tgsender/pkg/logging"
	"tgsender/pkg/storage"
)

func main() {
	ctx := context.Background()

	logger := logging.GetLogger()

	config.GetConfigs()

	srv := server.NewServer()

	stClient, err := storage.NewStorage(ctx)
	if err != nil {
		logger.Fatalf("can't init storage client, err: %v", err)
	}

	rp := repo.NewRepo(stClient, logger)

	app := app2.NewApp(logger, srv, rp)

	err = app.Init()
	if err != nil {
		logger.Fatalf("can't init app: %v", err)
	}
}
