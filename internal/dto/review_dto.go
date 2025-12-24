package dto

type ReviewCreateRequest struct {
	AuthorID uint   `json:"author_id" binding:"required,gt=0"`
	Text     string `json:"text" binding:"required,min=3"`
	Rating   int    `json:"rating" binding:"required,min=1,max=5"`
}
type ReviewUpdateRequest struct {
	Text   *string `json:"text"`
	Rating *int    `json:"rating"`
}
