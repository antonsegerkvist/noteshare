package upload

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/model/file"
	"github.com/noteshare/session"
)

//
// Post handles file uploads.
//
var Post = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, p httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: /service/file/v1/upload/:fid`)
		}

		userID, err := strconv.ParseUint(s.Id, 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusUnauthorized)
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		fileID, err := strconv.ParseUint(p.ByName(":fid"), 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		err = file.LookupUnprocessedFileFromUserID(fileID, userID)
		if err == file.ErrFileNotFound {
			log.NotifyError(err, http.StatusNotFound)
			log.RespondJSON(w, `{}`, http.StatusNotFound)
			return
		} else if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, `{}`, http.StatusOK)

	},
)
