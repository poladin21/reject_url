package main

import (
	"log/slog"
	"os"
	"redirect_url/internal/config"
	"redirect_url/internal/servises"
)

func main() {

	// config inst

	cfg := config.New()

	// loger inst логер для логирования инфы чекни в консоль там данные можешь посмотреть ошибки нужно это для просмотра всего вышеперечисленного на сервакее
	loger := newLog(cfg.Loggers.Level)

	// storeg(база данных) inst (дб постгер sql)

	// servis inst

	servises := servises.New(nil, loger)

	// interfece (во всех ниже перечисленных)

	// hendler init

	// start servis

}

func newLog(level string) *slog.Logger {
	var lvl slog.Level
	switch level {
	case "debug":
		lvl = slog.LevelDebug
	default:
		lvl = slog.LevelInfo
	}

	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: lvl}),
	)
}
