package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/model/user"
	"github.com/noteshare/session"
)

//
// Get returns the users that are in the same account.
//
var Get = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		p httprouter.Params,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: /service/api/v1/user/:id`)
		}

		var err error
		targetUserID := s.UserID
		targetUser := p.ByName("id")
		if targetUser != "me" {
			targetUserID, err = strconv.ParseUint(targetUser, 10, 64)
			if err != nil {
				log.NotifyError(err, http.StatusBadRequest)
				log.RespondJSON(w, `{}`, http.StatusBadRequest)
				return
			}
		}

		response, err := user.GetUser(targetUserID, s.UserID)
		if err == user.ErrUserNotFound {
			log.NotifyError(err, http.StatusNotFound)
			log.RespondJSON(w, `{}`, http.StatusNotFound)
			return
		} else if err != nil {
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
