package http_server

import (
	"applicationDesignTest/internal/http-server/handlers/orders/create"
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/internal/model"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type BookingService interface {
	Create(ctx context.Context, order model.Order) error
}

func NewRouter(log *logger.Logger, service BookingService) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Post("/orders", create.New(log, service))

	return router
}
