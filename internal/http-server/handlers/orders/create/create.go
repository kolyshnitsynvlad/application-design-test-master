package create

import (
	resp "applicationDesignTest/internal/lib/api/responce"
	"applicationDesignTest/internal/lib/logger"
	"applicationDesignTest/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"time"
)

type Request struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

type Response struct {
	resp.Response
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

type OrderCreator interface {
	Create(ctx context.Context, order model.Order) error
}

func New(log *logger.Logger, orderCreator OrderCreator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.LogInfo("Start order create request")

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом.
			log.LogErrorf("request body is empty")
			render.JSON(w, r, resp.Error("request body is empty"))
			return
		}
		if err != nil {
			log.LogErrorf("failed to decode request body: %v", err)
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}
		//TODO need content
		err = orderCreator.Create(context.Background(), req.convertReqDataToOrder())
		if err != nil {
			log.LogErrorf("the service returned an error: %v", err)
			w.WriteHeader(http.StatusInternalServerError) //TODO it is need ?
			render.JSON(w, r, resp.Error(fmt.Sprintf("can't create order, error: %v", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, Response{
			Response:  resp.OK(),
			HotelID:   req.HotelID,
			RoomID:    req.RoomID,
			UserEmail: req.UserEmail,
			From:      req.From,
			To:        req.To,
		})

		log.LogInfo("Order successfully created: %v", req)
	}
}

func (r *Request) convertReqDataToOrder() model.Order {
	return model.Order{
		HotelID:   r.HotelID,
		RoomID:    r.RoomID,
		UserEmail: r.UserEmail,
		From:      r.From,
		To:        r.To,
	}
}
