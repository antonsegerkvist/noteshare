package check

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/session"
)

//
// Post checks if the users jwt is still valid.
//
var Post = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		_ session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: ` + r.URL.Path)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			log.NotifyError(errors.New(`Unsupported media-type`), http.StatusUnsupportedMediaType)
			log.RespondJSON(w, `{}`, http.StatusUnsupportedMediaType)
			return
		}

		log.RespondJSON(w, `{}`, http.StatusOK)

	},
)
