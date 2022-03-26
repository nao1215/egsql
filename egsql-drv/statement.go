package egsql

import "database/sql/driver"

type egsqlStmt struct{}

// Close closes the statement.
func (stmt *egsqlStmt) Close() error {
	return nil
}

// NumInput returns the number of placeholder parameters.
func (stmt *egsqlStmt) NumInput() int {
	return 0
}

// Exec executes a query that doesn't return rows, such as an INSERT or UPDATE.
// Deprecated: Drivers should implement StmtExecContext instead (or additionally).
func (stmt *egsqlStmt) Exec(args []driver.Value) (driver.Result, error) {
	return &egsqlResult{}, nil
}

// Query executes a query that may return rows, such as a SELECT.
// Deprecated: Drivers should implement StmtQueryContext instead (or additionally).
func (stmt *egsqlStmt) Query(args []driver.Value) (driver.Rows, error) {
	return &egsqlRows{}, nil
}
