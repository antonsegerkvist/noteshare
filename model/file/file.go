package file

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// NotifyFileUpload saves a notification of a file upload to the database and
// generates a file upload id.
//
func NotifyFileUpload(
	name,
	filename string,
	folderID,
	filesize uint64,
	checksum uint32,
	userID,
	accountID uint64,
) (uint64, error) {

	const query = `
		insert into t_file_upload
		(c_account_id, c_folder_id, c_name, c_filename, c_filesize, c_checksum, c_user_id)
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
		folderID,
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

//
// LookupUploadFile checks if a unprocessed file exists.
//
func LookupUploadFile(
	fileID,
	userID,
	accountID uint64,
) (uint64, uint32, error) {

	const query = `
		select c_filesize, c_checksum
		from t_file_upload
		where c_id = ?
		and c_account_id = ?
		and c_user_id = ?
		and c_is_uploaded = 0
		and c_is_processed = 0
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, 0, err
	}
	defer stmt.Close()

	var filesize uint64
	var checksum uint32
	row := stmt.QueryRow(
		fileID,
		accountID,
		userID,
	)
	err = row.Scan(&filesize, &checksum)
	if err == sql.ErrNoRows {
		return 0, 0, ErrFileNotFound
	} else if err != nil {
		return 0, 0, err
	}

	return filesize, checksum, nil

}

//
// MarkFileAsUploaded marks the file as uploaded.
//
func MarkFileAsUploaded(
	fileID,
	filesize uint64,
	checksum uint32,
	userID,
	accountID uint64,
) error {

	const query = `
		insert into t_file (
			c_account_id,
			c_folder_id,
			c_has_preview,
			c_type,
			c_name,
			c_filename,
			c_filesize,
			c_checksum,
			c_created_by_user_id,
			c_modified_by_user_id,
		)
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		filesize,
		checksum,
		fileID,
		accountID,
		userID,
	)
	if err != nil {
		return err
	}

	return nil

}
