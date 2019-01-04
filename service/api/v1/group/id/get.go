package id

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	modelgroup "github.com/noteshare/model/group"
	"github.com/noteshare/session"
)

//
// Get fetches a list of all groups in an account.
//
var Get = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: ` + r.URL.Path)
		}

		groupID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		group, err := modelgroup.GetGroup(
			groupID,
			s.UserID,
			s.AccountID,
		)
		if err == modelgroup.ErrGroupNotFound {
			log.NotifyError(err, http.StatusNotFound)
			log.RespondJSON(w, `{}`, http.StatusNotFound)
			return
		} else if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		jsonBody, err := json.Marshal(group)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBody), http.StatusOK)

	},
)
