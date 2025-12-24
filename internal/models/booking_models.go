package models

import "github.com/mutsaevz/team-5-ambitious/internal/constants"

type Booking struct {
	Base

	TripID        uint   `json:"trip_id" gorm:"not null;index"`
	PassengerID   uint   `json:"passenger_id" gorm:"not null;index"`
	BookingStatus constants.BookingStatus `json:"booking_status" gorm:"type:varchar(50);not null;index"`
}
