package egsql

import (
	"context"
	"database/sql/driver"
)

type connector struct {
	cfg *Config // immutable private copy.
}

// Connect implements driver.Connector interface.
// Connect returns a connection to the database.
func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	return &egsqlConn{}, nil
}

// Driver implements driver.Connector interface.
// Driver returns &MySQLDriver{}.
func (c *connector) Driver() driver.Driver {
	return &Driver{}
}
