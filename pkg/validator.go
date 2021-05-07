package pkg

import "gopkg.in/go-playground/validator.v9"

// CustomValidator validation that handle validation
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
