package booking

import (
	"applicationDesignTest/internal/model"
	"sync"
	"time"
)

type repository struct {
	Orders       []model.Order
	Availability []model.RoomAvailability
	rwm          sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		Orders:       make([]model.Order, 0, 10),
		Availability: getAvailability(),
	}
}

func getAvailability() []model.RoomAvailability {
	return []model.RoomAvailability{
		{"reddison", "lux", date(2024, 1, 1), 1},
		{"reddison", "lux", date(2024, 1, 2), 1},
		{"reddison", "lux", date(2024, 1, 3), 1},
		{"reddison", "lux", date(2024, 1, 4), 1},
		{"reddison", "lux", date(2024, 1, 5), 0},
	}
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
