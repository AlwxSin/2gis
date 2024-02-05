package models

import (
	"context"
	"time"
)

// RoomAvailability represents room availability for specific hotel on specific date.
type RoomAvailability struct {
	HotelID string    `json:"hotelID"`
	RoomID  string    `json:"roomID"`
	Date    time.Time `json:"date"`
	Quota   int       `json:"quota"`
}

func GetRoomAvailability(ctx context.Context, db DB, hotelID, roomID string, from, to time.Time) ([]*RoomAvailability, error) {
	return db.GetRoomAvailability(ctx, hotelID, roomID, from, to)
}
