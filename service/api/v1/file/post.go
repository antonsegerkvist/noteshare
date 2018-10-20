package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/model/file"
	"github.com/noteshare/session"
)

//
// RequestData contains all the expected fields of a request.
//
type RequestData struct {
	file.ModelAddFile
}

//
// ParseRequest parses the request into the structure fields.
//
func (rd *RequestData) ParseRequest(r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(rd)
	if err != nil {
		return err
	}
	return nil
}

//
// PostResponseData contains the fields of a successful response.
//
type PostResponseData struct {
	FileID uint64 `json:"fileID"`
}

//
// Post returns the root folders that the user has access to.
//
var Post = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params, s session.Session) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: /service/api/v1/file`)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			log.NotifyError(errors.New(`Unsupported media-type`), http.StatusUnsupportedMediaType)
			log.RespondJSON(w, `{}`, http.StatusUnsupportedMediaType)
			return
		}

		userID, err := strconv.ParseUint(s.Id, 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusUnauthorized)
			log.RespondJSON(w, `{}`, http.StatusUnauthorized)
			return
		}

		requestData := RequestData{}
		err = requestData.ParseRequest(r)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		fileData := file.ModelAddFile{
			Type:            requestData.Type,
			FileReferenceID: requestData.FileReferenceID,
			Parent:          requestData.Parent,
			Name:            requestData.Name,
		}

		fileID, err := file.AddRootMultiplexerFromUserID(&fileData, userID)
		if err == file.ErrUnknownFileType {
			log.NotifyError(err, http.StatusPreconditionFailed)
			log.RespondJSON(w, `{}`, http.StatusPreconditionFailed)
			return
		} else if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		responseData := PostResponseData{FileID: fileID}
		jsonBytes, err := json.Marshal(responseData)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBytes), http.StatusOK)

	},
)
