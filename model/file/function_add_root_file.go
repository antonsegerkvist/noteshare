package file

import (
	"unicode/utf8"

	"github.com/noteshare/mysql"
)

//
// AddRootFileFromUserID adds a single root file to the users personal folder.
//
func AddRootFileFromUserID(file *ModelAddFile, userID uint64) (uint64, error) {

	const insertQuery = `
		insert into t_file (c_type, c_name, c_is_processed, c_is_uploaded, c_modified_by_user_id)
		values (?, ?, 0, 0, ?)
	`

	const insertRelationQuery = `
		insert into t_file_belongs_to_user (c_user_id, c_file_id)
		values (?, ?)
	`

	if utf8.RuneCountInString(file.Name) < 1 {
		return 0, ErrShortName
	}

	db := mysql.Open()

	insertStmt, err := db.Prepare(insertQuery)
	if err != nil {
		return 0, err
	}
	defer insertStmt.Close()

	insertRelationStmt, err := db.Prepare(insertRelationQuery)
	if err != nil {
		return 0, err
	}
	defer insertRelationStmt.Close()

	insertResult, err := insertStmt.Exec(TypeFile, file.Name, userID)
	if err != nil {
		return 0, err
	}

	insertRows, err := insertResult.RowsAffected()
	if err != nil {
		return 0, err
	} else if insertRows < 1 {
		return 0, ErrCouldNotInsert
	}

	insertID, err := insertResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	insertRelationResult, err := insertRelationStmt.Exec(userID, insertID)
	if err != nil {
		return 0, err
	}

	insertRelationRows, err := insertRelationResult.RowsAffected()
	if err != nil {
		return 0, err
	} else if insertRelationRows < 1 {
		return 0, ErrCouldNotInsert
	}

	return uint64(insertID), nil

}
