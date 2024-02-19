package service

import (
	"applicationDesignTest/internal/model"
	"context"
	"sync"
)

const defaultCapacity = 100

type (
	Writer interface {
		CreateOrder(order model.Order) error
	}
	Manager struct {
		writer Writer

		ctx    context.Context
		cancel context.CancelFunc

		mu    sync.Mutex
		order []model.Order
	}
)

func newServer(w Writer) *Manager {
	ctx, cancel := context.WithCancel(context.Background())

	return &Manager{
		writer: w,
		ctx:    ctx,
		cancel: cancel,
		order:  newOrders(),
	}
}

func newOrders() []model.Order {
	return make([]model.Order, 0, defaultCapacity)
}
