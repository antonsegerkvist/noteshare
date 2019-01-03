package rename

import "github.com/gorilla/mux"

//
// Mount mounts the get folder endpoint to the specified router.
//
func Mount(router *mux.Router) {
	router.HandleFunc("/service/api/v1/folder/rename", Patch).Methods("PATCH")
	router.HandleFunc("/service/api/v1/folder/rename", Options).Methods("OPTIONS")
}
