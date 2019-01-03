package login

import "github.com/gorilla/mux"

//
// Mount takes a router object and assigns the login handler to it.
//
func Mount(router *mux.Router) {
	router.HandleFunc(`/service/api/v1/login`, Options).Methods("OPTIONS")
	router.HandleFunc(`/service/api/v1/login`, Post).Methods("POST")
}
