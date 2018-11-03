package account

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// GetAccountLayout returns the account layout for the user.
//
func GetAccountLayout(accountID uint64) (*ModelAccountLayout, error) {

	const query = `
		select c_layout from t_account
		where c_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var json sql.NullString
	row := stmt.QueryRow(accountID)
	err = row.Scan(&json)
	if err == sql.ErrNoRows {
		return nil, ErrAccountNotFound
	} else if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret := ModelAccountLayout{}
	if json.Valid {
		ret.Layout = &json.String
	} else {
		ret.Layout = nil
	}

	return &ret, nil

}
