package booking

import (
	"applicationDesignTest/internal/model"
	"applicationDesignTest/internal/repository/memoryrep"
	"context"
	"time"
)

func (r *Repository) RoomReservation(_ context.Context, newOrder model.Order) (model.ReservedRoomsIDs, error) {

	daysToBook := daysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{}, len(daysToBook))
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}
	availabilityIndToReduceQuota := make([]int, 0, len(daysToBook))

	r.AvailabilityRep.Lock()
	defer r.AvailabilityRep.Unlock()

	for i, availability := range r.AvailabilityRep.Availability {
		if _, ok := unavailableDays[availability.Date]; ok && availability.Quota > 0 {
			delete(unavailableDays, availability.Date)
			availabilityIndToReduceQuota = append(availabilityIndToReduceQuota, i)
		}
	}
	var reservedRooms model.ReservedRoomsIDs

	if len(unavailableDays) != 0 {
		return reservedRooms, memoryrep.ErrNoAvailableRooms
	}

	for _, ind := range availabilityIndToReduceQuota {
		r.AvailabilityRep.Availability[ind].Quota--
	}
	reservedRooms.IDs = availabilityIndToReduceQuota
	reservedRooms.Quota = 1

	return reservedRooms, nil
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
