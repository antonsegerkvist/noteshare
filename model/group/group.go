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
	Name string `json:"name"`
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

	groups := ModelGroups{Groups: []ModelGroup{}}
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

	return &groups, nil

}

//
// GetUserGroups returns a list of all groups that a user is part of.
//
func GetUserGroups(forUserID, userID, accountID uint64) (*ModelGroups, error) {

	const query = `
		select g.c_id, g.c_name
		from t_group as g
		inner join t_user_belongs_to_group as ubtg
		on g.c_id = ubtg.c_group_id
		where ubtg.c_user_id = ?
		and g.c_account_id = ?
		and ubtg.c_account_id = ?
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(
		forUserID,
		accountID,
		accountID,
	)
	if err == sql.ErrNoRows {
		return &ModelGroups{}, nil
	} else if err != nil {
		return nil, err
	}

	groups := ModelGroups{Groups: []ModelGroup{}}
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

	return &groups, nil

}
