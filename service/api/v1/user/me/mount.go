package me

import "github.com/gorilla/mux"

//
// Mount mounts the me user handler to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/user/me", Get).Methods("GET")
	router.HandleFunc("/service/api/v1/user/me", Options).Methods("OPTIONS")
}
