package file

import (
	"github.com/noteshare/mysql"
	"github.com/noteshare/session"
)

//
// PostFile saves a notification of a file upload to the database.
//
func PostFile(
	name string,
	toFolder uint64,
	filesize uint64,
	checksum uint64,
	session *session.Session,
) (uint64, error) {

	const query = `
		insert into t_file
		(c_account_id, c_type, c_filename, c_name, c_filesize, c_checksum, c_created_by_user_id, c_modified_by_user_id)
		values (?, 0, '', ?, ?, ?, ?, ?)
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		session.AccountID,
		name,
		filesize,
		checksum,
		session.UserID,
		session.UserID,
	)
	if err != nil {
		return 0, err
	}

	fileID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(fileID), nil

}
