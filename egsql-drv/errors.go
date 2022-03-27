package egsql

import "errors"

var (
	// ErrNotGetEgSQLHomeDir indicates that the home directory path
	// that manages the files used by egsql did not get.
	ErrNotGetEgSQLHomeDir = errors.New("not found egsql home dirctory")
	// ErrNotCreateEgSQLHomeDir means that the egsql home directory
	// could not be created.
	ErrNotCreateEgSQLHomeDir = errors.New("not create egsql home dirctory")
)
