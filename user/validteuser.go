package user

import (
	"errors"
	"fmt"
	"regexp"
)

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) ValidateUser() []error {
	var errors []error

	if u.Name == "" {
		errors = append(errors, errors.New("Name: couldn't be empty"))
	}
	if len(u.Name) < 6 {
		errors = append(errors, errors.New("Name: length cannot be less than 6 characters"))
	}

	if u.Age <= 0 || u.Age > 120 {
		errors = append(errors, errors.New("Age: should be between 1 and 120"))
	}

	if u.Email == "" {
		errors = append(errors, errors.New("Email: couldn't be empty"))
	} else {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(u.Email) {
			errors = append(errors, errors.New("Email: should be in the format example@domain.com"))
		}
	}

	return errors
}

