package account

import "errors"

//
// ErrAccountNotFound is returned either when a account does not exist.
//
var ErrAccountNotFound = errors.New(`Account does not exist`)
