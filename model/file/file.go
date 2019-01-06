package file

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// ModelFileUpload contains information about a file upload.
//
type ModelFileUpload struct {
	FolderID uint64 `json:"folderID"`
	Name     string `json:"name"`
	Filename string `json:"filename"`
	Filesize uint64 `json:"filesize"`
	Checksum uint32 `json:"checksum"`
}

//
// NotifyFileUpload saves a notification of a file upload to the database and
// generates a file upload id.
//
func NotifyFileUpload(
	folderID uint64,
	name,
	filename string,
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
) (*ModelFileUpload, error) {

	const query = `
		select c_folder_id, c_name, c_filename, c_filesize, c_checksum
		from t_file_upload
		where c_id = ?
		and c_account_id = ?
		and c_user_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ret := ModelFileUpload{}
	row := stmt.QueryRow(
		fileID,
		accountID,
		userID,
	)
	err = row.Scan(
		&ret.FolderID,
		&ret.Name,
		&ret.Filename,
		&ret.Filesize,
		&ret.Checksum,
	)
	if err == sql.ErrNoRows {
		return nil, ErrFileNotFound
	} else if err != nil {
		return nil, err
	}

	return &ret, nil

}

//
// MarkFileAsUploaded marks the file as uploaded.
//
func MarkFileAsUploaded(
	fileID,
	folderID uint64,
	name string,
	filename string,
	filesize uint64,
	checksum uint32,
	userID,
	accountID uint64,
) error {

	const deleteQuery = `
		delete from t_file_upload
		where c_id = ? and c_account_id = ? and c_user_id = ?
	`

	const insertQuery = `
		insert into t_file (
			c_account_id,
			c_folder_id,
			c_name,
			c_filename,
			c_filesize,
			c_checksum,
			c_created_by_user_id,
			c_modified_by_user_id
		) values (?, ?, ?, ?, ?, ?, ?, ?)
	`

	db := mysql.Open()

	deleteStmt, err := db.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer deleteStmt.Close()

	insertStmt, err := db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	_, err = deleteStmt.Exec(
		fileID,
		accountID,
		userID,
	)
	if err != nil {
		return err
	}

	_, err = insertStmt.Exec(
		accountID,
		folderID,
		name,
		filename,
		filesize,
		checksum,
		userID,
		userID,
	)
	if err != nil {
		return err
	}

	return nil

}
