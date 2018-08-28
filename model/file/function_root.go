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
		select f.c_id, f.c_type, f.c_file_reference_id, f.c_file_reference_count, f.c_parent, f.c_name
		from t_file as f
		inner join t_file_belongs_to_user as fbtu
		on fbtu.c_file_id = f.c_id
		where fbtu.c_user_id = ?
		and f.c_parent is null and f.c_is_processed = 1
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
		var ID uint64
		var Type uint64
		var FileReferenceID sql.NullInt64
		var FileReferenceCount uint64
		var Parent sql.NullInt64
		var Name string

		err = rows.Scan(
			&ID,
			&Type,
			&FileReferenceID,
			&FileReferenceCount,
			&Parent,
			&Name,
		)
		if err != nil {
			return nil, err
		}

		buffer := ModelFile{
			ID:                 ID,
			Type:               Type,
			FileReferenceID:    0,
			FileReferenceCount: FileReferenceCount,
			Parent:             0,
			Name:               Name,
		}
		if FileReferenceID.Valid {
			buffer.FileReferenceID = uint64(FileReferenceID.Int64)
		}
		if Parent.Valid {
			buffer.Parent = uint64(Parent.Int64)
		}
		files = append(files, buffer)
	}

	return &files, nil

}

//
// AddRootFileFromUserID adds a single root folder from the specified
//
