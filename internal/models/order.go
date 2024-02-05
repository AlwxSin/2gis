package models

import (
	"context"
	"time"
)

// Order represents room order for specific user on specific dates.
type Order struct {
	HotelID   string    `json:"hotelID"`
	RoomID    string    `json:"roomID"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

type CreateOrderOptions struct {
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}

func CreateOrder(ctx context.Context, db DB, opts *CreateOrderOptions) (*Order, error) {
	return db.CreateOrder(ctx, opts.HotelID, opts.RoomID, opts.UserEmail, opts.From, opts.To)
}
