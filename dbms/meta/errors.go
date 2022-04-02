package meta

import "errors"

var (
	// ErrNotMatchColumnNum means that "'number of column names' and 'number of column types' do not match"
	ErrNotMatchColumnNum = errors.New("'number of column names' and 'number of column types' do not match")
	// ErrColumnBelowMinNum means that "columns are below the minimum number"
	ErrColumnBelowMinNum = errors.New("columns are below the minimum number")
	// ErrEmptyColumnName means that use an empty ("") string for a column name.
	ErrEmptyColumnName = errors.New("column name is empty")
	// ErrInvalidPrimaryKey means that the primary key is invalid.
	// For example, if you specify a column name that does not exist.
	ErrInvalidPrimaryKey = errors.New("invalid primary key")
)
