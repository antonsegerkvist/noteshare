package file

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// GetRootFilesFromUserID returns a list of all root files.
//
func GetRootFilesFromUserID(userID uint64) (*[]ModelFile, error) {

	const query = `
		select f.c_id, f.c_type, f.c_name
		from t_file as f
		inner join t_file_belongs_to_user as fbtu
		on fbtu.c_file_id = f.c_id
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
		return &[]ModelFile{}, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := []ModelFile{}
	for rows.Next() {
		buffer := ModelFile{}
		err = rows.Scan(
			&buffer.ID,
			&buffer.Type,
			&buffer.Name,
		)
		if err != nil {
			return nil, err
		}
		files = append(files, buffer)
	}

	return &files, nil

}
