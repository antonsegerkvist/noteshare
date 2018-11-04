package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"unicode/utf8"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/model/file"
	"github.com/noteshare/session"
)

//
// PostRequestData contains the fields needed to create a new file.
//
type PostRequestData struct {
	Name     string `json:"name"`
	Filename string `json:"filename"`
	ToFolder uint64 `json:"toFolder"`
	Filesize uint64 `json:"filesize"`
	Checksum uint32 `json:"checksum"`
}

//
// ParseRequest parses the request into the fields.
//
func (rd *PostRequestData) ParseRequest(reader io.ReadCloser) error {
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(rd)
	return err
}

//
// PostResponseData contains the fields needed to upload a new file.
//
type PostResponseData struct {
	FileID uint64 `json:"fileID"`
}

//
// Post adds a single folder as a child to the specified folder id.
//
var Post = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		p httprouter.Params,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: /service/api/v1/file`)
		}

		requestData := PostRequestData{}
		err := requestData.ParseRequest(r.Body)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		if utf8.RuneCountInString(requestData.Name) == 0 {
			log.NotifyError(
				errors.New(`Missing JSON parameter: name`),
				http.StatusBadRequest,
			)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}
		if utf8.RuneCountInString(requestData.Filename) == 0 {
			log.NotifyError(
				errors.New(`Missing JSON parameter: filename`),
				http.StatusBadRequest,
			)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}
		if requestData.ToFolder == 0 {
			log.NotifyError(
				errors.New(`Missing JSON parameter: toFolder`),
				http.StatusBadRequest,
			)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		responseData := PostResponseData{}
		responseData.FileID, err = file.PostFile(
			requestData.Name,
			requestData.Filename,
			requestData.ToFolder,
			requestData.Filesize,
			requestData.Checksum,
			&s,
		)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(responseData)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBytes), http.StatusOK)

	},
)
