package layout

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	modellayout "github.com/noteshare/model/account"
	"github.com/noteshare/session"
)

//
// PostRequestData contains the expected fields of the json body.
//
type PostRequestData struct {
	Layout string `json:"layout"`
}

//
// ParseRequest parses the request json body into the structure fields.
//
func (rd *PostRequestData) ParseRequest(reader io.ReadCloser) error {
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(rd)
	if err != nil {
		return err
	}
	return nil
}

//
// Post updates the account layout that the user is a part of.
//
var Post = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		p httprouter.Params,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: ` + r.URL.Path)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			log.NotifyError(errors.New(`Unsupported media-type`), http.StatusUnsupportedMediaType)
			log.RespondJSON(w, `{}`, http.StatusUnsupportedMediaType)
			return
		}

		requestData := PostRequestData{}
		err := requestData.ParseRequest(r.Body)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		err = modellayout.UpdateAccountLayout(requestData.Layout, s.AccountID)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, `{}`, http.StatusOK)
	},
)
