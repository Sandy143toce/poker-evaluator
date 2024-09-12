package models

// ErrorResponse struct
type ErrorResponse struct {
	Code    int    `json:"code"`
	Key     string `json:"key"`
	Message string `json:"message"`
	Details string `json:"details"`
	Error   string `json:"error"`
}
