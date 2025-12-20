package models

type Booking struct {
	Base

	TripID        uint
	PassengerID   uint
	BookingStatus string
}
