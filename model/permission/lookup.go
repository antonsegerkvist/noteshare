package permission

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"

	"github.com/noteshare/mysql"
	"github.com/noteshare/session"
)

//
// LookupGroupPermission gets a list of permissions from the set of permission
// keys.
//
func LookupGroupPermission(permissionSet []uint32, session *session.Session) (*ModelPermissionSet, error) {

	if len(permissionSet) == 0 {
		return &ModelPermissionSet{}, nil
	}

	var queryBuffer = bytes.Buffer{}
	queryBuffer.WriteString(`select gp.c_key, gp.c_value from t_group_permission as gp `)
	queryBuffer.WriteString(`inner join t_user_belongs_to_group as ubtg on ubtg.c_group_id = gp.c_group_id `)
	queryBuffer.WriteString(`where ubtg.c_user_id = ? and gp.c_account_id = ? and gp.c_key in `)
	queryBuffer.WriteString(`(%d`)
	queryBuffer.WriteString(strings.Repeat(",%d", len(permissionSet)-1))
	queryBuffer.WriteString(`)`)

	ret := ModelPermissionSet{}
	for _, v := range permissionSet {
		ret[v] = ModelPermission{
			Key:   v,
			Value: 0,
		}
	}

	interfacePermissionSet := make([]interface{}, len(permissionSet))
	for i := range permissionSet {
		interfacePermissionSet[i] = permissionSet[i]
	}

	db := mysql.Open()

	stmt, err := db.Prepare(fmt.Sprintf(queryBuffer.String(), interfacePermissionSet...))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(session.UserID, session.AccountID)
	if err == sql.ErrNoRows {
		return &ret, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	var keyBuffer uint32
	var valueBuffer uint32
	for rows.Next() {
		err = rows.Scan(&keyBuffer, &valueBuffer)
		if err != nil {
			return nil, err
		}
		ret[keyBuffer] = ModelPermission{
			Key:   keyBuffer,
			Value: valueBuffer,
		}
	}

	return &ret, nil

}
