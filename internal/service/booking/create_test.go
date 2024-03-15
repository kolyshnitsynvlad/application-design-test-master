package booking

import (
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/internal/model"
	mock_booking "applicationDesignTest/internal/service/booking/mock"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestService_Booking_Create(t *testing.T) {

	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_booking.NewMockBookingRepository(ctl)
	ctx := context.Background()

	service := NewService(repo, logger.New())

	tests := []struct {
		name  string
		order model.Order
		err   error
	}{
		{
			name: "Positive",
			order: model.Order{
				HotelID:   "reddison",
				RoomID:    "lux",
				UserEmail: "guest@mail.ru",
				From:      time.Now().AddDate(0, 0, 3),
				To:        time.Now().AddDate(0, 0, 5),
			},
			err: nil,
		},
	}
	roomsIDs := model.ReservedRoomsIDs{
		Quota: 2,
		IDs:   []int{1, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo.EXPECT().RoomReservation(ctx, tt.order).Return(roomsIDs, nil).Times(1)
			//repo.EXPECT().CancelReservation(ctx, roomsIDs).Return(nil)
			repo.EXPECT().CreateOrder(ctx, tt.order).Return(nil).Times(1)

			err := service.Create(ctx, tt.order)
			assert.NoError(t, err)
		})
	}

}
