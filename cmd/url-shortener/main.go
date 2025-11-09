package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"restApp/internal/config"
	sl "restApp/internal/lib/logger/slog"
	"restApp/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	fmt.Printf("CONFIG_PATH value: '%s'\n", configPath)

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Debug("debug messages are enabled", slog.String("env", cfg.Env))
	log.Info("starting url-shortener", slog.String("env", cfg.Env))

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// TODO: init storage: sqlite
	// TODO: init router: chi, "chi render"
	// TODO: run server:
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
