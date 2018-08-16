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
// RequestData contains the request fields.
//
type RequestData struct {
	Name string `json:"name"`
}

//
// ParseRequest parses the request into the request fields.
//
func (rd *RequestData) ParseRequest(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(rd)
	if err != nil {
		return err
	}
	return nil
}

//
// Post handles adding of new root directories.
//
var Post = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: /service/api/v1/folder`)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			log.RespondJSON(w, `{}`, http.StatusUnsupportedMediaType)
			return
		}

		userID, err := strconv.ParseUint(s.Id, 10, 64)
		if err != nil {
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		requestData := RequestData{}
		err = requestData.ParseRequest(r)
		if err != nil {
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		err = folder.AddRootFolderFromUserID(requestData.Name, userID)
		if err == folder.ErrShortName {
			log.RespondJSON(w, `{}`, http.StatusPreconditionFailed)
			return
		} else if err != nil {
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, `{}`, http.StatusOK)

	},
)
