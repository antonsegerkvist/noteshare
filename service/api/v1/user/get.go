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
// GetResponseData contains the fields of the response.
//
type GetResponseData struct {
	User *user.ModelUser `json:"user"`
}

//
// Get returns the users that are in the same account.
//
var Get = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, p httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: /service/api/v1/user/:id`)
		}

		userID, err := strconv.ParseUint(s.Id, 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusUnauthorized)
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		targetUserID := userID
		targetUser := p.ByName("id")
		if targetUser != "me" {
			targetUserID, err = strconv.ParseUint(s.Id, 10, 64)
			if err != nil {
				log.NotifyError(err, http.StatusBadRequest)
				log.RespondJSON(w, `{}`, http.StatusBadRequest)
				return
			}
		}

		response := GetResponseData{}
		response.User, err = user.GetUser(targetUserID, userID)
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
