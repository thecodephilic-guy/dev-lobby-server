package models

type SuccessResponse struct {
	StatusCode int    `json:"status"`
	Title      string `json:"title"`
	Data       any    `json:"data"`
}
