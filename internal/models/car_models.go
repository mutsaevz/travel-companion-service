package models

type Car struct {
	Base

	OwnerID  uint
	Brand    string
	CarModel string
	Seats    int
}
