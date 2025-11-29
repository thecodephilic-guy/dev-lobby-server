package models

type ErrorResponse struct {
	StatusCode int    `json:"status"`
	Title      string `json:"title"`
	Error      string `json:"error"`
}
