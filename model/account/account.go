package account

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// ModelAccount contains information of a single account.
//
type ModelAccount struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

//
// GetAccount returns the account the sepcified user belongs to.
//
func GetAccount(accountID uint64) (*ModelAccount, error) {

	const query = `
		select c_id, c_name from t_account
		where c_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	account := ModelAccount{}
	row := stmt.QueryRow(accountID)
	err = row.Scan(&account.ID, &account.Name)
	if err == sql.ErrNoRows {
		return nil, ErrAccountNotFound
	} else if err != nil {
		return nil, err
	}

	return &account, nil

}
