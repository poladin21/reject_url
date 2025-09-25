package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"redirect_url/internal/config"
	"redirect_url/internal/handlers"
	"redirect_url/internal/servises"
	"redirect_url/internal/storegs/url"

	_ "github.com/lib/pq"
)

func main() {

	// config inst
	cfg := config.New()

	// loger inst логер для логирования инфы чекни в консоль там данные можешь посмотреть ошибки нужно это для просмотра всего вышеперечисленного на сервакее
	loger := newLog(cfg.Loggers.Level)

	// storeg(база данных) inst (дб постгер sql)
	db, err := NewDB(*cfg)
	if err != nil {
		loger.Error("fatl inet db")
		return
	}

	storegs := url.New(db)

	// servis inst
	servises := servises.New(nil, loger)

	// interfece (во всех ниже перечисленных)

	// hendler init

	// Регистрируем обработчик для всех запросов
	h := handlers.NeHandlers()

	http.HandleFunc("/", h.Get)
	http.HandleFunc("/exit/", h.Post)

	// start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
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

func NewDB(cfg config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name,
	)
	return sql.Open("postgres", psqlInfo)

}
