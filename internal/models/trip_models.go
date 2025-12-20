package models

import (
	"time"
)

type Trip struct {
	Base

	DriverID       uint
	CarID          uint
	From           string
	To             string
	StartTime      time.Time
	DurationMin    int
	TotalSeats     int
	AvailableSeats int
	Price          int
	TripStatus     string
	AvgRating      float64
}
