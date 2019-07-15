package bindings

import (
	"errors"

	"github.com/saveav/pelacore/app/database"
)

// DatabaseQuery is a payload about running an SQL query
type DatabaseQuery struct {
	Connection struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
		SSLMode  bool   `json:"sslmode"`
	} `json:"connection"`
	Query string `json:"query"`
}

var (
	// ErrBlankDriver is an error when the `driver` is blank
	ErrBlankDriver = errors.New("driver cannot be blank")
	// ErrBlankHost is an error when the `host` is blank
	ErrBlankHost = errors.New("host cannot be blank")
	// ErrBlankUsername is an error when the `username` is blank
	ErrBlankUsername = errors.New("username cannot be blank")
	// ErrBlankQuery is when the query is blank
	ErrBlankQuery = errors.New("query cannot be blank")
)

// Validate is an implementation of Validatable for DatabaseQuery
func (dbq *DatabaseQuery) Validate() error {
	errs := new(PayloadErrors)
	if dbq.Connection.Driver == "" {
		errs.Append(ErrBlankDriver)
	}

	if dbq.Connection.Host == "" {
		errs.Append(ErrBlankHost)
	}

	if dbq.Connection.Username == "" {
		errs.Append(ErrBlankUsername)
	}

	if dbq.Query == "" {
		errs.Append(ErrBlankQuery)
	}

	return errs
}

// DatabaseConnection instantiates a database connection object
func (dbq *DatabaseQuery) DatabaseConnection() database.Connection {
	conn := database.Connection{
		Driver:   dbq.Connection.Driver,
		Host:     dbq.Connection.Host,
		Port:     dbq.Connection.Port,
		Username: dbq.Connection.Username,
		Password: dbq.Connection.Password,
		Database: dbq.Connection.Database,
		SSLMode:  dbq.Connection.SSLMode,
	}
	return conn
}
