package main

import (
	"butter/application"
	"butter/auth"
	"butter/database"
	"butter/routing"
	"butter/service"
	"butter/sys"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/hugorut/gofile"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// global variables defined so as to facilitate testing and mocking
var app *application.App
var db database.ORM

func main() {
	log.Fatal(Serve(routes))
}

func Serve([]routing.ApplicationRoute) error {
	// Load the environemnt configuration from the route .env file
	godotenv.Load(".env")

	// open the default gorm db connection
	db, err := database.OpenGormDbConnection()
	defer db.Close()

	// if there is a database issue then we shouldn't boot the app
	if err != nil {
		log.Fatalf("Could not establish database connection\n error met: %s", err.Error())
	}

	// create the application with the defaults, an s3 filesystem and
	// a mailgun mailer see docs for more configuration possibilities
	app = &application.App{
		DB: db,
		FileSystem: gofile.NewS3FileSystem(
			service.EnvOrDefault("S3_REGION", "eu-west-1"),
			os.Getenv("S3_BUCKET"),
			&credentials.EnvProvider{},
		),
		Mailer: sys.NewMailer(),
		Time:   new(sys.OSTime),
		Logger: new(sys.StdLogger),
	}

	// add routes to the application using the default gorilla mux default routing
	// option, routins is secified in the routes.go file in the root of your application
	router := routing.NewGorillaRouter().AddRoutes(routing.ApplyRoutes(app, routes, auth.Middled))

	// create the server listening as default on 8082 but feel free to change in your .env
	return http.ListenAndServe(service.EnvOrDefault("APP_PORT", ":8082"), router)
}
