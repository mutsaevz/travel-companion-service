package models

type Review struct {
	Base

	AuthorID uint
	TripID   uint
	Text     string
	Rating   int
}
