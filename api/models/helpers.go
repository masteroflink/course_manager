package models

import (
	"errors"

	"github.com/badoux/checkmail"
	"github.com/nyaruka/phonenumbers"
)

func ValidatePhoneNumber(value string) error {
	// Phone is not required
	if len(value) <= 0 {
		return nil
	}

	phone, err := phonenumbers.Parse(value, "US")

	if err != nil {
		return err
	}

	if phonenumbers.IsPossibleNumber(phone) {
		return nil
	} else {
		return errors.New("Invalid phone number")
	}
}

func ValidateEmail(value string) error {
	if err := checkmail.ValidateFormat(value); err != nil {
		return errors.New("Invalid Email")
	}
	return nil
}
