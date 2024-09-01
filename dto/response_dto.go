package dto

type ResponseDTO[T any] struct {
	Response T `json:"response"`
}
