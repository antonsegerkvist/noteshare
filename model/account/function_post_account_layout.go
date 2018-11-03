package account

import "github.com/noteshare/mysql"

//
// PostAccountLayout updates the account layout provided the user has the
// right to do so.
//
func PostAccountLayout(userID, accountID uint64) error {

	const updateQuery = `
		update t_account set c_layout where c_id = ?
	`

	db := mysql.Open()

	updateStmt, err := db.Prepare(updateQuery)
	if err != nil {
		return err
	}
	defer updateStmt.Close()

	_, err = updateStmt.Exec(accountID)
	if err != nil {
		return err
	}

	return nil

}
