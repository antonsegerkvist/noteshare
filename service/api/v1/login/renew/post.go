package renew

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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
		fmt.Println(`==> POST: /service/api/v1/login/renew`)
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.RespondJSON(w, `{}`, http.StatusUnsupportedMediaType)
		return
	}

	oldRefreshCookie, err := r.Cookie(config.SessionRefreshCookieName)
	if err != nil {
		log.RespondJSON(w, `{}`, http.StatusUnauthorized)
		return
	}

	oldRefeshToken := oldRefreshCookie.Value
	oldRefreshSession := session.Session{}
	err = oldRefreshSession.Parse(oldRefeshToken)
	if err != nil {
		log.RespondJSON(w, `{}`, http.StatusUnauthorized)
		return
	}

	claims := session.Session{
		StandardClaims: jwt.StandardClaims{
			Id:        oldRefreshSession.Id,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(config.SessionTime),
			Issuer:    "noteshare",
		},
	}

	refresh := session.Session{
		StandardClaims: jwt.StandardClaims{
			Id:        oldRefreshSession.Id,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(config.RefreshTime),
			Issuer:    "noteshare",
		},
	}

	token, err := claims.Stringify()
	if err != nil {
		log.RespondJSON(w, `{}`, http.StatusInternalServerError)
		return
	}

	refreshToken, err := refresh.Stringify()
	if err != nil {
		log.RespondJSON(w, `{}`, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     config.SessionCookieName,
		Value:    token,
		Domain:   config.SessionDomain,
		Path:     config.SessionPath,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     config.SessionRefreshCookieName,
		Value:    refreshToken,
		Domain:   config.SessionDomain,
		Path:     config.SessionPath,
		HttpOnly: true,
	})

	log.RespondJSON(w, `{}`, http.StatusOK)

}
