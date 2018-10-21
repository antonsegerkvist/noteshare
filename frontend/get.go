package public

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//
// Get handles serving of the frontend resources.
//
func Get(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fs := http.FileServer(http.Dir("frontend/client/dist"))
	http.StripPrefix("/frontend", fs).ServeHTTP(w, r)
}
