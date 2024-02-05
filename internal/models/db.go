package models

import (
	"context"
	"fmt"
	"sync"
	"time"

	"applicationDesignTest/internal"
)

type DB interface {
	GetRoomAvailability(ctx context.Context, hotelID, roomID string, from, to time.Time) ([]*RoomAvailability, error)
	markRoomsAvailable(ctx context.Context, hotelID, roomID string, from, to time.Time) error

	CreateOrder(ctx context.Context, hotelID, roomID, userEmail string, from, to time.Time) (*Order, error)
}

type InMemory struct {
	m            sync.Mutex
	orders       []*Order
	availability map[string]map[string]map[time.Time]int // hotel -> room_id -> date -> quota
}

func (i *InMemory) GetRoomAvailability(_ context.Context, hotelID, roomID string, from, to time.Time) ([]*RoomAvailability, error) {
	i.m.Lock()
	defer i.m.Unlock()

	daysToBook := internal.DaysBetween(from, to)

	hotelAv, ok := i.availability[hotelID]
	if !ok {
		return nil, fmt.Errorf("hotel %s not found", hotelID)
	}

	roomAv, ok := hotelAv[roomID]
	if !ok {
		return nil, fmt.Errorf("room %s in hotel %s not found", roomID, hotelID)
	}

	avs := make([]*RoomAvailability, 0, len(daysToBook))
	for _, dtb := range daysToBook {
		if dayQuota, ok := roomAv[dtb]; ok && dayQuota > 0 {
			avs = append(avs, &RoomAvailability{
				HotelID: hotelID,
				RoomID:  roomID,
				Date:    dtb,
				Quota:   dayQuota,
			})
		}
	}
	return avs, nil
}

func (i *InMemory) markRoomsAvailable(_ context.Context, hotelID, roomID string, from, to time.Time) error {
	daysToBook := internal.DaysBetween(from, to)

	hotelAv, ok := i.availability[hotelID]
	if !ok {
		return fmt.Errorf("hotel %s not found", hotelID)
	}

	roomAv, ok := hotelAv[roomID]
	if !ok {
		return fmt.Errorf("room %s in hotel %s not found", roomID, hotelID)
	}
	for _, dtb := range daysToBook {
		if dayQuota, ok := roomAv[dtb]; ok && dayQuota > 0 {
			roomAv[dtb] = 0
		}
	}
	return nil
}

func (i *InMemory) CreateOrder(ctx context.Context, hotelID, roomID, userEmail string, from, to time.Time) (*Order, error) {
	i.m.Lock()
	defer i.m.Unlock()

	err := i.markRoomsAvailable(ctx, hotelID, roomID, from, to)
	if err != nil {
		return nil, err
	}

	order := &Order{
		HotelID:   hotelID,
		RoomID:    roomID,
		UserEmail: userEmail,
		From:      from,
		To:        to,
	}
	i.orders = append(i.orders, order)

	return order, nil
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func NewInMemoryDB() *InMemory {
	return &InMemory{
		orders: make([]*Order, 0),
		availability: map[string]map[string]map[time.Time]int{
			"reddison": {
				"lux": map[time.Time]int{
					date(2024, 1, 1): 1,
					date(2024, 1, 2): 1,
					date(2024, 1, 3): 1,
					date(2024, 1, 4): 1,
					date(2024, 1, 5): 0,
					date(2025, 2, 1): 1,
				},
			},
		},
	}
}
