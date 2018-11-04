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
	"github.com/noteshare/model/folder"
	"github.com/noteshare/session"
)

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

		folderData := &folder.ModelAddFolder{
			Name:   requestData.Name,
			Parent: folderID,
		}
		err = folder.AddFolderFromFolderData(folderData, s.UserID, s.AccountID)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, `{}`, http.StatusOK)

	},
)
