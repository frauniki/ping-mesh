package validator

import (
	"time"

	"github.com/go-playground/validator"
)

func parseDuration(fl validator.FieldLevel) bool {
	if _, err := time.ParseDuration(fl.Field().String()); err != nil {
		return false
	}

	return true
}
