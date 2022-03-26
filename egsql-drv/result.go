package egsql

type egsqlResult struct {
	affectedRows int64
	insertID     int64
}

// LastInsertId returns the ID of the last record inserted.
func (res *egsqlResult) LastInsertId() (int64, error) {
	return res.insertID, nil
}

// RowsAffected returns the number of data affected by the query operation.
func (res *egsqlResult) RowsAffected() (int64, error) {
	return res.affectedRows, nil
}
