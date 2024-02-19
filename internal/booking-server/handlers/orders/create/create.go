package create

import (
	"applicationDesignTest/internal/logger"
	"applicationDesignTest/internal/model"
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"time"
)

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type Request struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

// TODO create general structure
type Response struct {
	Status    string    `json:"status"`
	Error     string    `json:"error,omitempty"`
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

type OrderBooker interface {
	BookOrder(order model.Order) error
}

func New(log *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.LogInfo("Start orders handler")

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом.
			log.LogErrorf("request body is empty")
			render.JSON(w, r, Response{
				Status: StatusError,
				Error:  "request body is empty",
			})
			return
		}
		if err != nil {
			log.LogErrorf("failed to decode request body: %v", err)
			render.JSON(w, r, Response{
				Status: StatusError,
				Error:  "failed to decode request",
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, Response{
			Status:    StatusOK,
			HotelID:   req.HotelID,
			RoomID:    req.RoomID,
			UserEmail: req.UserEmail,
			From:      req.From,
			To:        req.To,
		})

		log.LogInfo("Order successfully created: %v", req)
	}
}
