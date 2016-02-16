package main

import (
	"errors"
	"regexp"
)

//Letter is the model for the user-answer
type Letter struct {
	Char string `json:"char"`
}

//IsValid checks the validity of the specified letter,
//in short it must be a single-digit a-z char, everything
//else will throw a soft-error
func (l *Letter) IsValid() error {

	r, err := regexp.Compile(`^[a-z]$`)

	if err != nil {
		// Problem with the regexp
		panic(err)
	}

	if !r.MatchString(l.Char) {
		return errors.New("Only single-digit, a-z characters allowed")
	}

	return nil
}
