package helper

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)


func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func ParseStringToInt(value string) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Error(err)
		return 0
	}
	return i
}