package booking

import (
	"applicationDesignTest/internal/model"
	"context"
	"testing"
)

func TestRepository_CancelReservation(t *testing.T) {

	rep := NewRepository()

	type args struct {
		ctx           context.Context
		reservedRooms model.ReservedRoomsIDs
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				ctx: nil,
				reservedRooms: model.ReservedRoomsIDs{
					Quota: 1,
					IDs:   []int{1, 2, 3},
				},
			},
			wantErr: false,
		},
		{
			name: "Positive",
			args: args{
				ctx: nil,
				reservedRooms: model.ReservedRoomsIDs{
					Quota: 1,
					IDs:   []int{1, 2, 3},
				},
			},
			wantErr: false,
		},
		{
			name: "[]int{} mass",
			args: args{
				ctx: nil,
				reservedRooms: model.ReservedRoomsIDs{
					Quota: 1,
					IDs:   []int{},
				},
			},
			wantErr: false,
		},
		{
			name: "nil mass",
			args: args{
				ctx: nil,
				reservedRooms: model.ReservedRoomsIDs{
					Quota: 1,
					IDs:   nil,
				},
			},
			wantErr: false,
		},
		{
			name: "our of range mass",
			args: args{
				ctx: nil,
				reservedRooms: model.ReservedRoomsIDs{
					Quota: 1,
					IDs:   []int{1, 2, 3, 999},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rep.CancelReservation(tt.args.ctx, tt.args.reservedRooms); (err != nil) != tt.wantErr {
				t.Errorf("CancelReservation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
