package log

import (
	"fmt"
	"net/http"
)

//
// RespondJSON sets the content-type of the header to application/json and
// sets the response code and body.
//
func RespondJSON(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, message)
}
