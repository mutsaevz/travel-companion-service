package models

type Car struct {
	Base

	OwnerID  uint   `json:"owner_id" gorm:"not null;index"`
	Brand    string `json:"brand" gorm:"type:varchar(255);not null"`
	CarModel string `json:"car_model" gorm:"type:varchar(255);not null"`
	Seats    int    `json:"seats" gorm:"not null;check:seats > 0"`
}
