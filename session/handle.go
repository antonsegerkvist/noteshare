package session

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
)

//
// Handle is the httprouter handle object with an added parameter
// containing the session claims.
//
type Handle func(
	http.ResponseWriter,
	*http.Request,
	httprouter.Params,
	Session,
)

//
// Authenticate is a handle for authenticating the user. The provided handle
// is called if the user is authenticated.
//
func Authenticate(s Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		cookie, err := r.Cookie(config.SessionCookieName)
		if err != nil {
			log.NotifyError(err, http.StatusUnauthorized)
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		token := cookie.Value
		session := Session{}
		err = session.Parse(token)
		if err != nil {
			log.NotifyError(err, http.StatusUnauthorized)
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		s(w, r, p, session)
	}
}
