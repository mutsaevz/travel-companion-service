package dto

type CarCreateRequest struct {
	OwnerID  uint   `json:"owner_id"`
	Brand    string `json:"brand"`
	CarModel string `json:"car_model"`
	Seats    int    `json:"seats"`
}

type CarUpdateRequest struct {
	Brand    *string `json:"brand"`
	CarModel *string `json:"car_model"`
	Seats    *int    `json:"seats"`
}
