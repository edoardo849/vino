package main

import (
	"errors"
	"regexp"
)

type Letter struct {
	Char string `json:"char"`
}

func (l *Letter) isValid() error {

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
