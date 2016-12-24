package butter

import (
	"butter/auth"
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct holds information about a specific endpoint
type Route struct {
	Method      string
	URI         string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type ApplicationRoute struct {
	Name       string
	Method     string
	URI        string
	Func       ApplicationHandleFunc
	Middleware auth.Chain
}

// ApplicationHandleFunc is a function that wraps a handler func with an
// App context so that the handler can access the application core
type ApplicationHandleFunc func(app *App) http.HandlerFunc

type Router interface {
	Methods(...string) Routeable
	ServeHTTP(http.ResponseWriter, *http.Request)
	AddRoutes(routes Routes) Router
}

type Routeable interface {
	Path(string) Routeable
	HandlerFunc(f func(http.ResponseWriter, *http.Request)) Routeable
}

type GorillaRouter struct {
	Router *mux.Router
}

// Sets the methods on the gorilla mux
func (r *GorillaRouter) Methods(methods ...string) Routeable {
	route := r.Router.Methods(methods...)
	return &GorillaRouting{route}
}

// Serve http by defaulting to the gorilla implementation
func (r *GorillaRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Router.ServeHTTP(w, req)
}

// AddRoutes adds a list of routes to the underlying gorilla router
func (r *GorillaRouter) AddRoutes(routes Routes) Router {
	for _, route := range routes {
		r.Router.Methods(route.Method, "OPTIONS").Path(route.URI).HandlerFunc(route.HandlerFunc)
	}

	return r
}

type GorillaRouting struct {
	Route *mux.Route
}

// Sets a Path for the route within Gorilla Implementation.
func (r *GorillaRouting) Path(tpl string) Routeable {
	r.Route.Path(tpl)
	return r
}

// return a pointer to a new gorilla router which is a wrapper
// interface around the concrete mux implementation
func NewGorillaRouter() Router {
	return &GorillaRouter{
		mux.NewRouter(),
	}
}

// Sets a handler for the route within Gorilla Implementation.
func (r *GorillaRouting) HandlerFunc(f func(http.ResponseWriter, *http.Request)) Routeable {
	r.Route.HandlerFunc(f)
	return r
}

// ApplyRoutes returns a list of routes to apply to a router from a given slice of
// application routes
func ApplyRoutes(app *App, appRoutes []ApplicationRoute, middleFunc auth.MiddlewareCallable) Routes {
	var routes Routes

	for _, route := range appRoutes {
		var handler http.HandlerFunc
		handler = route.Func(app)

		// if there is middleware to apply lets do it here
		if len(route.Middleware) > 0 {
			handler = middleFunc(handler, route.Middleware...)
		}

		routes = append(routes, Route{
			Method:      route.Method,
			URI:         route.URI,
			HandlerFunc: handler,
		})
	}

	return routes
}
