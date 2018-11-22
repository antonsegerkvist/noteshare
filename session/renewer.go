package session

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/noteshare/config"
)

//
// Renewer handler renewing of old sessions.
//
func Renewer(w http.ResponseWriter, r *http.Request) error {
	oldRefreshCookie, err := r.Cookie(config.SessionRefreshCookieName)
	if err != nil {
		return ErrUnauthorized
	}

	oldRefeshToken := oldRefreshCookie.Value
	oldRefreshSession := Session{}
	err = oldRefreshSession.Parse(oldRefeshToken)
	if err != nil {
		return ErrUnauthorized
	}

	claims := Session{
		AccountID: oldRefreshSession.AccountID,
		UserID:    oldRefreshSession.UserID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(config.SessionTime),
			Issuer:    "noteshare",
		},
	}

	refresh := Session{
		AccountID: oldRefreshSession.AccountID,
		UserID:    oldRefreshSession.UserID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(config.RefreshTime),
			Issuer:    "noteshare",
		},
	}

	token, err := claims.Stringify()
	if err != nil {
		return err
	}

	refreshToken, err := refresh.Stringify()
	if err != nil {
		return err
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

	return nil
}
