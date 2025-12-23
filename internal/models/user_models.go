package models

type User struct {
	Base

	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Balance int    `json:"balance"`
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
