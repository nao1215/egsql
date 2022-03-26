package egsql

type egsqlTx struct {
}

// Commit confirms changes to the database
func (tx *egsqlTx) Commit() (err error) {
	return nil
}

// Rollback undoes changes to the database.
func (tx *egsqlTx) Rollback() (err error) {
	return nil
}
