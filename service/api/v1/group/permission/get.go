package permission

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/noteshare/config"
	"github.com/noteshare/log"
	"github.com/noteshare/model/permission"
	"github.com/noteshare/session"
)

//
// Get handles fetching of specific user permissions.
//
var Get = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: ` + r.URL.Path)
		}

		var keys = []uint32{}
		queries := r.URL.Query()
		permissions := queries["key"]

		for _, v := range permissions {
			key, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				log.NotifyError(err, http.StatusBadRequest)
				log.RespondJSON(w, `{}`, http.StatusBadRequest)
				return
			}
			keys = append(keys, uint32(key))
		}

		permissionJSON, err := permission.LookupGroupPermissionJSON(keys, &s)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(*permissionJSON)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBytes), http.StatusOK)

	},
)
