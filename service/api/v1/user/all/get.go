package all

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: ` + r.URL.Path)
		}

		users, err := modeluser.GetUsers(s.AccountID)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		jsonBody, err := json.Marshal(users)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBody), http.StatusOK)

	},
)
