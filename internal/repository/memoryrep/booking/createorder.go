package booking

import (
	"applicationDesignTest/internal/model"
	"context"
)

func (r *Repository) CreateOrder(_ context.Context, newOrder model.Order) error {

	r.OrdersRep.Lock()
	defer r.OrdersRep.Unlock()
	r.OrdersRep.Orders = append(r.OrdersRep.Orders, newOrder)

	return nil
}
