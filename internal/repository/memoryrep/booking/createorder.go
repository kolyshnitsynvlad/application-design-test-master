package booking

import (
	"applicationDesignTest/internal/model"
	"applicationDesignTest/internal/repository/memoryrep"
	"context"
	"time"
)

func (r *repository) CreateOrder(_ context.Context, newOrder model.Order) error {

	daysToBook := daysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}
	availabilityIndToReduceQuota := make([]int, 0, len(daysToBook))

	//maybe need use only Lock() (because maybe write error)
	//r.rwm.RLock()
	r.rwm.Lock()
	defer r.rwm.Unlock()

	for i, availability := range r.Availability {
		if _, ok := unavailableDays[availability.Date]; ok && availability.Quota > 0 {
			delete(unavailableDays, availability.Date)
			availabilityIndToReduceQuota = append(availabilityIndToReduceQuota, i)
		}
	}
	//r.rwm.RUnlock()

	if len(unavailableDays) != 0 {
		return memoryrep.ErrNoAvailableRooms
	}

	//r.rwm.Lock()
	//defer r.rwm.Unlock()

	for _, ind := range availabilityIndToReduceQuota {
		r.Availability[ind].Quota--
	}
	r.Orders = append(r.Orders, newOrder)

	return nil
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
