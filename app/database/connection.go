package database

import "strings"

// Connection represents a database connection parameters
type Connection struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  bool
}

func (dbcon *Connection) String() (string, error) {
	switch dbcon.Driver {
	case Postgres:
		return dbcon.pgConString(), nil
	default:
		return "", ErrUnknownDriver
	}
}

// pgConString returns a connection string to connect to Postgres
// please refer to the following doc for list of available params
// https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
func (dbcon *Connection) pgConString() string {
	params := []string{}

	if dbcon.Host != "" {
		params = append(params, "host="+dbcon.Host)
	}

	if dbcon.Port != "" {
		params = append(params, "port="+dbcon.Port)
	}

	if dbcon.Username != "" {
		params = append(params, "user="+dbcon.Username)
	}

	if dbcon.Password != "" {
		params = append(params, "password="+dbcon.Password)
	}

	if dbcon.Database != "" {
		params = append(params, "dbname="+dbcon.Database)
	}

	if dbcon.SSLMode == true {
		params = append(params, "sslmode=require")
	} else {
		params = append(params, "sslmode=disable")
	}

	return strings.Join(params, " ")
}
