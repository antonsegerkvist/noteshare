package id

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/session"
)

var Post = session.Authenticate(
	func(w http.ResponseWriter, r *http.Request, _ httprouter.Params, s session.Session) {
	},
)
