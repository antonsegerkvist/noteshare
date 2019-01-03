package renew

import "github.com/gorilla/mux"

//
// Mount takes a router object and assigns the login renew handler to it.
//
func Mount(router *mux.Router) {
	router.HandleFunc(`/service/api/v1/login/renew`, Options).Methods("OPTIONS")
	router.HandleFunc(`/service/api/v1/login/renew`, Post).Methods("POST")
}
