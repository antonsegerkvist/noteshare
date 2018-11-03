package permission

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/log"
	"github.com/noteshare/session"
)

//
// Get handles fetching of specific user permissions.
//
var Get = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, p httprouter.Params, s session.Session) {

		log.RespondJSON(w, `{}`, http.StatusOK)

	},
)
