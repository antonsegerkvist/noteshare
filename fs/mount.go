package public

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noteshare/session"
)

//
// Mount mounts the public handler to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/frontend/", http.StatusTemporaryRedirect)
	}).Methods("GET")
	router.PathPrefix("/login").HandlerFunc(GetLogin).Methods("GET")
	router.PathPrefix("/frontend").HandlerFunc(session.FileSystemAuthenticate(GetClient, "/login/")).Methods("GET")
}
