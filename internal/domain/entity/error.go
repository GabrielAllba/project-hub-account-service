package entity

import "errors"

var (
	ErrNotFound      = errors.New("requested entity not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrInternal      = errors.New("internal error")
	ErrAlreadyExists = errors.New("entity already exists")
)
