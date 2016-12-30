package butter

import (
	"butter/auth"
	"butter/data"
	"butter/filesystem"
	"butter/mail"
	"butter/sys"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Serve serves a butter application with http routing
func Serve(routes []ApplicationRoute) (*App, error) {
	// Load the environemnt configuration from the route .env file
	godotenv.Load(".env")

	// boot up the logging
	logger := sys.NewStdLogger()

	// open the default mysql connection and wrap the connection with a GormORM
	db, err := data.NewMySQLDBConnection()
	// if there is a database issue then we shouldn't boot the app
	if err != nil {
		logger.Log(sys.FATAL, fmt.Sprintf("Could not establish database connection\n error met: %s", err.Error()))
	}

	orm, err := data.WrapSqlInGorm(db)
	if err != nil {
		logger.Log(sys.ERROR, err.Error())
	}

	defer db.Close()

	// create the application with the outputs of the env configuration
	app := &App{
		DB:         db,
		ORM:        orm,
		Store: 	    data.NewRedisStore(),
		FileSystem: filesystem.NewFileSystem(),
		Mailer:     mail.NewMailer(),
		Time:       new(sys.OSTime),
		Logger:     logger,
	}

	// add routes to the application using the specified routing option
	// routes are specified in the routes.go file in the root of your application
	router := NewGorillaRouter().AddRoutes(ApplyRoutes(app, routes, auth.Middled))

	// create the server listening as default on 8082 but feel free to change in your .env
	return app, http.ListenAndServe(sys.EnvOrDefault("APP_PORT", ":8082"), router)
}
