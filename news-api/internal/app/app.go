package app

import (
	"log"

	"github.com/danilbushkov/test-tasks/internal/app/api"
	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/danilbushkov/test-tasks/internal/app/context"
	"github.com/danilbushkov/test-tasks/internal/app/db"
	"github.com/sirupsen/logrus"
)

type App struct {
	Context *context.AppContext
}

func New() *App {
	cf := config.New()
	log := logrus.New()
	d, err := db.New(cf.DB, log)
	if err != nil {
		log.Fatal(err)
	}
	app := &App{
		Context: context.New(config.New(), d, log),
	}

	api.New(app.Context).Reg()

	return app
}

func (a *App) Run() {
	defer a.Context.DB.Close()

	log.Fatal(a.Context.Fiber.Listen(":3000"))
}
