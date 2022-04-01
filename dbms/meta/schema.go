package meta

import "github.com/nao1215/egsql/misc/slice"

// DataType is the data type of the table column. It is Enum.
type DataType uint8

const (
	// Int is an Integer type.
	Int DataType = iota + 1
	// Varchar is a CHARACTER VARYING type. It means a variable-length string.
	Varchar
)

// Table represents a DB table.
type Table struct {
	// Name is DB table name.
	Name string
	// Columns is an slice that holds everything involved in the table.
	Columns []Column
}

// Column represents a DB table column.
type Column struct {
	// Name is column name in table.
	Name string
	// Type is column data type.
	Type DataType
	// Primary is a flag indicating whether the column is a primary key or not.
	Primary bool
}

// Scheme is the definition of tables and Columns
type Scheme struct {
	// TableName is table name.
	TableName string `json:"tableName"`
	// ColumnNames is an slice of all column names.
	ColumnNames []string `json:"columnNames"`
	// ColumnDataTypes is an slice of all column data type.
	ColumnDataTypes []DataType `json:"dataTypes"`
	// PrimaryKey is primary key.
	PrimaryKey string `json:"pk"`
}

// NewScheme returns a pointer to the new schema.
func NewScheme(tableName string, columnNames []string, dataTypes []DataType, pk string) (*Scheme, error) {
	if err := validColumn(columnNames, dataTypes); err != nil {
		return nil, err
	}
	return &Scheme{
		TableName:       tableName,
		ColumnNames:     columnNames,
		ColumnDataTypes: dataTypes,
		PrimaryKey:      pk,
	}, nil
}

// validColumn is the validation when initializing the schema.
func validColumn(columnNames []string, dataTypes []DataType) error {
	if len(columnNames) == 0 || len(dataTypes) == 0 {
		return ErrColumnBelowMinNum
	}
	if len(columnNames) != len(dataTypes) {
		return ErrNotMatchColumnNum
	}
	if slice.Contains(columnNames, "") {
		return ErrEmptyColumnName
	}
	return nil
}

// String is stringer for DataType
func (d DataType) String() string {
	switch d {
	case Int:
		return "int"
	case Varchar:
		return "varchar"
	default:
		return "undefined"
	}
}

// ConvertToTable converts a Schema structure to a Table structure.
func (s *Scheme) ConvertToTable() *Table {
	var t Table
	t.Name = s.TableName

	var columns []Column
	for i := range s.ColumnNames {
		var col Column
		col.Name = s.ColumnNames[i]
		col.Type = s.ColumnDataTypes[i]
		col.Primary = (col.Name == s.PrimaryKey)
		columns = append(columns, col)
	}
	t.Columns = columns
	return &t
}
