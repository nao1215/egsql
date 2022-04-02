package storage

import "errors"

var (
	// ErrParseCatalogFile means that parsing of the catalog file (json file) failed
	ErrParseCatalogFile = errors.New("failed to parse catalog file")
	// ErrSaveCatalogFile means that saving of the catalog file (json file) failed
	ErrSaveCatalogFile = errors.New("failed to save catalog file")
)
