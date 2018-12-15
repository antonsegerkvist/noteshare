package folder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/mysql"
	"github.com/noteshare/session"
)

//
// ModelAddFolder contains information for adding a folder.
//
type ModelAddFolder struct {
	Parent uint64 `json:"parent"`
	Name   string `json:"name"`
}

//
// AddFolderFromFolderData adds a single folder from the spcified folder
// data.
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

//
// PostRequestData contains the fields needed to create a new folder.
//
type PostRequestData struct {
	Name string `json:"name"`
}

//
// ParseRequest parses the request into the fields.
//
func (rd *PostRequestData) ParseRequest(reader io.ReadCloser) error {
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(rd)
	return err
}

//
// Post adds a single folder as a child to the specified folder id.
//
var Post = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		p httprouter.Params,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: /service/api/v1/folder/:id`)
		}

		folderID, err := strconv.ParseUint(p.ByName("id"), 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		requestData := PostRequestData{}
		err = requestData.ParseRequest(r.Body)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		folderData := &ModelAddFolder{
			Name:   requestData.Name,
			Parent: folderID,
		}
		err = AddFolderFromFolderData(folderData, s.UserID, s.AccountID)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, `{}`, http.StatusOK)

	},
)
