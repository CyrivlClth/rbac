package rbac

import (
	"reflect"
	"testing"
)

func TestNewRole(t *testing.T) {
	tests := []struct {
		name string
		want *Role
	}{
		{
			name: "new role",
			want: &Role{
				permissions: map[Permission]struct{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRole(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRole_AddPermission(t *testing.T) {
	type fields struct {
		permissions map[Permission]struct{}
	}
	type args struct {
		permissions []Permission
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "add one to empty",
			fields: fields{
				permissions: map[Permission]struct{}{},
			},
			args: args{
				permissions: []Permission{"create"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Role{
				permissions: tt.fields.permissions,
			}
			r.AddPermission(tt.args.permissions...)
		})
	}
}

func TestRole_HasPermission(t *testing.T) {
	type fields struct {
		permissions map[Permission]struct{}
	}
	type args struct {
		permission Permission
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
		{
			name:   "has it",
			fields: fields{map[Permission]struct{}{"create": Exists}},
			args:   args{"create"},
			wantOk: true,
		},
		{
			name:   "not it",
			fields: fields{map[Permission]struct{}{"create": Exists}},
			args:   args{"update"},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Role{
				permissions: tt.fields.permissions,
			}
			if gotOk := r.HasPermission(tt.args.permission); gotOk != tt.wantOk {
				t.Errorf("HasPermission() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRole_RemovePermission(t *testing.T) {
	type fields struct {
		permissions map[Permission]struct{}
	}
	type args struct {
		permissions []Permission
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = &Role{
				permissions: tt.fields.permissions,
			}
		})
	}
}
