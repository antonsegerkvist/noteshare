package user

import (
	"database/sql"
	"unicode/utf8"

	"github.com/noteshare/mysql"
)

//
// PerformLogin checks if the email and password exists in the database and returns
// the login model data accosiated with it.
//
func PerformLogin(email, password string) (*ModelLogin, error) {
	const query = `
		select c_id from t_user
		where c_email = ? and c_password_hash = SHA2(CONCAT(c_password_salt, ?), 256)
	`

	if utf8.RuneCountInString(email) < 3 {
		return nil, ErrShortEmail
	} else if utf8.RuneCountInString(email) > 320 {
		return nil, ErrLongEmail
	} else if utf8.RuneCountInString(password) < 5 {
		return nil, ErrShortPassword
	}

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	ret := &ModelLogin{}
	row := stmt.QueryRow(email, password)
	err = row.Scan(&ret.UserID)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return ret, nil
}
