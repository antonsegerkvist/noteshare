package group

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the group endpoints to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/group/:id", Get)
	router.OPTIONS("/service/api/v1/group/:id", Options)
}
