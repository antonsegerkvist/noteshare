package renew

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/session"
)

//
// Post handles renewing of authentication tokens.
//
func Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if config.BuildDebug == true {
		fmt.Println(`==> POST: ` + r.URL.Path)
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.NotifyError(errors.New(`Unsupported media-type`), http.StatusUnsupportedMediaType)
		log.RespondJSON(w, `{}`, http.StatusUnsupportedMediaType)
		return
	}

	err := session.Renewer(w, r)
	if err == session.ErrUnauthorized {
		log.NotifyError(err, http.StatusUnauthorized)
		log.RespondJSON(w, `{}`, http.StatusUnauthorized)
		return
	} else if err != nil {
		log.NotifyError(err, http.StatusUnauthorized)
		log.RespondJSON(w, `{}`, http.StatusInternalServerError)
		return
	}

	log.RespondJSON(w, `{}`, http.StatusOK)

}
