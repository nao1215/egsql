package meta

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewScheme(t *testing.T) {
	type args struct {
		tableName   string
		columnNames []string
		dataTypes   []DataType
		pk          string
	}
	tests := []struct {
		name string
		args args
		want *Scheme
	}{
		{
			name: "[Success] generate new table",
			args: args{
				tableName:   "this_is_table_name",
				columnNames: []string{"id", "user_id", "group_id", "name"},
				dataTypes:   []DataType{Int, Int, Int, Varchar},
				pk:          "id",
			},
			want: &Scheme{
				TableName:       "this_is_table_name",
				ColumnNames:     []string{"id", "user_id", "group_id", "name"},
				ColumnDataTypes: []DataType{Int, Int, Int, Varchar},
				PrimaryKey:      "id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewScheme(tt.args.tableName, tt.args.columnNames, tt.args.dataTypes, tt.args.pk)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestDataType_String(t *testing.T) {
	tests := []struct {
		name string
		d    DataType
		want string
	}{
		{
			name: "[Succes] get 'int' from Int type",
			d:    Int,
			want: "int",
		},
		{
			name: "[Succes] get 'varchar' from Varchar type",
			d:    Varchar,
			want: "varchar",
		},
		{
			name: "[Error] get 'undefined' from undifiend type",
			d:    DataType(0),
			want: "undefined",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("DataType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheme_ConvertToTable(t *testing.T) {
	type fields struct {
		TableName       string
		ColumnNames     []string
		ColumnDataTypes []DataType
		PrimaryKey      string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Table
	}{
		{
			name: "[Succes] convert schema to table",
			fields: fields{
				TableName:       "test_table",
				ColumnNames:     []string{"id", "user_id", "group_id", "name"},
				ColumnDataTypes: []DataType{Int, Int, Int, Varchar},
				PrimaryKey:      "id",
			},
			want: &Table{
				Name: "test_table",
				Columns: []Column{
					{
						Name:    "id",
						Type:    Int,
						Primary: true,
					},
					{
						Name:    "user_id",
						Type:    Int,
						Primary: false,
					},
					{
						Name:    "group_id",
						Type:    Int,
						Primary: false,
					},
					{
						Name:    "name",
						Type:    Varchar,
						Primary: false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Scheme{
				TableName:       tt.fields.TableName,
				ColumnNames:     tt.fields.ColumnNames,
				ColumnDataTypes: tt.fields.ColumnDataTypes,
				PrimaryKey:      tt.fields.PrimaryKey,
			}
			got := s.ConvertToTable()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
