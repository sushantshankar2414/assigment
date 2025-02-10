package utils

import (
	"errors"
)

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	return nil
}

func ValidateDate(date string) error {
	if date == "" {
		return errors.New("date cannot be empty")
	}
	return nil
}