package repository

import "errors"

// Общие sentinel-ошибки, возвращаемые слоями репозиториев
var (
	ErrNotFound = errors.New("resource not found")
)
