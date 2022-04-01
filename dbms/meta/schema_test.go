package meta

import (
	"errors"
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
		name      string
		args      args
		want      *Scheme
		wantErr   bool
		wantErrIs error
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
			wantErr:   false,
			wantErrIs: nil,
		},
		{
			name: "[Error] Column name slice is empty",
			args: args{
				tableName:   "this_is_table_name",
				columnNames: []string{},
				dataTypes:   []DataType{Int, Int, Int, Varchar},
				pk:          "id",
			},
			want:      nil,
			wantErr:   true,
			wantErrIs: ErrColumnBelowMinNum,
		},
		{
			name: "[Error] Column data type slice is empty",
			args: args{
				tableName:   "this_is_table_name",
				columnNames: []string{"id", "user_id", "group_id", "name"},
				dataTypes:   []DataType{},
				pk:          "id",
			},
			want:      nil,
			wantErr:   true,
			wantErrIs: ErrColumnBelowMinNum,
		},
		{
			name: "[Error] 'Number of column names' and 'Number of column types' do not match",
			args: args{
				tableName:   "this_is_table_name",
				columnNames: []string{"id", "user_id", "group_id", "name"},
				dataTypes:   []DataType{Int},
				pk:          "id",
			},
			want:      nil,
			wantErr:   true,
			wantErrIs: ErrNotMatchColumnNum,
		},
		{
			name: "[Error] An empty string is specified in the column name.",
			args: args{
				tableName:   "this_is_table_name",
				columnNames: []string{"id", "user_id", "group_id", ""},
				dataTypes:   []DataType{Int, Int, Int, Varchar},
				pk:          "id",
			},
			want:      nil,
			wantErr:   true,
			wantErrIs: ErrEmptyColumnName,
		},
		{
			name: "[Error]  if you specify a column name that does not exist.",
			args: args{
				tableName:   "this_is_table_name",
				columnNames: []string{"id", "user_id", "group_id", "name"},
				dataTypes:   []DataType{Int, Int, Int, Varchar},
				pk:          "not_exist_pk",
			},
			want:      nil,
			wantErr:   true,
			wantErrIs: ErrInvalidPrimaryKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewScheme(tt.args.tableName, tt.args.columnNames, tt.args.dataTypes, tt.args.pk)
			if (err != nil) != tt.wantErr && !errors.Is(err, tt.wantErrIs) {
				t.Errorf("NewScheme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_validColumn(t *testing.T) {
	type args struct {
		columnNames []string
		dataTypes   []DataType
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantErrIs error
	}{
		{
			name: "[Success] valid ok",
			args: args{
				columnNames: []string{"id", "user_id", "group_id", "name"},
				dataTypes:   []DataType{Int, Int, Int, Varchar},
			},
			wantErr:   false,
			wantErrIs: nil,
		},
		{
			name: "[Error] Column name slice is empty",
			args: args{
				columnNames: []string{},
				dataTypes:   []DataType{Int, Int, Int, Varchar},
			},
			wantErr:   true,
			wantErrIs: ErrColumnBelowMinNum,
		},
		{
			name: "[Error] Column data type slice is empty",
			args: args{
				columnNames: []string{"id", "user_id", "group_id", "name"},
				dataTypes:   []DataType{},
			},
			wantErr:   true,
			wantErrIs: ErrColumnBelowMinNum,
		},
		{
			name: "[Error] 'Number of column names' and 'Number of column types' do not match",
			args: args{
				columnNames: []string{"id", "user_id", "group_id", "name"},
				dataTypes:   []DataType{Int},
			},
			wantErr:   true,
			wantErrIs: ErrNotMatchColumnNum,
		},
		{
			name: "[Error] An empty string is specified in the column name.",
			args: args{
				columnNames: []string{"id", "", "group_id", "name"},
				dataTypes:   []DataType{Int},
			},
			wantErr:   true,
			wantErrIs: ErrEmptyColumnName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validColumn(tt.args.columnNames, tt.args.dataTypes); (err != nil) != tt.wantErr && !errors.Is(err, tt.wantErrIs) {
				t.Errorf("validColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validPrimaryKey(t *testing.T) {
	type args struct {
		columnNames []string
		pk          string
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantErrIs error
	}{
		{
			name: "[Success] valid ok",
			args: args{
				columnNames: []string{"id", "user_id", "group_id", "name"},
				pk:          "id",
			},
			wantErr:   false,
			wantErrIs: nil,
		},
		{
			name: "[Error] if you specify a column name that does not exist.",
			args: args{
				columnNames: []string{"id", "user_id", "group_id", "name"},
				pk:          "not_exist_pk",
			},
			wantErr:   true,
			wantErrIs: ErrInvalidPrimaryKey,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validPrimaryKey(tt.args.columnNames, tt.args.pk); (err != nil) != tt.wantErr {
				t.Errorf("validPrimaryKey() error = %v, wantErr %v", err, tt.wantErr)
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
