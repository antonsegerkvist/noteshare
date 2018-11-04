package folder

import "github.com/noteshare/mysql"

//
// AddFolderFromFolderData returns a list of folder with the specified folder
// id as a parent.
//
func AddFolderFromFolderData(
	folderData *ModelAddFolder,
	userID,
	accountID uint64,
) error {

	const query = `
		insert into t_folder (
			c_account_id, c_parent,
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

	_, err = stmt.Exec(
		accountID,
		folderData.Parent,
		folderData.Name,
		userID,
		userID,
	)
	if err != nil {
		return err
	}

	return nil

}
