package http_server

import (
	"applicationDesignTest/internal/app"
	"applicationDesignTest/internal/http-server/handlers/orders/create"
	"applicationDesignTest/internal/lib/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(log *logger.Logger, service app.BookingService) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Post("/orders", create.New(log, service))

	return router
}
