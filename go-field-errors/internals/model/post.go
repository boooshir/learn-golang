package model

type Post struct {
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=100"`
}

type PostResponse struct {
	Data Post `json:"data"`
}

type ErrorResponse struct {
	Errors map[string]string `json:"errors"`
}
