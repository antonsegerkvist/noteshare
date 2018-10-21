package folder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/model/folder"
	"github.com/noteshare/session"
)

//
// GetResponseData contains the fields of the response.
//
type GetResponseData struct {
	Folders *[]folder.ModelFolder `json:"folders"`
}

//
// Get returns the root folders that the user has access to.
//
var Get = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, p httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: /service/api/v1/folder/:id`)
		}

		userID, err := strconv.ParseUint(s.Id, 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusUnauthorized)
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		folderID, err := strconv.ParseUint(p.ByName("id"), 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		response := GetResponseData{}
		response.Folders, err = folder.GetFoldersFromFolderID(folderID, userID)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		jsonBody, err := json.Marshal(response)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBody), http.StatusOK)

	},
)
