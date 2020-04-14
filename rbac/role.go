package rbac

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

func (r *RoleGroup) AddRole(roles ...*Role) {
	for i := range roles {
		r.roles[roles[i].Id] = roles[i]
	}
}

func (r *RoleGroup) RemoveRole(roles ...*Role) {
	for i := range roles {
		delete(r.roles, roles[i].Id)
	}
}

func (r *RoleGroup) HasRole(role *Role) (ok bool) {
	_, ok = r.roles[role.Id]
	return
}

type Role struct {
	Id          string
	Name        string
	Code        string
	Parent      *Role
	SystemCode  string
	permissions map[Permission]struct{}
}

func NewRole() *Role {
	return &Role{permissions: make(map[Permission]struct{})}
}

func (r *Role) AddPermission(permissions ...Permission) {
	for _, p := range permissions {
		r.permissions[p] = Exists
	}
}

func (r *Role) RemovePermission(permissions ...Permission) {
	for _, p := range permissions {
		delete(r.permissions, p)
	}
}

func (r *Role) HasPermission(permission Permission) (ok bool) {
	_, ok = r.permissions[permission]
	return
}
