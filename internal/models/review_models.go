package models

type Review struct {
	Base

	AuthorID uint   `json:"author_id"`
	TripID   uint   `json:"trip_id"`
	Text     string `json:"text"`
	Rating   int    `json:"rating"`
}

type ReviewCreateRequest struct {
	AuthorID uint   `json:"author_id"`
	TripID   uint   `json:"trip_id" `
	Text     string `json:"text"`
	Rating   int    `json:"rating"`
}

type ReviewUpdateRequest struct {
	Text   *string `json:"text"`
	Rating *int    `json:"rating"`
}
