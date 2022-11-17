package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type ResponseOK struct {
	Message string `json:"message"`
}
