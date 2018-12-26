package folder

import (
	"github.com/julienschmidt/httprouter"
)

//
// Mount mounts the get folder endpoint to the specified router.
//
func Mount(router *httprouter.Router) {
	router.DELETE("/service/api/v1/folder/:id", Delete)
	router.GET("/service/api/v1/folder/:id", Get)
	router.OPTIONS("/service/api/v1/folder/:id", Options)
	router.PATCH("/service/api/v1/folder/:id", Patch)
	router.POST("/service/api/v1/folder/:id", Post)
}
