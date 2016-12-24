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
)

// Middleware defines a fuction that applies logic to a handler func and
// passes the request to the next Middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain is an slice of Middleware
type Chain []Middleware

// MiddlewareCallable calls a chain of Middleware creating a nested tree of functions
type MiddlewareCallable func(final http.HandlerFunc, chain ...Middleware) http.HandlerFunc

var ErrorJWTAuthError = errors.New("Authentication failed")

// SkipMiddleware skips all Middleware functions so that a route chain can be tested
func SkipMiddleware(final http.HandlerFunc, chain ...Middleware) http.HandlerFunc {
	return final
}

// wrap a given handler func with a chain of Middleware
func Middled(final http.HandlerFunc, chain ...Middleware) http.HandlerFunc {
	handled := chain[len(chain)-1](final)

	for i := len(chain) - 2; i >= 0; i-- {
		handled = chain[i](handled)
	}

	return handled
}

// remove access control origin from requests, allow options and requests
// this should be removed or modified fro production
func CrossOrigin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")

		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	})
}

// protect a route with via a web token, if a valid token is passed then
// we decode the token and set the given entity within the context of the
// request.
func JWTProtected(next http.HandlerFunc) http.HandlerFunc {
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
