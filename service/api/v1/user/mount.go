package user

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the user handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/user/:id", Get)
	router.OPTIONS("/service/api/v1/user/:id", Options)
}
