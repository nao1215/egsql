package meta

import (
	"reflect"
	"testing"
)

func TestNewResultSet(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *ResultSet
	}{
		{
			name: "[Success] get new ResultSet pointer",
			args: args{
				message: "new result set",
			},
			want: &ResultSet{
				Message: "new result set",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResultSet(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResultSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
