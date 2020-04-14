package rbac

type Dept struct {
	*RoleGroup
	Id     string
	Name   string
	Code   string
	Parent *Dept
}
