package all

import "github.com/gorilla/mux"

//
// Mount mounts the user handler to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/user/all", Get).Methods("GET")
	router.HandleFunc("/service/api/v1/user/all", Options).Methods("OPTIONS")
}
