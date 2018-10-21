package user

//
// ModelLogin contains the data required during login.
//
type ModelLogin struct {
	UserID uint64
}

//
// ModelUser contains user information.
//
type ModelUser struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Created  string `json:"created"`
}
