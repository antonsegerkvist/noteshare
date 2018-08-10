package user

import "errors"

//
// ErrShortEmail is the error returned on short email.
//
var ErrShortEmail = errors.New(`Email too short`)

//
// ErrLongEmail is the error returned on too long email.
//
var ErrLongEmail = errors.New(`Email too long`)

//
// ErrShortPassword is the error returned on short email.
//
var ErrShortPassword = errors.New(`Password too short`)

//
// ErrUserNotFound is the error returned on invalid user information.
//
var ErrUserNotFound = errors.New(`User not found`)
