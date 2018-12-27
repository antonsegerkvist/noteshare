package all

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/log"
)

//
// Get fetches a list of all groups in an account.
//
func Get(
	w http.ResponseWriter,
	r *http.Request,
	p httprouter.Params,
) {

	log.RespondJSON(w, `{}`, http.StatusOK)

}
