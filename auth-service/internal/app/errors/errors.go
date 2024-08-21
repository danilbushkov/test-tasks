package errors

import (
	"errors"
	"fmt"
)

var (
	ErrDatabase        = errors.New("Database error")
	ErrToken           = errors.New("InternalServerError")
	ErrTokenIsNotValid = errors.New("Token is not valid")
)

func TokenError(err error) error {
	return fmt.Errorf("%w: %w", ErrToken, err)
}

func DatabaseError(err error) error {
	return fmt.Errorf("%w:  %w", ErrDatabase, err)
}
