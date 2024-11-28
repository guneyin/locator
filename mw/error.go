package mw

import "errors"

var (
	ErrNotFound  = errors.New("not found")
	ErrInvalidId = errors.New("invalid id")
)
