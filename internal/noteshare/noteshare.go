package noteshare

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	serviceapiv1login "github.com/noteshare/service/api/v1/login"
	serviceapiv1logincheck "github.com/noteshare/service/api/v1/login/check"
	serviceapiv1loginrenew "github.com/noteshare/service/api/v1/login/renew"
)

//
// Run handles starting of noteshare server
//
func Run() {

	router := httprouter.New()

	serviceapiv1login.Mount(router)
	serviceapiv1logincheck.Mount(router)
	serviceapiv1loginrenew.Mount(router)

	log.Fatal(http.ListenAndServe(":8081", router))

}
