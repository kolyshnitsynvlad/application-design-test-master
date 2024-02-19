package booking_server

import (
	"applicationDesignTest/internal/booking-server/handlers/orders/create"
	"applicationDesignTest/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(log *logger.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Post("/orders", create.New(log))

	return router
}
