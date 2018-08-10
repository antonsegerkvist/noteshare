package renew

import "github.com/julienschmidt/httprouter"

//
// Mount takes a router object and assigns the login renew handler to it.
//
func Mount(router *httprouter.Router) {
	router.OPTIONS(`/service/api/v1/login/renew`, Options)
	router.POST(`/service/api/v1/login/renew`, Post)
}
