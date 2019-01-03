package public

import (
	"net/http"
)

//
// GetClient handles serving of the frontend resources.
//
func GetClient(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("frontend/client/dist"))
	http.StripPrefix("/frontend", fs).ServeHTTP(w, r)
}
