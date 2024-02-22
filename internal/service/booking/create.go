package booking

import (
	"applicationDesignTest/internal/model"
	"applicationDesignTest/internal/service"
	"context"
	"time"
)

func (s *Service) Create(ctx context.Context, order model.Order) error {

	//TODO Validation

	reservedRooms, err := s.bookingRepository.RoomReservation(ctx, order)
	if err != nil {
		s.log.LogErrorf("booking repository return error: %v", err)
		return err
	}

	// payment simulation
	payment := paymentSimulation(false)

	if !payment {
		err = s.bookingRepository.CancelReservation(ctx, reservedRooms)
		if err != nil {
			s.log.LogErrorf("failed to cancel reservation, err: %v", err)
		}
		return service.ErrPaymentFailed
	}

	err = s.bookingRepository.CreateOrder(ctx, order)
	if err != nil {
		s.log.LogErrorf("booking repository return error: %v", err)
		return err
	}

	return nil
}

func paymentSimulation(ans bool) bool {
	time.Sleep(time.Second)
	return ans
}
