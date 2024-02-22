package booking

import (
	"applicationDesignTest/internal/model"
	"context"
	"errors"
	"fmt"
)

func (r *Repository) CancelReservation(_ context.Context, reservedRooms model.ReservedRoomsIDs) error {

	r.AvailabilityRep.Lock()
	defer r.AvailabilityRep.Unlock()

	notCanceledRoomID := make([]int, 0)

	for _, ind := range reservedRooms.IDs {
		if ind < 0 || ind >= len(r.AvailabilityRep.Availability) {
			notCanceledRoomID = append(notCanceledRoomID, ind)
			continue
		}
		r.AvailabilityRep.Availability[ind].Quota -= reservedRooms.Quota
		if r.AvailabilityRep.Availability[ind].Quota < 0 {
			r.AvailabilityRep.Availability[ind].Quota = 0
		}
	}
	if len(notCanceledRoomID) > 0 {
		return errors.New(
			fmt.Sprintf("can't cansel reserved in rooms: %v, quota: %d",
				notCanceledRoomID, reservedRooms.Quota))
	}

	return nil
}
