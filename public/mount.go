package public

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//
// Mount mounts the public handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.Redirect(w, r, "/frontend/", http.StatusTemporaryRedirect)
	})
	router.GET("/frontend/*filepath", Get)
}
