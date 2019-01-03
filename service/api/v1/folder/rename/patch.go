package rename

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/noteshare/log"
	modelfolder "github.com/noteshare/model/folder"
	"github.com/noteshare/session"
)

//
// PatchRequestData contains the fields required to rename a folder.
//
type PatchRequestData struct {
	FolderID uint64 `json:"folderID"`
	Name     string `json:"name"`
}

//
// ParseRequest parses the request body into the structure fields.
//
func (rd *PatchRequestData) ParseRequest(reader io.ReadCloser) error {
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(rd)
	return err
}

//
// Patch handles renaming a single folder.
//
var Patch = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		s session.Session,
	) {

		if r.Header.Get("Content-Type") != "application/json" {
			log.NotifyError(errors.New(`Unsupported media-type`), http.StatusUnsupportedMediaType)
			log.RespondJSON(w, `{}`, http.StatusUnsupportedMediaType)
			return
		}

		requestData := PatchRequestData{}
		err := requestData.ParseRequest(r.Body)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		err = modelfolder.RenameFolder(
			requestData.Name,
			requestData.FolderID,
			s.UserID,
			s.AccountID,
		)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, `{}`, http.StatusOK)

	},
)
