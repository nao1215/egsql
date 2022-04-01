package meta

import "errors"

var (
	// ErrNotMatchColumnNum means that "'Number of column names' and 'Number of column types' do not match"
	ErrNotMatchColumnNum = errors.New("'Number of column names' and 'Number of column types' do not match")
	// ErrColumnBelowMinNum means that "Columns are below the minimum number"
	ErrColumnBelowMinNum = errors.New("Columns are below the minimum number")
)
