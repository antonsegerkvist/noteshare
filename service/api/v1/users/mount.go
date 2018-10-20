package users

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the users handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/users", Get)
	router.OPTIONS("/service/api/v1/users", Options)
}
