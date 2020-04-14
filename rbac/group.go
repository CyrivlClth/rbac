package rbac

type Group struct {
	*RoleGroup
	Id   string
	Name string
}
