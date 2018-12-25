package permission

import "github.com/noteshare/session"

//
// ModelPermissionSetJSON contains a collection of permissions that can be
// converted to json.
//
type ModelPermissionSetJSON struct {
	Permissions []ModelPermission `json:"permissions"`
}

//
// LookupGroupPermissionJSON gets a list of permissions from the set of
// permission keys.
//
func LookupGroupPermissionJSON(
	permissionSet []uint32,
	session *session.Session,
) (*ModelPermissionSetJSON, error) {

	permissionSetMap, err := LookupGroupPermission(permissionSet, session)
	if err != nil {
		return nil, err
	}

	ret := ModelPermissionSetJSON{Permissions: []ModelPermission{}}
	for _, v := range *permissionSetMap {
		ret.Permissions = append(ret.Permissions, ModelPermission{
			Key:   v.Key,
			Value: v.Value,
		})
	}

	return &ret, nil

}
