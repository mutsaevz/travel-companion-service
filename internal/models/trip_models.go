package models

import (
	"time"
)

type Trip struct {
	Base

	DriverID       uint      `json:"driver_id"`
	CarID          uint      `json:"car_id"`
	From           string    `json:"from"`
	To             string    `json:"to"`
	StartTime      time.Time `json:"start_time"`
	DurationMin    int       `json:"duration_min"`
	TotalSeats     int       `json:"total_seats"`
	AvailableSeats int       `json:"available_seats"`
	Price          int       `json:"price"`
	TripStatus     string    `json:"trip_status"`
	AvgRating      float64   `json:"avg_rating"`
}

type TripCreateRequest struct {
	DriverID       uint      `json:"driver_id"`
	CarID          uint      `json:"car_id"`
	From           string    `json:"from"`
	To             string    `json:"to"`
	StartTime      time.Time `json:"start_time"`
	DurationMin    int       `json:"duration_min"`
	TotalSeats     int       `json:"total_seats"`
	AvailableSeats int       `json:"available_seats"`
	Price          int       `json:"price"`
	TripStatus     string    `json:"trip_status"`
}

type TripUpdateRequest struct {
	DriverID       uint      `json:"driver_id"`
	CarID          uint      `json:"car_id"`
	From           string    `json:"from"`
	To             string    `json:"to"`
	StartTime      time.Time `json:"start_time"`
	DurationMin    int       `json:"duration_min"`
	TotalSeats     int       `json:"total_seats"`
	AvailableSeats int       `json:"available_seats"`
	Price          int       `json:"price"`
	TripStatus     string    `json:"trip_status"`
}
