package account

import "errors"

//
// ErrInvalidUserOrAccount is returned either when a user or account does
// not exist.
//
var ErrInvalidUserOrAccount = errors.New(`User or account does not exist`)
