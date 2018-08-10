package renew

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noteshare/config"
	"github.com/noteshare/log"
)

//
// Options simply tells the client which methods, headers and origins
// are allowed.
//
func Options(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if config.BuildDebug == true {
		fmt.Println(`==> OPTIONS: /service/api/v1/login/renew`)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin")
	log.RespondJSON(w, ``, http.StatusOK)
}
