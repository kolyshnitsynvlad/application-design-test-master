package app

import (
	"applicationDesignTest/internal/config"
	bookingserver "applicationDesignTest/internal/http-server"
	"applicationDesignTest/internal/lib/logger"
	"context"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(cfg config.Config, log *logger.Logger) *App {
	a := &App{}
	a.initDeps(cfg, log)
	return a
}

func (a *App) Run() error {
	return a.httpServer.ListenAndServe()
}
func (a *App) Shutdown(ctx context.Context) error {
	return a.httpServer.Shutdown(ctx)
}

func (a *App) initDeps(cfg config.Config, log *logger.Logger) {
	inits := []func(config.Config, *logger.Logger){
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		f(cfg, log)
	}
}
func (a *App) initServiceProvider(_ config.Config, log *logger.Logger) {
	a.serviceProvider = newServiceProvider(log)
}

// роутер можно вынести во входные параметры при создании
func (a *App) initHTTPServer(cfg config.Config, log *logger.Logger) {
	a.httpServer = &http.Server{
		Addr:         cfg.Address,
		Handler:      bookingserver.NewRouter(log, a.serviceProvider.BookingService()),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
