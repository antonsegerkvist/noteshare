package me

import "github.com/gorilla/mux"

//
// Mount mounts the group endpoints to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/group/me", Get).Methods("GET")
	router.HandleFunc("/service/api/v1/group/me", Options).Methods("OPTIONS")
}
