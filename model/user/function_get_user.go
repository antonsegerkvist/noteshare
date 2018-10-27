package user

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// GetUser returns all the users that the user can see
// in the same account.
//
func GetUser(targetUserID, userID uint64) (*ModelUser, error) {

	const query = `
		select u1.c_id, u1.c_email, u1.c_username, u1.c_created
		from t_user as u1
		inner join t_user as u2 on u1.c_account_id = u2.c_account_id
		where u1.c_id = ? and u2.c_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	ret := ModelUser{}
	row := stmt.QueryRow(targetUserID, userID)
	err = row.Scan(
		&ret.ID,
		&ret.Email,
		&ret.Username,
		&ret.Created,
	)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return &ret, nil

}
