package auth

import (
	"butter/database"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// TokenGenerator interface which returns a token to use as an apikey in
// the header of a request.
type TokenGenerator interface {
	GenerateToken(user database.User) string
}

// JWTGenerator is the json web token implementation of a TokenGenerator
type JWTGenerator struct {
	Secret []byte
}

// GenerateToken generates a jwt token from a valid user object
// adds custom claims to the token with the ID of the user
func (g *JWTGenerator) GenerateToken(user database.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
	})

	tokenString, _ := token.SignedString(g.Secret)
	return tokenString
}

// GetSecret returns jwt signing secret from an environment variable
func GetSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
