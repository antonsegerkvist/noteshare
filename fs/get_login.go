package public

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//
// GetLogin handles serving of the login resources.
//
func GetLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fs := http.FileServer(http.Dir("frontend/login/dist"))
	http.StripPrefix("/login", fs).ServeHTTP(w, r)
}
