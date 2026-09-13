package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return "validation error"
}

func repository(id int) error {
	if id == 0 {
		return ErrNotFound
	}

	if id < 0 {
		return &ValidationError{Field: "id"}
	}

	return nil
}

func service(id int) error {
	err := repository(id)
	if err != nil {
		return fmt.Errorf("service failed: %w", err)
	}

	return nil
}

// handler must handle ErrNotFound and *ValidationError, and return other errors unchanged.
func handler(id int) error {
	err := service(id)

	var validationErr *ValidationError

	if errors.As(err, &validationErr) {
		fmt.Println("invalid field:", validationErr.Field)
		return nil
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println("не найдено")
		return nil
	}
	if err != nil {
		return err
	}
	return nil
}

func main() {
}
