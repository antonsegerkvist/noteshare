package file

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the root file handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/file", Get)
	router.OPTIONS("/service/api/v1/file", Options)
	router.POST("/service/api/v1/file", Post)
}
