package upload

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"hash/adler32"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/mysql"
	"github.com/noteshare/session"
)

//
// LookupUploadFile checks if a unprocessed file exists.
//
func LookupUploadFile(
	fileID uint64,
	accountID uint64,
	userID uint64,
) (uint64, uint32, error) {

	const query = `
		select c_filesize, c_checksum
		from t_file_upload
		where c_id = ?
		and c_account_id = ?
		and c_user_id = ?
		and c_is_uploaded = 0
		and c_is_processed = 0
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, 0, err
	}
	defer stmt.Close()

	var filesize uint64
	var checksum uint32
	row := stmt.QueryRow(
		fileID,
		accountID,
		userID,
	)
	err = row.Scan(&filesize, &checksum)
	if err == sql.ErrNoRows {
		return 0, 0, ErrFileNotFound
	} else if err != nil {
		return 0, 0, err
	}

	return filesize, checksum, nil

}

//
// MarkFileAsUploaded marks the file as uploaded.
//
func MarkFileAsUploaded(
	fileID uint64,
	filesize uint64,
	checksum uint32,
	accountID uint64,
	userID uint64,
) error {

	const query = `
		update t_file
		set c_is_uploaded = 1,
		c_filesize = ?,
		c_checksum = ?
		where c_id = ?
		and c_account_id = ?
		and c_modified_by_user_id = ?
		and c_is_uploaded = 0
		and c_is_processed = 0
	`

	db := mysql.Open()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		filesize,
		checksum,
		fileID,
		accountID,
		userID,
	)
	if err != nil {
		return err
	}

	return nil

}

//
// ResponseData contains the fields of a response.
//
type ResponseData struct {
	Filesize uint64 `json:"filesize"`
	Checksum uint32 `json:"checksum"`
}

//
// Post handles file uploads.
//
var Post = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		p httprouter.Params,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> POST: /service/file/v1/upload/:fid`)
		}

		fileID, err := strconv.ParseUint(p.ByName("fid"), 10, 64)
		if err != nil {
			log.NotifyError(err, http.StatusBadRequest)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		filesize, checksum, err := LookupUploadFile(fileID, s.AccountID, s.UserID)
		if err == ErrFileNotFound {
			log.NotifyError(err, http.StatusNotFound)
			log.RespondJSON(w, `{}`, http.StatusNotFound)
			return
		} else if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		fds, err := os.Create(
			path.Join(
				config.FileRootDir,
				config.FileOriginalDir,
				strconv.FormatUint(fileID, 10),
			),
		)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}
		defer fds.Close()

		filesizeOK, err := io.Copy(fds, r.Body)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		_, err = fds.Seek(0, 0)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		adler32Hash := adler32.New()
		_, err = io.Copy(adler32Hash, fds)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}
		checksumOK := adler32Hash.Sum32()

		if filesize != 0 && filesize != uint64(filesizeOK) {
			log.NotifyError(
				errors.New(`Filesize missmatch`),
				http.StatusBadRequest,
			)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		if checksum != 0 && checksum != uint32(checksumOK) {
			log.NotifyError(
				errors.New(`Checksum missmatch`),
				http.StatusBadRequest,
			)
			log.RespondJSON(w, `{}`, http.StatusBadRequest)
			return
		}

		err = MarkFileAsUploaded(
			fileID,
			uint64(filesizeOK),
			uint32(checksumOK),
			s.AccountID,
			s.UserID,
		)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		responseData := ResponseData{
			Filesize: uint64(filesizeOK),
			Checksum: checksumOK,
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
