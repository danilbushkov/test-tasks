package context

import (
	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/gofiber/fiber/v2"
)

type AppContext struct {
	Config *config.Config
	Fiber  *fiber.App
}

func New() *AppContext {
	return &AppContext{
		Config: config.New(),
		Fiber:  fiber.New(),
	}
}
