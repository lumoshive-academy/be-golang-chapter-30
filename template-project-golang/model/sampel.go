package model

type Sampel struct {
	Basic
	Name  *string `json:"name" validate:"required"`
	Email string  `json:"email" validate:"required,email"`
}
