package id

import "github.com/gorilla/mux"

//
// Mount takes a router object and assigns the upload file handler to it.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/file/v1/upload/{fid:[0-9]+}", Post).Methods("POST")
}
