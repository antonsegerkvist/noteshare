package folder

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// AddChildFolderForFolderIDFromUserID adds a child folder to the folder
// specified folderID.
//
func AddChildFolderForFolderIDFromUserID(name string, folderID, userID uint64) error {

	const permissionQuery = `
		select c_user_id from t_folder_belongs_to_user as fbtu
		where fbtu.c_user_id = ? and fbtu.c_folder_id = ?
	`

	const insertQuery = `
		insert into t_folder (c_parent, c_name)
		values (?, ?)
	`

	const insertRelationQuery = `
		insert into t_folder_belongs_to_user (c_user_id, c_folder_id)
		values (?, ?)
	`

	db := mysql.Open()

	permissionStmt, err := db.Prepare(permissionQuery)
	if err != nil {
		return err
	}
	defer permissionStmt.Close()

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

	var tmpID uint64
	permissionRow := permissionStmt.QueryRow(userID, folderID)
	err = permissionRow.Scan(&tmpID)
	if err == sql.ErrNoRows {
		return ErrUserNotAllowed
	} else if err != nil {
		return err
	}

	insertResult, err := insertStmt.Exec(folderID, name)
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
// GetChildFoldersForFolderIDFromUserID returns the chid folders of a folder
// denoted by folderID giver the scope of userID.
//
func GetChildFoldersForFolderIDFromUserID(folderID, userID uint64) (*[]ModelFolder, error) {

	const query = `
		select f.c_id, f.c_name
		from t_folder as f
		inner join t_folder_belongs_to_user as fbtu
		on fbtu.c_folder_id = f.c_id
		where fbtu.c_user_id = ?
		and f.c_parent = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID, folderID)
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
