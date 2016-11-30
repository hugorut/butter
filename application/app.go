package application

import (
	"butter/database"
	"butter/sys"
	"os"

	"github.com/hugorut/gofile"
)

type App struct {
	DB         database.ORM
	FileSystem gofile.FileSystem
	Mailer     sys.Mailer
	Time       sys.Time
	Logger     sys.Logger
}

func (a App) GetAppHost() string {
	return os.Getenv("APP_URL") + os.Getenv("APP_PORT")
}

func NewMockApplication() (*App, *database.MockORM, *gofile.MockFileSystem, *sys.MockTime) {
	db := new(database.MockORM)
	files := gofile.NewMockFilesystem()
	time := new(sys.MockTime)
	app := &App{
		DB:         db,
		FileSystem: files,
		Mailer:     sys.NewMockMailer(),
		Time:       time,
		Logger:     new(sys.StdLogger),
	}

	return app, db, files, time
}
