package upload

import "github.com/julienschmidt/httprouter"

//
// Mount takes a router object and assigns the upload file handler to it.
//
func Mount(router *httprouter.Router) {
	router.POST("/service/file/v1/upload/:fid", Post)
}
