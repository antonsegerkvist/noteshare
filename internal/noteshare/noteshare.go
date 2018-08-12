package noteshare

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	serviceapiv1folder "github.com/noteshare/service/api/v1/folder"
	serviceapiv1login "github.com/noteshare/service/api/v1/login"
	serviceapiv1logincheck "github.com/noteshare/service/api/v1/login/check"
	serviceapiv1loginrenew "github.com/noteshare/service/api/v1/login/renew"
	servicefilev1download "github.com/noteshare/service/file/v1/download"
	servicefilev1upload "github.com/noteshare/service/file/v1/upload"
)

//
// Run handles starting of noteshare server
//
func Run() {

	router := httprouter.New()

	serviceapiv1folder.Mount(router)

	serviceapiv1login.Mount(router)
	serviceapiv1logincheck.Mount(router)
	serviceapiv1loginrenew.Mount(router)

	servicefilev1download.Mount(router)
	servicefilev1upload.Mount(router)

	log.Fatal(http.ListenAndServe(":8081", router))

}
