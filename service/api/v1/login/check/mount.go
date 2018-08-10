package check

import "github.com/julienschmidt/httprouter"

//
// Mount takes a router object and assigns the login check handler to it.
//
func Mount(router *httprouter.Router) {
	router.OPTIONS(`/service/api/v1/login/check`, Options)
	router.POST(`/service/api/v1/login/check`, Post)
}
