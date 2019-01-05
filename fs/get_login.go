package public

import (
	"fmt"
	"net/http"

	"github.com/noteshare/config"
)

//
// GetLogin handles serving of the login resources.
//
func GetLogin(w http.ResponseWriter, r *http.Request) {

	if config.BuildDebug {
		fmt.Println(`==> GET: ` + r.URL.Path)
	}

	fs := http.FileServer(http.Dir("frontend/login/dist"))
	http.StripPrefix("/login", fs).ServeHTTP(w, r)

}
