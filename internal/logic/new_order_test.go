package logic

import (
	"context"
	"reflect"
	"testing"
	"time"

	"applicationDesignTest/internal/models"
)

func TestNewOrder(t *testing.T) {
	db := models.NewInMemoryDB() // for demo project purpose, we know dates in inmemory db
	order, err := CreateOrder(context.Background(), db, &NewOrderRequest{
		HotelID:   "reddison",
		RoomID:    "lux",
		UserEmail: "asinichkin@usetech.ru",
		From:      time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
		To:        time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Errorf("should not be an error on order create, %s", err)
	}

	want := &NewOrderResponse{Order: &models.Order{
		HotelID:   "reddison",
		RoomID:    "lux",
		UserEmail: "asinichkin@usetech.ru",
		From:      time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC),
		To:        time.Date(2024, time.January, 4, 0, 0, 0, 0, time.UTC),
	}}
	if !reflect.DeepEqual(order, want) {
		t.Errorf("NewOrder() = %v, want %v", order, want)
	}
}

func TestNewOrderUnavailableDates(t *testing.T) {
	db := models.NewInMemoryDB() // for demo project purpose, we know dates in inmemory db
	order, err := CreateOrder(context.Background(), db, &NewOrderRequest{
		HotelID:   "reddison",
		RoomID:    "lux",
		UserEmail: "asinichkin@usetech.ru",
		From:      time.Date(2024, time.January, 3, 0, 0, 0, 0, time.UTC),
		To:        time.Date(2024, time.January, 5, 0, 0, 0, 0, time.UTC),
	})

	if err == nil {
		t.Errorf("should be an error on order create - got %v, ", order)
	}
}
