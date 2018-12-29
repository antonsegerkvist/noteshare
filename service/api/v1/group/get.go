package group

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	groupall "github.com/noteshare/service/api/v1/group/all"
	groupme "github.com/noteshare/service/api/v1/group/me"
	grouppermission "github.com/noteshare/service/api/v1/group/permission"
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
		case "all":
			groupall.Get(w, r, p, s)
			break
		case "me":
			groupme.Get(w, r, p, s)
			break
		case "permission":
			grouppermission.Get(w, r, p, s)
			break
		}

	},
)
