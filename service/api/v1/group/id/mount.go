package id

import "github.com/gorilla/mux"

//
// Mount mounts the group endpoints to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/group/{id:[0-9]+}", Get).Methods("GET")
	router.HandleFunc("/service/api/v1/group/{id:[0-9]+}", Options).Methods("OPTIONS")
}
