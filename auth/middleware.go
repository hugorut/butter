package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hugorut/butter"
)

var ErrorJWTAuthError = errors.New("Authentication failed")

// TokenGenerator interface which returns a token to use as an apikey in
// the header of a request.
type TokenGenerator interface {
	GenerateToken(id int) string
}

// JWTGenerator is the json web token implementation of a TokenGenerator
type JWTGenerator struct {
	Secret []byte
}

// GenerateToken generates a jwt token from a valid user object
// adds custom claims to the token with the ID of the entity
// that is is making the request.
func (g *JWTGenerator) GenerateToken(id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
	})

	tokenString, _ := token.SignedString(g.Secret)
	return tokenString
}

// GetSecret returns jwt signing secret from an environment variable
func GetSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// JWTProtected protects a route via a web token, if a valid token is passed then
// we decode the token and set the given entity within the context of the request.
func JWTProtected(next http.HandlerFunc, app *butter.App) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read the token out of the request header
		head := r.Header.Get("Authorization")
		tokenString := strings.TrimSpace(strings.Replace(head, "Bearer ", "", 1))

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			w.WriteHeader(422)
			json.NewEncoder(w).Encode(map[string]string{"error": ErrorJWTAuthError.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// if the claims pass lets set the subject on the request
			ctx := context.WithValue(r.Context(), "sub", claims["sub"])
			r = r.WithContext(ctx)
		} else {
			w.WriteHeader(422)
			json.NewEncoder(w).Encode(map[string]string{"error": ErrorJWTAuthError.Error()})
			return
		}

		next(w, r)
	})
}
