package layout

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the account layout handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/account/layout", Get)
	router.OPTIONS("/service/api/v1/account/layout", Options)
	router.POST("/service/api/v1/account/layout", Post)
}
