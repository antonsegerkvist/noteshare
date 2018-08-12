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
// ResponseData contains the fields of the response.
//
type ResponseData struct {
	Folders []folder.ModelFolder `json:"folders"`
}

//
// Get returns the root folders that the user has access to.
//
var Get = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: /service/api/v1/folder`)
		}

		userID, err := strconv.ParseUint(s.Id, 10, 64)
		if err != nil {
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		folders, err := folder.GetFoldersFromUserID(userID)
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
