package dto

type SuccessResponse[T any] struct {
	Status  int    `json:"status"`
	Data    []T    `json:"data"`
	Message string `json:"message"`
}
