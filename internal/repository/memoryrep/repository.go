package memoryrep

import "errors"

var (
	ErrNoAvailableRooms = errors.New("all rooms are occupied at this time")
)
