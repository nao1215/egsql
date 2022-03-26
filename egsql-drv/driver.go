package egsql

import (
	"database/sql"
	"database/sql/driver"
)

// Driver is exported to make the sql driver directly accessible.
// In general the driver is used via the database/sql package.
type Driver struct{}

// init registers the egsql driver using database/sql.Register().
func init() {
	sql.Register("egsql", &Driver{})
}

// Open new Connection.
func (d Driver) Open(dsn string) (driver.Conn, error) {
	return &egsqlConn{}, nil
}

// OpenConnector implements driver.DriverContext.
func (d Driver) OpenConnector(dsn string) (driver.Connector, error) {
	return &connector{}, nil
}
