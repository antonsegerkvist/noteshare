package config

const (

	//
	// SessionCookieName is the name of the cookie containing the session
	// information.
	//
	SessionCookieName string = "nst"

	//
	// SessionRefreshCookieName is the name of the cookie containing the refresh
	// token information.
	//
	SessionRefreshCookieName string = "nsrt"

	//
	// SessionDomain is the domain the cookie is issued under.
	//
	SessionDomain string = ""

	//
	// SessionPath is the path the cookie is issued under.
	//
	SessionPath string = "/"

	//
	// SessionTime is the number of seconds the cookie is valid for.
	//
	SessionTime uint64 = 3600

	//
	// RefreshTime is the number of seconds the refresh token is valid for.
	//
	RefreshTime uint64 = 43200

	//
	// SessionSecretKey is the secret session key all tokens are encrypted using.
	//
	SessionSecretKey string = "CB6345FAD4DBF5266197B27A11E8371782C0E62791B6A861766EDAEDB1AAB031"
)
