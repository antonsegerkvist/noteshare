package file

import (
	"database/sql"

	"github.com/noteshare/mysql"
	"github.com/noteshare/session"
)

//
// LookupUploadFile checks if a unprocessed file exists.
//
func LookupUploadFile(
	fileID uint64,
	session *session.Session,
) (uint64, uint32, error) {

	const query = `
		select c_filesize, c_checksum
		from t_file
		where c_id = ?
		and c_account_id = ?
		and c_modified_by_user_id = ?
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
		session.AccountID,
		session.UserID,
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
	fileID uint64,
	filesize uint64,
	checksum uint32,
	session *session.Session,
) error {

	const query = `
		update t_file
		set c_is_uploaded = 1,
		c_filesize = ?,
		c_checksum = ?
		where c_id = ?
		and c_account_id = ?
		and c_modified_by_user_id = ?
		and c_is_uploaded = 0
		and c_is_processed = 0
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
		session.AccountID,
		session.UserID,
	)
	if err != nil {
		return err
	}

	return nil

}
