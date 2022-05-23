package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

var customValidator *validator.Validate

func NewValidator() *validator.Validate {
	if customValidator == nil {
		fmt.Println("create validator")
		customValidator = validator.New()
	}
	return customValidator
}
