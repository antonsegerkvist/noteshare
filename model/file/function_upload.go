package file

import (
	"database/sql"

	"github.com/noteshare/mysql"
	"github.com/noteshare/session"
)

//
// LookupUploadFile checks if a unprocessed file exists.
//
func LookupUploadFile(fileID uint64, session *session.Session) error {

	const query = `
		select c_id
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
		return err
	}
	defer stmt.Close()

	var tmp uint64
	row := stmt.QueryRow(
		fileID,
		session.AccountID,
		session.UserID,
	)
	err = row.Scan(&tmp)
	if err == sql.ErrNoRows {
		return ErrFileNotFound
	} else if err != nil {
		return err
	}

	return nil

}

//
// MarkFileAsUploaded marks the file as uploaded.
//
func MarkFileAsUploaded(fileID uint64, session *session.Session) error {

	const query = `
		update t_file
		set c_is_uploaded = 1
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
		fileID,
		session.AccountID,
		session.UserID,
	)
	if err != nil {
		return err
	}

	return nil

}
