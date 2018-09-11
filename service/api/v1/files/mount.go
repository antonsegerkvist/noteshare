package file

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the root files handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/files", Get)
	router.OPTIONS("/service/api/v1/files", Options)
	router.POST("/service/api/v1/files", Post)
}
