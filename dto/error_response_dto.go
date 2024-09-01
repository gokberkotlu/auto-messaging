package dto

type ErrorResponseDTO struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
