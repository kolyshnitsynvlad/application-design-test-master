package booking

import (
	"applicationDesignTest/internal/model"
	"sync"
	"time"
)

type Repository struct {
	OrdersRep       OrdersRep
	AvailabilityRep AvailabilityRep
}

type OrdersRep struct {
	sync.Mutex
	Orders []model.Order
}

type AvailabilityRep struct {
	sync.Mutex
	Availability []model.RoomAvailability
}

func NewRepository() *Repository {
	return &Repository{
		OrdersRep: OrdersRep{
			Orders: make([]model.Order, 0, 10),
		},
		AvailabilityRep: AvailabilityRep{
			Availability: getAvailability(),
		},
	}
}

func getAvailability() []model.RoomAvailability {
	return []model.RoomAvailability{
		{"reddison", "lux", date(2024, 3, 1), 1},
		{"reddison", "lux", date(2024, 3, 2), 1},
		{"reddison", "lux", date(2024, 3, 3), 1},
		{"reddison", "lux", date(2024, 3, 4), 1},
		{"reddison", "lux", date(2024, 3, 5), 0},
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
