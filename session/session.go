package session

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/noteshare/config"
)

//
// Session contains all session claim information.
//
type Session struct {
	AccountID uint64 `json:"aid,omitempty"`
	UserID    uint64 `json:"uid,omitempty"`
	jwt.StandardClaims
}

// Parse parses a request and fills the claim with the accosiated data.
func (s *Session) Parse(jwtString string) error {
	token, err := jwt.ParseWithClaims(jwtString, &Session{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method or invalid string")
		}
		return []byte(config.SessionSecretKey), nil
	})

	if session, ok := token.Claims.(*Session); ok && token.Valid {
		s.ExpiresAt = session.ExpiresAt
		s.AccountID = session.AccountID
		s.UserID = session.UserID
	}

	return err
}

// Stringify stringifies claims.
func (s *Session) Stringify() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, s)
	str, err := token.SignedString([]byte(config.SessionSecretKey))
	return str, err
}
