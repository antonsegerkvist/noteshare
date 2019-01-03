package file

import "github.com/gorilla/mux"

//
// Mount mounts the get folder endpoint to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/file", Options).Methods("OPTIONS")
	router.HandleFunc("/service/api/v1/file", Post).Methods("POST")
}
