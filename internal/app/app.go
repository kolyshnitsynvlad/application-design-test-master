package app

import (
	booking_server "applicationDesignTest/internal/booking-server"
	"applicationDesignTest/internal/config"
	"applicationDesignTest/internal/logger"
	"context"
	"net/http"
)

type App struct {
	httpServer *http.Server
}

func NewApp(cfg config.Config, log *logger.Logger) *App {
	a := &App{}
	a.initHTTPServer(cfg, log)
	return a
}

func (a *App) Run() error {
	return a.httpServer.ListenAndServe()
}
func (a *App) Shutdown(ctx context.Context) error {
	return a.httpServer.Shutdown(ctx)
}

func (a *App) initHTTPServer(cfg config.Config, log *logger.Logger) {
	a.httpServer = &http.Server{
		Addr:         cfg.Address,
		Handler:      booking_server.NewRouter(log),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
