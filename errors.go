package main

import (
	"fmt"
	"strings"
)

type MultiError struct {
	Errors []error
}

func (e *MultiError) Error() string {
	errors := []string{}

	for _, err := range e.Errors {
		errors = append(errors, err.Error())
	}
	return fmt.Sprintf(
		"Recieved the following errors: \n%s",
		strings.Join(errors, "\n"),
	)
}

func (e *MultiError) GetError() error {
	for _, err := range e.Errors {
		if err != nil {
			return err
		}
	}
	return nil
}
