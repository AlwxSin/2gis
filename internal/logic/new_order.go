package logic

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"applicationDesignTest/internal"
	"applicationDesignTest/internal/models"
)

// NewOrderRequest represents all necessary fields to create an order.
type NewOrderRequest struct {
	HotelID   string    `json:"hotelID"`
	RoomID    string    `json:"roomID"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (no *NewOrderRequest) Validate() error {
	if no.HotelID == "" {
		return errors.New("missing required hotel_id field")
	}
	if no.RoomID == "" {
		return errors.New("missing required room_id field")
	}
	if no.UserEmail == "" {
		return errors.New("missing required email field")
	}

	if no.From.After(no.To) {
		return errors.New("from field can't be greater than to field")
	}

	return nil
}

func (no *NewOrderRequest) Bind(_ *http.Request) error {
	return no.Validate()
}

type NewOrderResponse struct {
	*models.Order
}

func (rd *NewOrderResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func CreateOrder(ctx context.Context, db models.DB, newOrder *NewOrderRequest) (*NewOrderResponse, error) {
	availability, err := db.GetRoomAvailability(ctx, newOrder.HotelID, newOrder.RoomID, newOrder.From, newOrder.To) // should be select for update
	if err != nil {
		return nil, err
	}

	daysToBook := internal.DaysBetween(newOrder.From, newOrder.To)

	availableDays := make(map[time.Time]struct{})
	for _, av := range availability {
		availableDays[av.Date] = struct{}{}
	}

	unavailableDays := make(map[time.Time]struct{})

	for _, dayToBook := range daysToBook {
		if _, ok := availableDays[dayToBook]; !ok {
			unavailableDays[dayToBook] = struct{}{}
		}
	}

	if len(unavailableDays) != 0 {
		slog.Error("Hotel room is not available for selected dates", "order", newOrder, "unavailable_days", unavailableDays)
		return nil, errors.New("hotel room is not available for selected dates")
	}

	order, err := models.CreateOrder(ctx, db, &models.CreateOrderOptions{
		HotelID:   newOrder.HotelID,
		RoomID:    newOrder.RoomID,
		UserEmail: newOrder.UserEmail,
		From:      newOrder.From,
		To:        newOrder.To,
	})
	if err != nil {
		return nil, err
	}

	slog.Info("Order successfully created", "order", newOrder)

	return &NewOrderResponse{Order: order}, nil
}
