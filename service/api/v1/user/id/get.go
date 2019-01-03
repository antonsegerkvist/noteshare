package id

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	modeluser "github.com/noteshare/model/user"
	"github.com/noteshare/session"
)

//
// Get returns the user that the access token is associated with.
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

		userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		user, err := modeluser.GetUser(userID, s.AccountID)
		if err == modeluser.ErrUserNotFound {
			log.NotifyError(err, http.StatusNotFound)
			log.RespondJSON(w, `{}`, http.StatusNotFound)
			return
		} else if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		jsonBody, err := json.Marshal(user)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBody), http.StatusOK)

	},
)
