package id

import "github.com/gorilla/mux"

//
// Mount mounts the get folder endpoint to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/folder/{id:[0-9]+}", Delete).Methods("DELETE")
	router.HandleFunc("/service/api/v1/folder/{id:[0-9]+}", Get).Methods("GET")
	router.HandleFunc("/service/api/v1/folder/{id:[0-9]+}", Options).Methods("OPTIONS")
	router.HandleFunc("/service/api/v1/folder/{id:[0-9]+}", Post).Methods("POST")
}
