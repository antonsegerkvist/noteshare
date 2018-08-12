package folder

import "errors"

//
// ErrShortName is returned when folder name is too short.
//
var ErrShortName = errors.New(`Name too short`)

//
// ErrCouldNotInsert is returned on failed insert.
//
var ErrCouldNotInsert = errors.New(`Could not insert`)
