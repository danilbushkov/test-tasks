package app

import (
	"log"
	"time"

	"github.com/danilbushkov/test-tasks/internal/app/api"
	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/danilbushkov/test-tasks/internal/app/context"
	"github.com/danilbushkov/test-tasks/internal/app/db"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type App struct {
	Context *context.AppContext
}

func NewWithDBAndTimeNow(db *db.DB, timeNow func() time.Time) *App {
	cf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	log := logrus.New()

	app := &App{
		Context: context.New(cf, db, log, timeNow),
	}

	api.New(app.Context).Reg()

	return app
}

func New() *App {
	cf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	log := logrus.New()

	d, err := db.New(cf.DB)
	if err != nil {
		log.Fatal(err)
	}
	app := &App{
		Context: context.New(cf, d, log, time.Now),
	}

	api.New(app.Context).Reg()

	return app
}

func (a *App) Api() *fiber.App {
	return a.Context.Fiber
}

func (a *App) Run() {
	defer a.Close()

	log.Fatal(a.Context.Fiber.Listen(":3000"))
}

func (a *App) Close() {
	a.Context.DB.Close()
}
