package folder

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	foldermove "github.com/noteshare/service/api/v1/folder/move"
	folderrename "github.com/noteshare/service/api/v1/folder/rename"
	"github.com/noteshare/session"
)

//
// Patch handles all actions related to modifying a single folder.
//
var Patch = session.Authenticate(
	func(
		w http.ResponseWriter,
		r *http.Request,
		p httprouter.Params,
		s session.Session,
	) {

		if config.BuildDebug == true {
			fmt.Println(`==> PATCH: ` + r.URL.Path)
		}

		switch p.ByName("id") {

		case "move":
			foldermove.Patch(w, r, p, s)
			break

		case "rename":
			folderrename.Patch(w, r, p, s)
			break

		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Not Found")
			break

		}

	},
)
