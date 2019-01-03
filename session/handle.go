package session

import (
	"fmt"
	"net/http"

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
	Session,
)

//
// Authenticate is a handle for authenticating the user. The provided handle
// is called if the user is authenticated.
//
func Authenticate(s Handle) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		s(w, r, session)
	}
}

//
// FileSystemAuthenticate is a handler that checks the user authentication
// token before letting the the handle serve the actual content.
//
func FileSystemAuthenticate(s http.HandlerFunc, redirectURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(config.SessionCookieName)
		if err != nil {
			http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
			return
		}

		token := cookie.Value
		session := Session{}
		err = session.Parse(token)
		if err != nil {

			refreshCookie, err := r.Cookie(config.SessionRefreshCookieName)
			if err != nil {
				http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
				return
			}

			refreshToken := refreshCookie.Value
			session := Session{}
			err = session.Parse(refreshToken)
			if err != nil {
				http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
				return
			}

			err = Renewer(w, r)
			if err == ErrUnauthorized {
				http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
				return
			} else if err != nil {
				log.NotifyError(err, http.StatusInternalServerError)
				w.Header().Set("Content-Type", "text/html")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `Internal server error`)
				return
			}

		}

		s(w, r)
	}
}
