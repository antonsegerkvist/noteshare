package account

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// GetAccountFromUserID returns the account the spcified user belongs to.
//
func GetAccountFromUserID(userID uint64) (*ModelAccount, error) {

	const query = `
		select a.c_id, a.c_name from t_account as a
		inner join t_user as u
		on a.c_id = u.c_account_id
		where u.c_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	account := ModelAccount{}
	row := stmt.QueryRow(userID)
	err = row.Scan(&account.ID, &account.Name)
	if err == sql.ErrNoRows {
		return nil, ErrInvalidUserOrAccount
	} else if err != nil {
		return nil, err
	}

	return &account, nil

}
