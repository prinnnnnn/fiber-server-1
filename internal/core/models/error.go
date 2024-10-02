package models

import "errors"

var (
	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternalServer = errors.New("internal server error")
	// an error for when data requested not found in database
	ErrDataNotFound = errors.New("data not found")
)
