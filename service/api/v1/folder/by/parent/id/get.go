package id

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/session"
)

//
// Get handles getting of folders with a certail parent folder.
//
var Get = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params, s session.Session) {
	},
)
