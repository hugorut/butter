package butter

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/hugorut/butter/data"
	"github.com/hugorut/butter/filesystem"
	"github.com/hugorut/butter/mail"
	"github.com/hugorut/butter/sys"

	"github.com/joho/godotenv"

	"crypto/tls"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/acme/autocert"
)

// Env stores key value pairing
type Env struct {
	Key   string
	Value string
}

// Serve the butter application with manual env
func ServeWithEnv(routes []ApplicationRoute, env []Env) (*App, chan error) {
	for _, e := range env {
		os.Setenv(e.Key, e.Value)
	}

	return serve(routes)
}

// Serve serves a butter application with the given http routes
func Serve(routes []ApplicationRoute, files ...string) (*App, chan error) {
	if len(files) == 0 {
		godotenv.Load(".env")
	}

	for _, f := range files {
		err := godotenv.Load(f)

		if err != nil {
			log.Printf("Could not load file %s, err: %v", f, err)
		}
	}

	return serve(routes)
}

// Serve serves a butter application with the given http routes
func serve(routes []ApplicationRoute) (*App, chan error) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(os.Getenv("HTTPS_WHITELIST")),
		Cache:      autocert.DirCache("certs"),
	}

	// boot up the logging
	logger := sys.NewStdLogger()

	// open the default mysql connection and wrap the connection with a GormORM
	db, err := data.NewMySQLDBConnection()

	if err != nil {
		logger.Log(sys.FATAL, fmt.Sprintf("Could not establish database connection\n error met: %s", err.Error()))
	}

	orm, err := data.WrapSqlInGorm(db)
	if err != nil {
		logger.Log(sys.ERROR, err.Error())
	}

	// create the application with the outputs of the env configuration
	app := &App{
		DB:         db,
		ORM:        orm,
		Store:      data.NewRedisStore(),
		FileSystem: filesystem.NewFileSystem(),
		Mailer:     mail.NewMailer(),
		Time:       new(sys.OSTime),
		Logger:     logger,
	}

	router := NewGorillaRouter(logger)

	// if we wish to profile the application lets append the debug routes to the router
	if os.Getenv("APP_PROFILE") == "true" {
		routes, router = addDebugRoutes(routes, router)
	}

	// add routes to the application using the specified routing option
	// routes are specified in the routes.go file in the root of your application
	router.AddRoutes(ApplyRoutes(app, routes, Middled))

	// make a errors channel if to send the errors to if the http servers fail
	errs := make(chan error)

	if os.Getenv("APP_HTTPS") == "true" {
		go func() {
			app.Logger.Log(sys.INFO, fmt.Sprintf("starting the https service at port :%s", sys.EnvOrDefault("HTTPS_PORT", "5555")))
			server := &http.Server{
				Addr: ":" + sys.EnvOrDefault("HTTPS_PORT", "5555"),
				TLSConfig: &tls.Config{
					GetCertificate: certManager.GetCertificate,
				},
				Handler: router,
			}

			err := server.ListenAndServeTLS("", "")

			if err != nil {
				errs <- err
			}
		}()

	}

	go func() {
		app.Logger.Log(sys.INFO, fmt.Sprintf("starting the http service at port :%s", sys.EnvOrDefault("APP_PORT", "8082")))
		err := http.ListenAndServe(":"+sys.EnvOrDefault("APP_PORT", "8082"), router)

		if err != nil {
			errs <- err
		}
	}()

	// create the server listening as default on 8082 but feel free to change in your .env
	return app, errs
}

func addDebugRoutes(routes []ApplicationRoute, router Router) ([]ApplicationRoute, Router) {
	routes = append(routes, debugRoutes...)
	router.AddHandler("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.AddHandler("/debug/pprof/heap", pprof.Handler("heap"))
	router.AddHandler("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.AddHandler("/debug/pprof/block", pprof.Handler("block"))

	return routes, router
}
