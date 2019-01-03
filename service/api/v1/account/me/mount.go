package me

import "github.com/gorilla/mux"

//
// Mount mounts the account me handler to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/account/me", Get).Methods("GET")
	router.HandleFunc("/service/api/v1/account/me", Options).Methods("OPTIONS")
}
