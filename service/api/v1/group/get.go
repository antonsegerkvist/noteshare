package group

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/session"
)

//
// Get returns a list of all groups in the account associated with the access
// token.
//
var Get = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		p httprouter.Params,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> GET: ` + r.URL.Path)
		}

		switch p.ByName("id") {
		case "me":
			break
		case "all":
			break
		}

	},
)
