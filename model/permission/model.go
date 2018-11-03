package permission

//
// ModelPermission contains a single permission entry
//
type ModelPermission struct {
	Key   uint32 `json:"key"`
	Value uint32 `json:"value"`
}

//
// ModelPermissionSet contains a collection of permission entries.
//
type ModelPermissionSet map[uint32]ModelPermission

//
// ModelPermissionSetJSON contains a collection of permissions that can be
// converted to json.
//
type ModelPermissionSetJSON struct {
	Permissions []ModelPermission `json:"permissions"`
}
