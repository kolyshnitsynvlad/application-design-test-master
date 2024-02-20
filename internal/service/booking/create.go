package booking

import (
	"applicationDesignTest/internal/model"
	"context"
)

func (s *service) Create(ctx context.Context, order model.Order) error {

	// validation

	//Repository
	err := s.bookingRepository.CreateOrder(ctx, order)
	if err != nil {
		s.log.LogErrorf("booking repository return error: %w", err)
		return err
	}

	return nil
}
