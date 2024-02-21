package booking

import (
	"applicationDesignTest/internal/model"
	"context"
)

func (s *Service) Create(ctx context.Context, order model.Order) error {

	//TODO Validation

	//Repository
	err := s.bookingRepository.CreateOrder(ctx, order)
	if err != nil {
		s.log.LogErrorf("booking repository return error: %v", err)
		return err
	}

	return nil
}
