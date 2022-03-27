package egsql

import "errors"

var (
	// ErrNoEgSQLHomeDir indicates that the home directory that manages the files used by egsql did not exist.
	ErrNoEgSQLHomeDir = errors.New("not found egsql home dirctory")
)
