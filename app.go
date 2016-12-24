package butter

import (
	"butter/database"
	"butter/filesystem"
	"butter/mail"
	"butter/sys"
	"os"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// App is a struct which holds the butter application entities
type App struct {
	DB         database.DB
	ORM        database.ORM
	FileSystem filesystem.FileSystem
	Mailer     mail.Mailer
	Time       sys.Time
	Logger     sys.Logger
}

// GetAppHost return the app host
func (a App) GetAppHost() string {
	return os.Getenv("APP_URL") + os.Getenv("APP_PORT")
}

// NewMockApplication generates a new mock application for ease of testing
func NewMockApplication() (*App, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()

	app := &App{
		DB:         db,
		ORM:        new(database.MockORM),
		FileSystem: filesystem.NewMockFilesystem(),
		Mailer:     mail.NewMockMailer(),
		Time:       new(sys.MockTime),
		Logger:     new(sys.StdLogger),
	}

	return app, mock
}
