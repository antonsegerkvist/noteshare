package session

import "errors"

//
// ErrUnauthorized is returned when a user is unauthorized
//
var ErrUnauthorized = errors.New(`Unauthorized user`)
