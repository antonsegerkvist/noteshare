package noteshare

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	fs "github.com/noteshare/fs"

	service_api_v1_account_layout "github.com/noteshare/service/api/v1/account/layout"
	service_api_v1_account_me "github.com/noteshare/service/api/v1/account/me"

	service_api_v1_file "github.com/noteshare/service/api/v1/file"

	service_api_v1_folder "github.com/noteshare/service/api/v1/folder"

	service_api_v1_group_permission "github.com/noteshare/service/api/v1/group/permission"

	service_api_v1_login "github.com/noteshare/service/api/v1/login"
	service_api_v1_login_check "github.com/noteshare/service/api/v1/login/check"
	service_api_v1_login_renew "github.com/noteshare/service/api/v1/login/renew"

	service_api_v1_user "github.com/noteshare/service/api/v1/user"

	service_file_v1_download "github.com/noteshare/service/file/v1/download"
	service_file_v1_upload "github.com/noteshare/service/file/v1/upload"
)

//
// Run handles starting of noteshare server
//
func Run() {

	router := httprouter.New()

	fs.Mount(router)

	service_api_v1_account_layout.Mount(router)
	service_api_v1_account_me.Mount(router)

	service_api_v1_file.Mount(router)

	service_api_v1_folder.Mount(router)

	service_api_v1_group_permission.Mount(router)

	service_api_v1_login.Mount(router)
	service_api_v1_login_check.Mount(router)
	service_api_v1_login_renew.Mount(router)

	service_api_v1_user.Mount(router)

	service_file_v1_download.Mount(router)
	service_file_v1_upload.Mount(router)

	log.Fatal(http.ListenAndServe(":8081", router))

}
