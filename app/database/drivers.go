package database

import "errors"

const (
	// Postgres represents the postgres database
	Postgres = "postgres"
)

// ErrUnknownDriver is an error given when a driver name is not among the known types
var ErrUnknownDriver = errors.New("unknown driver")
