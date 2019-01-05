package public

import (
	"fmt"
	"net/http"

	"github.com/noteshare/config"
)

//
// GetClient handles serving of the frontend resources.
//
func GetClient(w http.ResponseWriter, r *http.Request) {

	if config.BuildDebug {
		fmt.Println(`==> GET: ` + r.URL.Path)
	}

	fs := http.FileServer(http.Dir("frontend/client/dist"))
	http.StripPrefix("/frontend", fs).ServeHTTP(w, r)

}
