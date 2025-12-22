package models

type Booking struct {
	Base

	TripID        uint   `json:"trip_id"`
	PassengerID   uint   `json:"passenger_id"`
	BookingStatus string `json:"booking_status"`
}

type BookingCreateRequest struct {
	TripID      uint `json:"trip_id" binding:"required"`
	PassengerID uint `json:"passenger_id" binding:"required"`
}

type BookingUpdateRequest struct {
	BookingStatus *string `json:"booking_status" binding:"required"`
}
