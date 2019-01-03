package public

import (
	"net/http"
)

//
// GetLogin handles serving of the login resources.
//
func GetLogin(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("frontend/login/dist"))
	http.StripPrefix("/login", fs).ServeHTTP(w, r)
}
