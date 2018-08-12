package folder

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the folder handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/folder", Get)
	router.OPTIONS("/service/api/v1/folder", Options)
	router.POST("/service/api/v1/folder", Post)
}
