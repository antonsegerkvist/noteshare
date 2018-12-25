package user

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// ModelUser contains information about a single user.
//
type ModelUser struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Created  string `json:"created"`
}

//
// ModelUsers contains information about multiple user.
//
type ModelUsers struct {
	Users []ModelUser `json:"users"`
}

//
// GetUser returns information about a single user in the specified account.
//
func GetUser(userID, accountID uint64) (*ModelUser, error) {

	const query = `
		select u.c_id, u.c_email, u.c_username, u.c_created
		from t_user as u where u.c_id = ? and u.c_account_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	ret := ModelUser{}
	row := stmt.QueryRow(userID, accountID)
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

//
// GetUsers returns information about all users in the spcified account.
//
func GetUsers(accountID uint64) (*ModelUsers, error) {

	const query = `
		select u.c_id, u.c_email, u.c_username, u.c_created
		from t_user as u where u.c_account_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret := ModelUsers{Users: []ModelUser{}}
	rows, err := stmt.Query(accountID)
	if err == sql.ErrNoRows {
		return &ModelUsers{Users: []ModelUser{}}, nil
	} else if err != nil {
		return nil, err
	}
	for rows.Next() {
		buffer := ModelUser{}
		err := rows.Scan(
			&buffer.ID,
			&buffer.Email,
			&buffer.Username,
			&buffer.Created,
		)
		if err != nil {
			return nil, err
		}
		ret.Users = append(ret.Users, buffer)
	}

	return &ret, nil

}
