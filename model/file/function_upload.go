package file

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// LookupUnprocessedFileFromUserID checks if a unprocessed file exists.
//
func LookupUnprocessedFileFromUserID(fileID, userID uint64) error {

	const query = `
		select c_id
		from t_file
		where c_id = ?
		and c_is_processed = 0
		and c_modified_by_user_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var tmp uint64
	row := stmt.QueryRow(fileID, userID)
	err = row.Scan(&tmp)
	if err == sql.ErrNoRows {
		return ErrFileNotFound
	} else if err != nil {
		return err
	}

	return nil

}

//
// MarkFileAsUploadedFromUserID marks the file as uploaded.
//
func MarkFileAsUploadedFromUserID(fileID uint64) error {

	const query = `
		update t_file
		set c_is_uploaded = 1
		where c_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(fileID)
	if err != nil {
		return err
	}

	return nil

}
