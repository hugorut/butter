package butter

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hugorut/butter/auth"
	"github.com/hugorut/butter/data"
	"github.com/hugorut/butter/filesystem"
	"github.com/hugorut/butter/mail"
	"github.com/hugorut/butter/sys"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Serve serves a butter application with http routing
func Serve(routes []ApplicationRoute) (*App, chan error) {
	// Load the environment configuration from the route .env file
	godotenv.Load(".env")

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

	// add routes to the application using the specified routing option
	// routes are specified in the routes.go file in the root of your application
	router := NewGorillaRouter().AddRoutes(ApplyRoutes(app, routes, auth.Middled))

	// make a errors channel if to send the errors to if the http servers fail
	errs := make(chan error)

	if os.Getenv("APP_HTTPS") == "true" {
		go func() {
			app.Logger.Log(sys.INFO, fmt.Sprintf("starting the https service at port :%s", sys.EnvOrDefault("HTTPS_PORT", "5555")))
			err := http.ListenAndServeTLS(
				":"+sys.EnvOrDefault("HTTPS_PORT", "5555"),
				sys.EnvOrDefault("CERT_FILE", "cert.crt"),
				sys.EnvOrDefault("CERT_KEY", "cert.key"),
				router,
			)

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
