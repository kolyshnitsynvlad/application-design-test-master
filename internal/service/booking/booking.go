package booking

import (
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/internal/model"
	"context"
)

type BookingRepository interface {
	CreateOrder(ctx context.Context, newOrder model.Order) error
}

type Service struct {
	bookingRepository BookingRepository
	log               logger.CustomLogger
}

func NewService(bookingRepository BookingRepository, log logger.CustomLogger) *Service {
	return &Service{
		bookingRepository: bookingRepository,
		log:               log,
	}
}
