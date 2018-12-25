package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/model/user"
	"github.com/noteshare/session"
)

//
// PostData contains the fields that the post body can contain.
//
type PostData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//
// ParseRequestBody parses the request body and fills the fields in
// the struct with it.
//
func (post *PostData) ParseRequestBody(request *http.Request) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(post)
	if err != nil {
		return errors.New("Could not parse the request body")
	}
	request.Body.Close()
	return nil
}

//
// Post handles login requests.
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

	var postData PostData
	err := postData.ParseRequestBody(r)
	if err != nil {
		log.NotifyError(err, http.StatusBadRequest)
		log.RespondJSON(w, `{}`, http.StatusBadRequest)
		return
	}

	email, password := postData.Email, postData.Password
	loginData, err := user.PerformLogin(email, password)
	if err == user.ErrShortEmail {
		log.NotifyError(err, http.StatusPreconditionFailed)
		log.RespondJSON(w, `{}`, http.StatusPreconditionFailed)
		return
	} else if err == user.ErrLongEmail {
		log.NotifyError(err, http.StatusPreconditionFailed)
		log.RespondJSON(w, `{}`, http.StatusPreconditionFailed)
		return
	} else if err == user.ErrShortPassword {
		log.NotifyError(err, http.StatusPreconditionFailed)
		log.RespondJSON(w, `{}`, http.StatusPreconditionFailed)
		return
	} else if err == user.ErrUserNotFound {
		log.NotifyError(err, http.StatusNotFound)
		log.RespondJSON(w, `{}`, http.StatusNotFound)
		return
	} else if err != nil {
		log.NotifyError(err, http.StatusInternalServerError)
		log.RespondJSON(w, `{}`, http.StatusInternalServerError)
		return
	}

	claims := session.Session{
		AccountID: loginData.AccountID,
		UserID:    loginData.UserID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(config.SessionTime),
			Issuer:    "noteshare",
		},
	}

	refresh := session.Session{
		AccountID: loginData.AccountID,
		UserID:    loginData.UserID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(config.RefreshTime),
			Issuer:    "noteshare",
		},
	}

	token, err := claims.Stringify()
	if err != nil {
		log.NotifyError(err, http.StatusInternalServerError)
		log.RespondJSON(w, `{}`, http.StatusInternalServerError)
		return
	}

	refreshToken, err := refresh.Stringify()
	if err != nil {
		log.NotifyError(err, http.StatusInternalServerError)
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
