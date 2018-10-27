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
// Get returns the folders that are children to the specified folders that the
// user has access to.
//
var Get = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, p httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: /service/api/v1/folder/:id`)
		}

		folderID, err := strconv.ParseUint(p.ByName("id"), 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		folders, err := folder.GetFoldersFromFolderID(folderID, s.UserID, s.AccountID)
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
