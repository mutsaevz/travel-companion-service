package models

type User struct {
	Base

	Name    string `json:"name" gorm:"type:varchar(255);not null"`
	Phone   string `json:"phone" gorm:"type:varchar(20);not null;unique;index"`
	Balance int    `json:"balance" gorm:"not null;default:0;check:balance >= 0"`
}

type UserCreateRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Balance int    `json:"balance"`
}

type UserUpdateRequest struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}

type UserFilter struct {
	Page     int
	PageSize int
}
