package main

import (
	"log/slog"
	"os"

	"github.com/ll1bles404/flowers-delivery/internal/client"
	"github.com/ll1bles404/flowers-delivery/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	log := setupLogger("local")
	log.Info("Старт приложения!")

	cfg := config.MustLoad()
	log.Info("Конфиг загружен!")

	_ = cfg
	bot, _ := client.NewBot("f9LHodD0cOJ0LXycZ2pLFGw3NdFWqd74wCc3lbDjzx5Ypu1-FHAweKskFCq1MeTL6JDTePAc8-Uq5uyzDe90")
	bot.EventProcessor()
	// storage, err := pgsql.New(*cfg)

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
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	}

	return log
}
