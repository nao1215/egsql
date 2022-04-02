package meta

// ResultSet is the result of a table operation with SQL Query
type ResultSet struct {
	Message     string
	ColumnNames []string
	Values      []string
}

// NewResultSet returns ResultSet pointer
func NewResultSet(message string) *ResultSet {
	return &ResultSet{
		Message: message,
	}
}
