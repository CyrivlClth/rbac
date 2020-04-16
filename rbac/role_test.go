package rbac

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewRole(t *testing.T) {
	type args struct {
		parent *Role
		opts   []RoleOption
	}
	var tests = []struct {
		name string
		args args
		want *Role
	}{
		{
			name: "顶级角色无继承",
			args: args{},
			want: &Role{
				permissions: make(map[Permission]struct{}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRole(tt.args.parent, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRole_AddPermission(t *testing.T) {
	parentRole := &Role{
		permissions: map[Permission]struct{}{"1": Exists, "2": Exists},
	}
	type fields struct {
		Id          string
		Name        string
		Code        string
		Parent      *Role
		SystemCode  string
		permissions map[Permission]struct{}
	}
	type args struct {
		permissions []Permission
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "顶级角色添加任意权限",
			fields: fields{
				permissions: make(map[Permission]struct{}),
			},
			args: args{
				permissions: []Permission{"1", "2"},
			},
			wantErr: false,
		},
		{
			name: "子级角色添加合法权限",
			fields: fields{
				Parent:      parentRole,
				permissions: make(map[Permission]struct{}),
			},
			args: args{
				permissions: []Permission{"1", "2"},
			},
			wantErr: false,
		},
		{
			name: "子级角色添加非法权限",
			fields: fields{
				Parent:      parentRole,
				permissions: make(map[Permission]struct{}),
			},
			args: args{
				permissions: []Permission{"1", "3"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Role{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Code:        tt.fields.Code,
				Parent:      tt.fields.Parent,
				SystemCode:  tt.fields.SystemCode,
				permissions: tt.fields.permissions,
			}
			if err := r.AddPermission(tt.args.permissions...); (err != nil) != tt.wantErr {
				t.Errorf("AddPermission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRoleTree_MarshalJSON(t *testing.T) {
	z := &RoleTree{
		Node: NewRole(nil),
		Children: []*RoleTree{
			{
				Node: NewRole(nil),
				Children: []*RoleTree{{
					Node:     NewRole(nil),
					Children: []*RoleTree{},
				},
					{
						Node:     NewRole(nil),
						Children: []*RoleTree{},
					},
					{
						Node:     NewRole(nil),
						Children: []*RoleTree{},
					}},
			},
			{
				Node: NewRole(nil),
				Children: []*RoleTree{{
					Node:     NewRole(nil),
					Children: []*RoleTree{},
				},
					{
						Node:     NewRole(nil),
						Children: []*RoleTree{},
					},
					{
						Node:     NewRole(nil),
						Children: []*RoleTree{},
					}},
			},
			{
				Node: NewRole(nil),
				Children: []*RoleTree{{
					Node:     NewRole(nil),
					Children: []*RoleTree{},
				},
					{
						Node:     NewRole(nil),
						Children: []*RoleTree{},
					},
					{
						Node:     NewRole(nil),
						Children: []*RoleTree{},
					}},
			},
		},
	}
	b, _ := json.Marshal(z)
	t.Log(string(b))
}
