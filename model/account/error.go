package account

import "errors"

//
// ErrAccountNotFound is returned either when a account does not exist.
//
var ErrAccountNotFound = errors.New(`Account does not exist`)

//
// ErrUserNotAllowed is returned when a user does not have permission to
// perform a certail action.
//
var ErrUserNotAllowed = errors.New(`User not allowed`)
