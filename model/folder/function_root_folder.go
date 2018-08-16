package folder

import (
	"database/sql"
	"unicode/utf8"

	"github.com/noteshare/mysql"
)

//
// AddRootFolderFromUserID adds a root folder to the database.
//
func AddRootFolderFromUserID(name string, userID uint64) error {

	const insertQuery = `
		insert into t_folder (c_name)
		values (?)
	`

	const insertRelationQuery = `
		insert into t_folder_belongs_to_user (c_user_id, c_folder_id)
		values (?, ?)
	`

	if utf8.RuneCountInString(name) < 1 {
		return ErrShortName
	}

	db := mysql.Open()

	insertStmt, err := db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	insertRelationStmt, err := db.Prepare(insertRelationQuery)
	if err != nil {
		return err
	}
	defer insertRelationStmt.Close()

	insertResult, err := insertStmt.Exec(name)
	if err != nil {
		return err
	}

	insertRows, err := insertResult.RowsAffected()
	if err != nil {
		return err
	} else if insertRows < 1 {
		return ErrCouldNotInsert
	}

	insertID, err := insertResult.LastInsertId()
	if err != nil {
		return err
	}

	insertRelationResult, err := insertRelationStmt.Exec(userID, insertID)
	if err != nil {
		return err
	}

	insertRelationRows, err := insertRelationResult.RowsAffected()
	if err != nil {
		return err
	} else if insertRelationRows < 1 {
		return ErrCouldNotInsert
	}

	return nil

}

//
// GetRootFoldersFromUserID returns a list of all root folders.
//
func GetRootFoldersFromUserID(userID uint64) (*[]ModelFolder, error) {

	const query = `
		select f.c_id, f.c_name
		from t_folder as f
		inner join t_folder_belongs_to_user as fbtu
		on fbtu.c_folder_id = f.c_id
		where fbtu.c_user_id = ?
		and f.c_parent is null
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err == sql.ErrNoRows {
		return &[]ModelFolder{}, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	folders := []ModelFolder{}
	for rows.Next() {
		buffer := ModelFolder{}
		err = rows.Scan(
			&buffer.ID,
			&buffer.Name,
		)
		if err != nil {
			return nil, err
		}
		folders = append(folders, buffer)
	}

	return &folders, nil

}
