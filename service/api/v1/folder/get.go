package folder

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/mysql"
	"github.com/noteshare/session"
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
// GetFolderChildrenFromFolderID returns a list of folder with the specified folder
// id as a parent.
//
func GetFolderChildrenFromFolderID(
	folderID,
	userID,
	accountID uint64,
) (*[]ModelFolder, error) {

	const query = `
		select c_id,
			c_parent,
			c_name,
			c_created_by_user_id,
			c_modified_by_user_id,
			c_created_date,
			c_modified_date
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
// Get returns the folders that are children to the specified folders that the
// user has access to.
//
var Get = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		p httprouter.Params,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: /service/api/v1/folder/:id`)
		}

		folderID, err := strconv.ParseUint(p.ByName("id"), 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		folders, err := GetFolderChildrenFromFolderID(
			folderID,
			s.UserID,
			s.AccountID,
		)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		jsonBody, err := json.Marshal(folders)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBody), http.StatusOK)

	},
)
