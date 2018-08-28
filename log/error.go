package log

import (
	"fmt"
	"strconv"
)

//
// NotifyError handles error notifications.
//
func NotifyError(err error, errorType int) {
	fmt.Println(``)
	fmt.Println(`==> Error type:    ` + strconv.FormatInt(int64(errorType), 10))
	fmt.Println(`==> Error message: ` + err.Error())
	fmt.Println(``)
}
