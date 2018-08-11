package upload

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/session"
)

//
// Post handles file uploads.
//
var Post = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: /service/file/v1/upload`)
		}

		_, err := strconv.ParseUint(s.Id, 10, 64)
		if err != nil {
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		log.RespondJSON(w, `{}`, http.StatusOK)

	},
)
