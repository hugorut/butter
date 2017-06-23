package butter

import (
	"net/http"

	"github.com/hugorut/butter/sys"
)

// Middleware defines a fuction that applies logic to a handler func and
// passes the request to the next Middleware
type Middleware func(http.HandlerFunc, *App) http.HandlerFunc

// Chain is an slice of Middleware
type Chain []Middleware

// MiddlewareCallable calls a chain of Middleware creating a nested tree of functions
type MiddlewareCallable func(final http.HandlerFunc, app *App, chain ...Middleware) http.HandlerFunc

// SkipMiddleware skips all Middleware functions so that a route chain can be tested
func SkipMiddleware(final http.HandlerFunc, app *App, chain ...Middleware) http.HandlerFunc {
	return final
}

// Middled wraps a given handler func with a chain of Middleware
func Middled(final http.HandlerFunc, app *App, chain ...Middleware) http.HandlerFunc {
	handled := chain[len(chain)-1](final, app)

	for i := len(chain) - 2; i >= 0; i-- {
		handled = chain[i](handled, app)
	}

	return handled
}

// CrossOrigin remove access control origin from requests, allow options and requests
// this should be removed or modified for production
func CrossOrigin(next http.HandlerFunc, app *App) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := sys.EnvOrDefault("Allow_Origin", "*")

		w.Header().Set(
			"Access-Control-Allow-Origin",
			origin,
		)
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"POST, GET, OPTIONS, PUT, DELETE",
		)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With",
		)

		// Stop here if its pre-flight OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	})
}
