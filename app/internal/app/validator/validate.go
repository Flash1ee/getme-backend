package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var customValidator *validator.Validate

func NewValidator() *validator.Validate {
	if customValidator == nil {
		fmt.Println("create validator")
		customValidator = validator.New()
	}
	return customValidator
}
