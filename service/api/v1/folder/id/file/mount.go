package file

import "github.com/gorilla/mux"

//
// Mount mounts the get folder files endpoint to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/folder/{id:[0-9]+}/file", Get).Methods("GET")
	router.HandleFunc("/service/api/v1/folder/{id:[0-9]+}/file", Options).Methods("OPTIONS")
}
