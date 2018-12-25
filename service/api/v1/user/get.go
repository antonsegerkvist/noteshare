package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	modeluser "github.com/noteshare/model/user"
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
			fmt.Println(`==> GET: ` + r.URL.Path)
		}

		switch p.ByName("id") {
		case "all":
			response, err := modeluser.GetUsers(s.AccountID)
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
			break

		case "me":
			response, err := modeluser.GetUser(s.UserID, s.AccountID)
			if err == modeluser.ErrUserNotFound {
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
			break

		default:
			userID, err := strconv.ParseUint(p.ByName("id"), 10, 64)
			if err != nil {
				log.NotifyError(err, http.StatusBadRequest)
				log.RespondJSON(w, `{}`, http.StatusBadRequest)
				return
			}

			response, err := modeluser.GetUser(userID, s.AccountID)
			if err == modeluser.ErrUserNotFound {
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
			break
		}

	},
)
