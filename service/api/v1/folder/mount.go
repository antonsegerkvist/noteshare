package folder

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the get folder endpoint to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/folder/:id", Get)
	router.OPTIONS("/service/api/v1/folder/:id", Options)
}
