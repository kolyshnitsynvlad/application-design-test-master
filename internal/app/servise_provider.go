package app

import (
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/internal/model"
	repo "applicationDesignTest/internal/repository/memoryrep/booking"
	service "applicationDesignTest/internal/service/booking"
	"context"
)

type BookingRepository interface {
	CreateOrder(ctx context.Context, newOrder model.Order) error
	RoomReservation(ctx context.Context, newOrder model.Order) (model.ReservedRoomsIDs, error)
	CancelReservation(ctx context.Context, reservedRooms model.ReservedRoomsIDs) error
}

type BookingService interface {
	Create(ctx context.Context, order model.Order) error
}

type serviceProvider struct {
	bookingRepository BookingRepository
	bookingService    BookingService
	log               *logger.Logger
}

func newServiceProvider(log *logger.Logger) *serviceProvider {
	return &serviceProvider{
		log: log,
	}
}

func (s *serviceProvider) UserRepository() BookingRepository {
	if s.bookingRepository == nil {
		s.bookingRepository = repo.NewRepository()
	}

	return s.bookingRepository
}

func (s *serviceProvider) BookingService() BookingService {
	if s.bookingService == nil {
		s.bookingService = service.NewService(s.UserRepository(), s.log)
	}

	return s.bookingService
}
