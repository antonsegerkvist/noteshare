package me

import (
	"encoding/json"
	"net/http"

	"github.com/noteshare/log"
	modelgroups "github.com/noteshare/model/group"
	"github.com/noteshare/session"
)

//
// Get fetches a list of all groups the user associated with the current access
// token is part of.
//
var Get = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		s session.Session,
	) {

		groups, err := modelgroups.GetUserGroups(
			s.UserID,
			s.UserID,
			s.AccountID,
		)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		jsonBody, err := json.Marshal(groups)
		if err != nil {
			log.NotifyError(err, http.StatusInternalServerError)
			log.RespondJSON(w, `{}`, http.StatusInternalServerError)
			return
		}

		log.RespondJSON(w, string(jsonBody), http.StatusOK)

	},
)
