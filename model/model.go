package model

import "errors"

var (
	// ErrIDProductDoesNotExists the product does not exist
	ErrIDProductDoesNotExists = errors.New("the product does not exist")

	// ErrUserNotExists user does not exist
	ErrUserNotExists = errors.New("user does not exist")
)
