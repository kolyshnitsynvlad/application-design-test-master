package main

import (
	"applicationDesignTest/internal/app"
	"applicationDesignTest/internal/config"
	"applicationDesignTest/internal/logger"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg := config.MustLoad()

	log := logger.New()

	log.LogErrorf("int if work??")

	ctx := context.Background()

	ap := app.NewApp(cfg, log)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := ap.Run(); err != nil {
			log.LogErrorf("failed to start server")
		}
	}()
	log.LogInfo("server started")

	<-done

	log.LogInfo("stopping server")
	//timeout to config
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := ap.Shutdown(ctx); err != nil {
		log.LogErrorf("failed to stop server err: %v", err)
		return
	}
	log.LogInfo("server stopped")
}
