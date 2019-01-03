package move

import "github.com/gorilla/mux"

//
// Mount mounts the get folder endpoint to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/folder/move", Patch).Methods("PATCH")
	router.HandleFunc("/service/api/v1/folder/move", Options).Methods("OPTIONS")
}
