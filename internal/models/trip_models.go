package models

import (
	"time"

	"github.com/mutsaevz/team-5-ambitious/internal/constants"
)

type Trip struct {
	Base

	DriverID       uint      `json:"driver_id" gorm:"not null;index"`
	CarID          uint      `json:"car_id" gorm:"not null;index"`
	FromCity       string    `json:"from_city" gorm:"type:varchar(100);not null;index"`
	ToCity         string    `json:"to_city" gorm:"type:varchar(100);not null;index"`
	StartTime      time.Time `json:"start_time" gorm:"not null;index"`
	DurationMin    int       `json:"duration_min" gorm:"not null"`
	TotalSeats     int       `json:"total_seats" gorm:"not null;check:total_seats > 0"`
	AvailableSeats int       `json:"available_seats" gorm:"not null;index;check:available_seats >= 0"`
	Price          int       `json:"price" gorm:"not null;check:price >= 0"`
	TripStatus     string    `json:"trip_status" gorm:"type:varchar(50);not null;index"`
	AvgRating      float64   `json:"avg_rating" gorm:"default:0.0;check:avg_rating >= 0 AND avg_rating <= 5"`
}

type TripCreateRequest struct {
	FromCity       string               `json:"from_city"`
	ToCity         string               `json:"to_city"`
	StartTime      time.Time            `json:"start_time"`
	DurationMin    int                  `json:"duration_min"`
	AvailableSeats int                  `json:"available_seats"`
	Price          int                  `json:"price"`
	TripStatus     constants.TripStatus `json:"trip_status"`
}

type TripFilter struct {
	FromCity       *string
	ToCity         *string
	StartTime      *time.Time
	AvailableSeats *int
	TripStatus     *constants.TripStatus

	Page     int
	PageSize int
}

type TripUpdateRequest struct {
	FromCity       *string               `json:"from_city"`
	ToCity         *string               `json:"to_city"`
	StartTime      *time.Time            `json:"start_time"`
	DurationMin    *int                  `json:"duration_min"`
	AvailableSeats *int                  `json:"available_seats"`
	Price          *int                  `json:"price"`
	TripStatus     *constants.TripStatus `json:"trip_status"`
}
