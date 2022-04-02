package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strconv"
	"sync"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nao1215/egsql/dbms/meta"
	"github.com/nao1215/egsql/misc/slice"
)

func TestNewEmtpyCatalog(t *testing.T) {
	tests := []struct {
		name string
		want *Catalog
	}{
		{
			name: "[Success] get new empty catalog",
			want: &Catalog{
				mutex: &sync.RWMutex{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEmtpyCatalog()
			if diff := cmp.Diff(*tt.want, *got, cmpopts.IgnoreFields(*got, "mutex")); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestLoadCatalog(t *testing.T) {
	type args struct {
		egsqlHomePath string
	}
	tests := []struct {
		name      string
		args      args
		want      *Catalog
		wantErr   bool
		wantErrAs error
	}{
		{
			name: "[Success] generate new catlog pointer (not unmarshal catalog.db)",
			args: args{
				egsqlHomePath: "",
			},
			want:      &Catalog{mutex: &sync.RWMutex{}},
			wantErr:   false,
			wantErrAs: nil,
		},
		{
			name: "[Success] load catlog file",
			args: args{
				egsqlHomePath: "./testdata/ok",
			},
			want: &Catalog{
				Schemes: []*meta.Scheme{
					{
						TableName:       "users",
						ColumnNames:     []string{"id", "user_id"},
						ColumnDataTypes: []meta.DataType{meta.Int, meta.Varchar},
						PrimaryKey:      "id",
					},
				},
				mutex: &sync.RWMutex{},
			},
			wantErr:   false,
			wantErrAs: nil,
		},
		{
			name: "[Error] failed to parse catalog.db",
			args: args{
				egsqlHomePath: "./testdata/ng",
			},
			want:      nil,
			wantErr:   true,
			wantErrAs: ErrParseCatalogFile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadCatalog(tt.args.egsqlHomePath)
			if (err != nil) != tt.wantErr && !errors.As(err, &tt.wantErrAs) {
				t.Errorf("LoadCatalog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil {
				return
			}

			if diff := cmp.Diff(*tt.want, *got, cmpopts.IgnoreFields(*got, "mutex")); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSaveCatalog(t *testing.T) {
	type args struct {
		egsqlHomePath string
		c             *Catalog
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
		wantErrAs error
		marshal   func(v any) ([]byte, error)
	}{
		{
			name: "[Success] Save catalog",
			args: args{
				egsqlHomePath: "testdata/save",
				c: &Catalog{
					Schemes: []*meta.Scheme{
						{
							TableName:       "users",
							ColumnNames:     []string{"id", "user_id"},
							ColumnDataTypes: []meta.DataType{meta.Int, meta.Varchar},
							PrimaryKey:      "id",
						},
					},
					mutex: &sync.RWMutex{},
				},
			},
			wantErr:   false,
			wantErrAs: nil,
			marshal:   json.Marshal,
		},
		{
			name: "[Error] failed to marshal Catlog",
			args: args{
				egsqlHomePath: "testdata/save",
				c: &Catalog{
					Schemes: []*meta.Scheme{
						{
							TableName:       "users",
							ColumnNames:     []string{"id", "user_id"},
							ColumnDataTypes: []meta.DataType{meta.Int, meta.Varchar},
							PrimaryKey:      "id",
						},
					},
					mutex: &sync.RWMutex{},
				},
			},
			wantErr:   false,
			wantErrAs: ErrSaveCatalogFile,
			marshal:   func(v any) ([]byte, error) { return []byte{}, errors.New("error") },
		},
		{
			name: "[Error] failed to write file",
			args: args{
				egsqlHomePath: "/no_exist_path",
				c: &Catalog{
					Schemes: []*meta.Scheme{
						{
							TableName:       "users",
							ColumnNames:     []string{"id", "user_id"},
							ColumnDataTypes: []meta.DataType{meta.Int, meta.Varchar},
							PrimaryKey:      "id",
						},
					},
					mutex: &sync.RWMutex{},
				},
			},
			wantErr:   false,
			wantErrAs: ErrSaveCatalogFile,
			marshal:   json.Marshal,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonMarshal = tt.marshal
			defer func() { jsonMarshal = json.Marshal }()

			err := SaveCatalog(tt.args.egsqlHomePath, tt.args.c)
			if (err != nil) != tt.wantErr && !errors.As(err, &tt.wantErrAs) {
				t.Errorf("SaveCatalog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			golden, err := ioutil.ReadFile("testdata/ok/catalog.db")
			if err != nil {
				t.Fatal(err)
			}

			got, err := ioutil.ReadFile("testdata/save/catalog.db")
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(golden, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestCatalog_Add(t *testing.T) {
	type fields struct {
		Schemes []*meta.Scheme
		mutex   *sync.RWMutex
	}
	type args struct {
		scheme *meta.Scheme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "[Success] add the new scheme into a memory",
			fields: fields{
				Schemes: make([]*meta.Scheme, 0),
				mutex:   &sync.RWMutex{},
			},
			args: args{
				scheme: &meta.Scheme{TableName: "success"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Catalog{
				Schemes: tt.fields.Schemes,
				mutex:   tt.fields.mutex,
			}
			c.Add(tt.args.scheme)

			if !slice.Contains(c.Schemes, tt.args.scheme) {
				t.Errorf("can't add schema in memory: %v", tt.args.scheme)
			}
		})
	}
}
func TestCatalog_Add_Concurrenct(t *testing.T) {
	var s []*meta.Scheme
	ctg := &Catalog{
		Schemes: s,
	}
	ctg.mutex = &sync.RWMutex{}

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		num := i
		go func() {
			scheme := &meta.Scheme{TableName: strconv.Itoa(num)}
			ctg.Add(scheme)
			wg.Done()
		}()
	}
	wg.Wait()

	want := 1000
	got := len(ctg.Schemes)
	if want != got {
		t.Errorf("mismatch: want=%d got=%d", want, got)
	}
}

func TestCatalog_HasScheme(t *testing.T) {
	type fields struct {
		Schemes []*meta.Scheme
		mutex   *sync.RWMutex
	}
	type args struct {
		tableName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "[Success] exist schema",
			fields: fields{
				Schemes: []*meta.Scheme{{TableName: "dummy"}, {TableName: "success"}},
				mutex:   &sync.RWMutex{},
			},
			args: args{
				tableName: "success",
			},
			want: true,
		},
		{
			name: "[Success] not exist schema",
			fields: fields{
				Schemes: []*meta.Scheme{{TableName: "dummy"}, {TableName: "success"}},
				mutex:   &sync.RWMutex{},
			},
			args: args{
				tableName: "not_exist",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Catalog{
				Schemes: tt.fields.Schemes,
				mutex:   tt.fields.mutex,
			}
			if got := c.HasScheme(tt.args.tableName); got != tt.want {
				t.Errorf("Catalog.HasScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCatalog_FetchScheme(t *testing.T) {
	type fields struct {
		Schemes []*meta.Scheme
		mutex   *sync.RWMutex
	}
	type args struct {
		tableName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *meta.Scheme
	}{
		{
			name: "[Success] get schema",
			fields: fields{
				Schemes: []*meta.Scheme{{TableName: "dummy"}, {TableName: "success"}},
				mutex:   &sync.RWMutex{},
			},
			args: args{
				tableName: "success",
			},
			want: &meta.Scheme{TableName: "success"},
		},
		{
			name: "[Error] not get schema",
			fields: fields{
				Schemes: []*meta.Scheme{{TableName: "dummy"}, {TableName: "success"}},
				mutex:   &sync.RWMutex{},
			},
			args: args{
				tableName: "not_exist",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Catalog{
				Schemes: tt.fields.Schemes,
				mutex:   tt.fields.mutex,
			}
			got := c.FetchScheme(tt.args.tableName)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
