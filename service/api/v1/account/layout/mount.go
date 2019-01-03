package layout

import "github.com/gorilla/mux"

//
// Mount mounts the account layout handler to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/account/layout", Get).Methods("GET")
	router.HandleFunc("/service/api/v1/account/layout", Options).Methods("OPTIONS")
	router.HandleFunc("/service/api/v1/account/layout", Post).Methods("POST")
}
