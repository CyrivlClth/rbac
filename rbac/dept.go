package rbac

// 部门
type Dept struct {
	*RoleGroup
	Id     string
	Name   string
	Code   string
	Parent *Dept
}

func (d *Dept) AddRole(roles ...*Role) error {
	if d.Parent != nil && !d.Parent.ContainsRole(roles...) {
		return UndatedPermissionErr
	}
	return d.RoleGroup.AddRole(roles...)
}

func (d *Dept) ContainsRole(roles ...*Role) bool {
	for i := range roles {
		if !d.RoleGroup.HasRole(roles[i]) {
			return false
		}
	}
	return true
}
