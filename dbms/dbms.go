package dbms

// EgSQLDB is the kernel of the DB management system.
type EgSQLDB struct {
}

// NewEgSQLDB return EgSQLDB instance.
func NewEgSQLDB() (*EgSQLDB, error) {
	// [ENV]
	// database name
	// user name
	// password
	// EgSQL HOME directory path
	return &EgSQLDB{}, nil
}
