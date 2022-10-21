package service

import "github.com/go-playground/validator"

// Validator is the interface that wraps the Validate function.
type Validator interface {
	Validate(i interface{}) error
}

// New func to create Custom validator object
func NewValidator() *Custom {
	return &Custom{validator: validator.New()}
}

// Custom wraper for validator
type Custom struct {
	validator *validator.Validate
}

// Validate function to comply with the interface of echo
func (cv *Custom) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
