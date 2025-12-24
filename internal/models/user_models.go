package models

type User struct {
	Base

	Name    string `json:"name" gorm:"type:varchar(255);not null"`
	Phone   string `json:"phone" gorm:"type:varchar(20);not null;unique;index"`
	Balance int    `json:"balance" gorm:"not null;default:0;check:balance >= 0"`
}
