package file

import "github.com/noteshare/mysql"

//
// PostFile saves a notification of a file upload to the database and
// generates a file upload id.
//
func PostFile(
	name string,
	filename string,
	toFolder uint64,
	filesize uint64,
	checksum uint32,
	accountID uint64,
	userID uint64,
) (uint64, error) {

	const query = `
		insert into t_file_upload
		(c_account_id, c_name, c_filename, c_filesize, c_checksum, c_user_id)
		values
		(?, ?, ?, ?, ?, ?, ?)
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		accountID,
		name,
		filename,
		filesize,
		checksum,
		userID,
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
