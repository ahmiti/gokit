package config

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(cfg interface{}) error {
	return validate.Struct(cfg)
}
