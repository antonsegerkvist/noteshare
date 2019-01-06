package noteshare

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	fs "github.com/noteshare/fs"

	service_api_v1_account_layout "github.com/noteshare/service/api/v1/account/layout"
	service_api_v1_account_me "github.com/noteshare/service/api/v1/account/me"

	service_api_v1_file "github.com/noteshare/service/api/v1/file"

	service_api_v1_folder_id "github.com/noteshare/service/api/v1/folder/id"
	service_api_v1_folder_id_file "github.com/noteshare/service/api/v1/folder/id/file"
	service_api_v1_folder_move "github.com/noteshare/service/api/v1/folder/move"
	service_api_v1_folder_rename "github.com/noteshare/service/api/v1/folder/rename"

	service_api_v1_group_all "github.com/noteshare/service/api/v1/group/all"
	service_api_v1_group_id "github.com/noteshare/service/api/v1/group/id"
	service_api_v1_group_me "github.com/noteshare/service/api/v1/group/me"
	service_api_v1_group_permission "github.com/noteshare/service/api/v1/group/permission"

	service_api_v1_login "github.com/noteshare/service/api/v1/login"
	service_api_v1_login_check "github.com/noteshare/service/api/v1/login/check"
	service_api_v1_login_renew "github.com/noteshare/service/api/v1/login/renew"

	service_api_v1_user_all "github.com/noteshare/service/api/v1/user/all"
	service_api_v1_user_id "github.com/noteshare/service/api/v1/user/id"
	service_api_v1_user_me "github.com/noteshare/service/api/v1/user/me"

	service_file_v1_download "github.com/noteshare/service/file/v1/download"
	service_file_v1_upload_id "github.com/noteshare/service/file/v1/upload/id"
)

//
// Run handles starting of noteshare server
//
func Run() {

	router := mux.NewRouter()

	fs.Mount(router)

	service_api_v1_account_layout.Mount(router)
	service_api_v1_account_me.Mount(router)

	service_api_v1_file.Mount(router)

	service_api_v1_folder_id.Mount(router)
	service_api_v1_folder_id_file.Mount(router)
	service_api_v1_folder_move.Mount(router)
	service_api_v1_folder_rename.Mount(router)

	service_api_v1_group_all.Mount(router)
	service_api_v1_group_id.Mount(router)
	service_api_v1_group_me.Mount(router)
	service_api_v1_group_permission.Mount(router)

	service_api_v1_login.Mount(router)
	service_api_v1_login_check.Mount(router)
	service_api_v1_login_renew.Mount(router)

	service_api_v1_user_all.Mount(router)
	service_api_v1_user_id.Mount(router)
	service_api_v1_user_me.Mount(router)

	service_file_v1_download.Mount(router)
	service_file_v1_upload_id.Mount(router)

	log.Fatal(http.ListenAndServe(":8081", router))

}
