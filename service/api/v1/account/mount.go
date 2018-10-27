package account

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the account handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/account/me", Get)
	router.OPTIONS("/service/api/v1/account/me", Options)
}
