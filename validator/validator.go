package validator

import govalidator "github.com/go-playground/validator"

// // Validator is the interface that wraps the Validate function.
// type Validator interface {
// 	Validate(i interface{}) error
// }

// New func to create Custom validator object
func New() *customValidator {
	return &customValidator{validator: govalidator.New()}
}

// customValidator wraper for validator
type customValidator struct {
	validator *govalidator.Validate
}

// Validate function to comply with the interface of echo
func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
