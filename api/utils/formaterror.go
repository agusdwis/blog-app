package utils

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "username") {
		return errors.New("Username Already Exists")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Exists")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Exists")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}

	return errors.New("Incorrect Details")
}
