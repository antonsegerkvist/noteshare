package file

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the get folder endpoint to the specified router.
//
func Mount(router *httprouter.Router) {
	router.OPTIONS("/service/api/v1/file", Options)
	router.POST("/service/api/v1/file", Post)
}
