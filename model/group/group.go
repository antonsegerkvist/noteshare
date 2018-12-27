package group

import (
	"database/sql"

	"github.com/noteshare/mysql"
)

//
// ModelGroup contains information of a single group.
//
type ModelGroup struct {
	ID   uint64 `json:"id"`
	Name uint64 `json:"name"`
}

//
// ModelGroups contains information of a collection of groups.
//
type ModelGroups struct {
	Groups []ModelGroup `json:"groups"`
}

//
// GetAllGroups returns a list of all groups in the same account.
//
func GetAllGroups(userID, accountID uint64) (*ModelGroups, error) {

	const query = `
		select c_id, c_name
		from t_group
		where c_account_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(accountID)
	if err == sql.ErrNoRows {
		return &ModelGroups{}, nil
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := ModelGroups{}
	for rows.Next() {
		buffer := ModelGroup{}
		err = rows.Scan(
			&buffer.ID,
			&buffer.Name,
		)
		if err != nil {
			return nil, err
		}
		groups.Groups = append(groups.Groups, buffer)
	}

	return nil, nil

}
