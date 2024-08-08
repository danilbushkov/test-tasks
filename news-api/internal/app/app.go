package app

import (
	"log"

	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	Config *config.Config
}

func New() *App {
	return &App{
		Config: config.New(),
	}
}

func (a *App) Run() {
	app := fiber.New()

	app.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
