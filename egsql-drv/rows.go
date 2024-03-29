package egsql

import (
	"database/sql/driver"
)

type egsqlRows struct{}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice. If a particular column name isn't known, an empty
// string should be returned for that entry.
func (rows *egsqlRows) Columns() []string {
	return []string{}
}

// Close closes the rows iterator.
func (rows *egsqlRows) Close() (err error) {
	return nil
}

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// Next should return io.EOF when there are no more rows.
//
// The dest should not be written to outside of Next. Care
// should be taken when closing Rows not to modify
// a buffer held in dest.
func (rows *egsqlRows) Next(dest []driver.Value) error {
	return nil
}
