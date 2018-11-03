package permission

import "github.com/julienschmidt/httprouter"

//
// Mount mounts the permission handler to the specified router.
//
func Mount(router *httprouter.Router) {
	router.GET("/service/api/v1/permission", Get)
	router.OPTIONS("/service/api/v1/permission", Options)
}
