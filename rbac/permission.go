package rbac

type Permission string

var Exists = struct{}{}

// 权限集合
type PermissionGroup interface {
	// 是否拥有某项权限
	HasPermission(permission Permission) (ok bool)
}
