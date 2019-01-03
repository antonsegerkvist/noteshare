package check

import "github.com/gorilla/mux"

//
// Mount takes a router object and assigns the login check handler to it.
//
func Mount(router *mux.Router) {
	router.HandleFunc(`/service/api/v1/login/check`, Options).Methods("OPTIONS")
	router.HandleFunc(`/service/api/v1/login/check`, Post).Methods("POST")
}
