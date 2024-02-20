package booking

import (
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/internal/model"
	"context"
)

type BookingRepository interface {
	CreateOrder(ctx context.Context, newOrder model.Order) error
}

type service struct {
	bookingRepository BookingRepository
	log               logger.CustomLogger
}

func NewService(bookingRepository BookingRepository, log logger.CustomLogger) *service {
	return &service{
		bookingRepository: bookingRepository,
		log:               log,
	}
}
