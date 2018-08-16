package id

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
// ResponseData contains the fields of the response.
//
type ResponseData struct {
	Folders []folder.ModelFolder `json:"folders"`
}

//
// Get returns the child folders of folderID that the user has access to.
//
var Get = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, p httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: /service/api/v1/folder/by/parent/id/:fid`)
		}

		userID, err := strconv.ParseUint(s.Id, 10, 64)
		if err != nil {
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		folderID, err := strconv.ParseUint(p.ByName("fid"), 10, 64)
		if err != nil {
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		folders, err := folder.GetChildFoldersForFolderIDFromUserID(folderID, userID)
		if err != nil {
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		responseData := ResponseData{Folders: *folders}
		jsonBody, err := json.Marshal(responseData)
		if err != nil {
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBody), http.StatusOK)

	},
)
