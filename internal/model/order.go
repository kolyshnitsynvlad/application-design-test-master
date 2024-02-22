package model

import (
	"errors"
	"time"
)

var (
	ErrIncorrectOrderData = errors.New("order data is incorrect")
)

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (o *Order) Validation() error {
	if len(o.HotelID) == 0 {
		return ErrIncorrectOrderData
	}
	if len(o.RoomID) == 0 {
		return ErrIncorrectOrderData
	}
	if len(o.UserEmail) == 0 {
		return ErrIncorrectOrderData
	}
	if o.From.Before(time.Now()) {
		return ErrIncorrectOrderData
	}
	if o.From.After(o.To) {
		return ErrIncorrectOrderData
	}
	return nil
}
