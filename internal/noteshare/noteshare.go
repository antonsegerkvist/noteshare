package noteshare

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	service_api_v1_account "github.com/noteshare/service/api/v1/account"

	service_api_v1_files "github.com/noteshare/service/api/v1/files"

	service_api_v1_login "github.com/noteshare/service/api/v1/login"
	service_api_v1_login_check "github.com/noteshare/service/api/v1/login/check"
	service_api_v1_login_renew "github.com/noteshare/service/api/v1/login/renew"

	service_file_v1_download "github.com/noteshare/service/file/v1/download"
	service_file_v1_upload "github.com/noteshare/service/file/v1/upload"
)

//
// Run handles starting of noteshare server
//
func Run() {

	router := httprouter.New()

	service_api_v1_account.Mount(router)

	service_api_v1_files.Mount(router)

	service_api_v1_login.Mount(router)
	service_api_v1_login_check.Mount(router)
	service_api_v1_login_renew.Mount(router)

	service_file_v1_download.Mount(router)
	service_file_v1_upload.Mount(router)

	log.Fatal(http.ListenAndServe(":8081", router))

}
