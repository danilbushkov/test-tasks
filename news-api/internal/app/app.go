package app

import (
	"log"

	"github.com/danilbushkov/test-tasks/internal/app/api"
	"github.com/danilbushkov/test-tasks/internal/app/context"
)

type App struct {
	Context *context.AppContext
}

func New() *App {
	app := &App{
		Context: context.New(),
	}

	api.New(app.Context).Reg()

	return app
}

func (a *App) Run() {

	log.Fatal(a.Context.Fiber.Listen(":3000"))
}
