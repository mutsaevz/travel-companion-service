package models

type Review struct {
	Base

	AuthorID uint   `json:"author_id" gorm:"not null;index"`
	TripID   uint   `json:"trip_id" gorm:"not null;index"`
	Text     string `json:"text" gorm:"type:text;not null"`
	Rating   int    `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
}

type ReviewCreateRequest struct {
	AuthorID uint   `json:"author_id"`
	TripID   uint   `json:"trip_id"`
	Text     string `json:"text"`
	Rating   int    `json:"rating"`
}

type ReviewUpdateRequest struct {
	Text   *string `json:"text"`
	Rating *int    `json:"rating"`
}
