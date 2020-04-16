package rbac

import (
	"encoding/json"
	"errors"
	"strconv"
)

var (
	UndatedPermissionErr = errors.New("UNDATED_PERMISSION_ERR")
)

// 角色集合
type RoleGroup struct {
	roles map[string]*Role
}

func (r *RoleGroup) HasPermission(permission Permission) (ok bool) {
	for k := range r.roles {
		if r.roles[k].HasPermission(permission) {
			return true
		}
	}
	return false
}

// 添加角色
func (r *RoleGroup) AddRole(roles ...*Role) error {
	for i := range roles {
		r.roles[roles[i].Id] = roles[i]
	}
	return nil
}

// 移除角色
func (r *RoleGroup) RemoveRole(roles ...*Role) {
	for i := range roles {
		delete(r.roles, roles[i].Id)
	}
}

// 是否拥有某项角色
func (r *RoleGroup) HasRole(role *Role) (ok bool) {
	_, ok = r.roles[role.Id]
	return
}

type RoleOption func(role *Role)

// 角色
type Role struct {
	Id          string                  `comment:"角色ID"`
	Name        string                  `json:"name,omitempty" comment:"角色名称"`
	Code        string                  `json:"code,omitempty" comment:"角色对应code码"`
	Parent      *Role                   `json:"-" comment:"父级角色"`
	SystemCode  string                  `json:"name,omitempty" comment:"角色系统码"`
	permissions map[Permission]struct{} // 权限集合
}

// func (r *Role) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(map[string]string{"id": r.Id})
// }

var id = 1

func NewRole(parent *Role, opts ...RoleOption) *Role {
	role := &Role{
		Id:          strconv.Itoa(id),
		Parent:      parent,
		permissions: make(map[Permission]struct{}),
	}
	for _, opt := range opts {
		opt(role)
	}
	id++
	return role
}

// 继承父级角色的所有权限
func WithExtendParent() RoleOption {
	return func(role *Role) {
		if role.Parent != nil {
			for p := range role.Parent.permissions {
				role.permissions[p] = Exists
			}
		}
	}
}

// 添加权限
// 只能继承父级权限
func (r *Role) AddPermission(permissions ...Permission) (err error) {
	if r.Parent != nil && !r.Parent.ContainsPermission(permissions...) {
		return UndatedPermissionErr
	}
	for _, p := range permissions {
		r.permissions[p] = Exists
	}
	return
}

// 获取所有权限列表
func (r *Role) ListPermissions() []Permission {
	ps := make([]Permission, 0)
	for k := range r.permissions {
		ps = append(ps, k)
	}
	return ps
}

// 移除权限
func (r *Role) RemovePermission(permissions ...Permission) {
	for _, p := range permissions {
		delete(r.permissions, p)
	}
}

func (r *Role) HasPermission(permission Permission) (ok bool) {
	_, ok = r.permissions[permission]
	return
}

func (r *Role) ContainsPermission(permissions ...Permission) bool {
	for k := range permissions {
		if _, ok := r.permissions[permissions[k]]; !ok {
			return false
		}
	}
	return true
}

type RoleTree struct {
	Node     *Role
	Children []*RoleTree
}

func (r RoleTree) MarshalJSON() ([]byte, error) {
	// if len(r.Children) == 0 {
	// 	return json.Marshal(r.Node)
	// }
	return json.Marshal(map[string]interface{}{"id": r.Node.Id, "children": r.Children})
}

func MakeTree(roles []*Role, root *RoleTree) {
	children, ok := haveChild(roles, root)
	if ok {
		root.Children = append(root.Children, children[0:]...)
		for _, v := range children {
			_, has := haveChild(roles, v)
			if has {
				MakeTree(roles, v)
			}
		}
	}
}

func haveChild(roles []*Role, root *RoleTree) (children []*RoleTree, ok bool) {
	for i := range roles {
		if roles[i].Parent == root.Node {
			children = append(children, &RoleTree{Node: roles[i]})
		}
	}
	ok = children != nil
	return
}
