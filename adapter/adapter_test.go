package adapter

import (
	"reflect"
	"testing"

	"github.com/CyrivlClth/rbac/rbac"
)

func TestAccountAdapter(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want *rbac.Account
	}{
		{
			name: "",
			args: args{username: "acc"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AccountAdapter(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}
