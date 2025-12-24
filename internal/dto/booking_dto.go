package dto

import "github.com/mutsaevz/team-5-ambitious/internal/constants"

type BookingCreateRequest struct {
	TripID      uint `json:"trip_id" binding:"required"`
	PassengerID uint `json:"passenger_id" binding:"required"`
}

type BookingUpdateRequest struct {
	BookingStatus *constants.BookingStatus `json:"booking_status" binding:"required"`
}
