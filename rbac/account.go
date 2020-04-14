package rbac

type Account struct {
	Username string
	groups   map[string]*Group
	depts    map[string]*Dept
}

func (a *Account) HasPermission(permission Permission) (ok bool) {
	for k := range a.groups {
		if a.groups[k].HasPermission(permission) {
			return true
		}
	}
	for k := range a.depts {
		if a.depts[k].HasPermission(permission) {
			return true
		}
	}
	return false
}
