package domain

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email              string `json:"email"`
	jwt.StandardClaims `swaggerignore:"true"`
}
