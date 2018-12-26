package folder

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// ModelFolder contains information of a single folder.
//
type ModelFolder struct {
	ID               uint64 `json:"id"`
	Parent           uint64 `json:"parent"`
	Name             string `json:"name"`
	CreatedByUserID  uint64 `json:"createdByUserID"`
	ModifiedByUserID uint64 `json:"modifiedByUserID"`
	CreatedDate      string `json:"createdDate"`
	ModifiedDate     string `json:"modifiedDate"`
}

//
// GetFolderChildren returns a list of folder with the specified folder
// id as a parent.
//
func GetFolderChildren(
	folderID,
	userID,
	accountID uint64,
) (*[]ModelFolder, error) {

	const query = `
		select c_id, c_parent, c_name, c_created_by_user_id, c_modified_by_user_id, c_created_date, c_modified_date
		from t_folder where c_parent = ? and c_account_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(folderID, accountID)
	if err == sql.ErrNoRows {
		return &[]ModelFolder{}, nil
	} else if err != nil {
		return nil, err
	}

	ret := []ModelFolder{}
	for rows.Next() {
		buffer := ModelFolder{}
		err = rows.Scan(
			&buffer.ID,
			&buffer.Parent,
			&buffer.Name,
			&buffer.CreatedByUserID,
			&buffer.ModifiedByUserID,
			&buffer.CreatedDate,
			&buffer.ModifiedDate,
		)
		if err != nil {
			return nil, err
		}
		ret = append(ret, buffer)
	}

	return &ret, nil
}

//
// AddFolder adds a single folder from the spcified folder
// data.
//
func AddFolder(
	name string,
	parent,
	userID,
	accountID uint64,
) error {

	const query = `
		insert into t_folder (
			c_account_id,
			c_parent,
			c_name,
			c_created_by_user_id,
			c_modified_by_user_id
		) values (?, ?, ?, ?, ?)
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		accountID,
		parent,
		name,
		userID,
		userID,
	)
	if err != nil {
		return err
	}

	return nil

}

//
// RenameFolder changes the name of the specified folder.
//
func RenameFolder(
	name string,
	folderID,
	userID,
	accountID uint64,
) error {

	const query = `
		update t_folder set
		c_name = ?,
		c_modified_by_user_id = ?,
		c_modified_date = NOW()
		where c_id = ? and c_account_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		name,
		userID,
		folderID,
		accountID,
	)
	if err != nil {
		return err
	}

	return nil

}

//
// MoveFolder changes a folders parent to the specified parent.
//
func MoveFolder(
	folderID,
	parentID,
	userID,
	accountID uint64,
) error {

	const query = `
		update t_folder set
		c_parent = ?,
		c_modified_by_user_id = ?,
		c_modified_date = NOW()
		where c_id = ? and c_account_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		parentID,
		userID,
		folderID,
		accountID,
	)
	if err != nil {
		return err
	}

	return nil

}

//
// DeleteFolder removes a folder from the specified folder id.
//
func DeleteFolder(
	folderID,
	userID,
	accountID uint64,
) error {

	const query = `
		delete from t_folder
		where c_id = ? and c_account_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		folderID,
		accountID,
	)
	if err != nil {
		return err
	}

	return nil

}
