package rbac

type Permission string

var Exists = struct{}{}

type PermissionGroup interface {
	HasPermission(permission Permission) (ok bool)
}
