package validator

import (
	"github.com/go-playground/validator"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("duration-string", parseDuration)
	return validate
}
