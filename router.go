package butter

import (
	"errors"
	"net/http"
	"net/http/pprof"

	"github.com/hugorut/butter/auth"

	"runtime"
	"strings"

	"fmt"

	"os"

	"github.com/gorilla/mux"
	"github.com/hugorut/butter/sys"
)

// debug routes used for profiling the application
var debugRoutes []ApplicationRoute = []ApplicationRoute{
	{Method: "GET", URI: "/debug/pprof/", Func: func(*App) http.HandlerFunc { return http.HandlerFunc(pprof.Index) }},
	{Method: "GET", URI: "/debug/pprof/cmdline", Func: func(*App) http.HandlerFunc { return http.HandlerFunc(pprof.Cmdline) }},
	{Method: "GET", URI: "/debug/pprof/profile", Func: func(*App) http.HandlerFunc { return http.HandlerFunc(pprof.Profile) }},
	{Method: "GET", URI: "/debug/pprof/symbol", Func: func(*App) http.HandlerFunc { return http.HandlerFunc(pprof.Symbol) }},
	{Method: "POST", URI: "/debug/pprof/symbol", Func: func(*App) http.HandlerFunc { return http.HandlerFunc(pprof.Symbol) }},
	{Method: "GET", URI: "/debug/pprof/trace", Func: func(*App) http.HandlerFunc { return http.HandlerFunc(pprof.Trace) }},
}

// Route struct holds information about a specific endpoint
type Route struct {
	Method      string
	URI         string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// ApplicationRoute provides a struct to hold information for a entire route
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

// Router is a core struct which is responsible for serving routes for different http requests
type Router interface {
	Methods(...string) Routeable
	ServeHTTP(http.ResponseWriter, *http.Request)
	AddRoutes(routes Routes) Router
	AddHandler(string, http.Handler) Router
}

// Routeable defines an entity that is able to be routed
type Routeable interface {
	Path(string) Routeable
	HandlerFunc(f func(http.ResponseWriter, *http.Request)) Routeable
}

// GorillaRouter provides a wrapper around the gorilla mux package
type GorillaRouter struct {
	Router *mux.Router
	Logger sys.Logger
}

// AddHandler adds a handler to the mux
func (r *GorillaRouter) AddHandler(path string, handler http.Handler) Router {
	r.Router.Handle(path, handler)

	return r
}

// Sets the methods on the gorilla mux
func (r *GorillaRouter) Methods(methods ...string) Routeable {
	route := r.Router.Methods(methods...)
	return &GorillaRouting{Route: route}
}

// Serve http by defaulting to the gorilla implementation
func (r *GorillaRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error

	// gracefully recover by default
	if os.Getenv("APP_GRACEFUL_RECOVER") != "false" {
		defer func() {
			rec := recover()
			if rec != nil {
				switch t := rec.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}

				file, line := IdentifyPanic()
				r.Logger.Log(
					sys.CRITICAL,
					fmt.Sprintf("Panic recovered from handler\nmethod: %s\nreq: %s\nname: %s\nline: %v\nerr: %s",
						req.Method,
						req.URL.Path,
						file,
						line,
						err.Error()),
				)
				http.Error(w, "Woops, something wen't wrong", http.StatusInternalServerError)
			}
		}()
	}

	r.Router.ServeHTTP(w, req)
}

// IdentifyPanic identifies the line of the panic string
func IdentifyPanic() (string, int) {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	switch {
	case name != "":
		return name, line
	case file != "":
		return file, line
	}

	return "unknown", 0
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

// Path sets a Path for the route within Gorilla Implementation.
func (r *GorillaRouting) Path(tpl string) Routeable {
	r.Route.Path(tpl)
	return r
}

// NewGorillaRouter return a pointer to a new gorilla router which is a wrapper
// interface around the concrete mux implementation
func NewGorillaRouter(logger sys.Logger) Router {
	return &GorillaRouter{
		Router: mux.NewRouter(),
		Logger: logger,
	}
}

// HandlerFunc Sets a handler for the route within Gorilla Implementation.
func (r *GorillaRouting) HandlerFunc(f func(http.ResponseWriter, *http.Request)) Routeable {
	r.Route.HandlerFunc(f)
	return r
}

// ApplyRoutes returns a list of routes to apply to a router from a given slice of application routes
func ApplyRoutes(app *App, appRoutes []ApplicationRoute, middleFunc auth.MiddlewareCallable) Routes {
	var routes Routes

	for _, route := range appRoutes {
		var handler http.HandlerFunc
		handler = route.Func(app)

		// if there is middleware to apply lets do it here
		if route.Middleware != nil && len(route.Middleware) > 0 {
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
