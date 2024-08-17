package context

import (
	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/danilbushkov/test-tasks/internal/app/db"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AppContext struct {
	Config *config.Config
	Fiber  *fiber.App
	DB     *db.DB
	Log    *logrus.Logger
}

func New(cf *config.Config, database *db.DB, log *logrus.Logger) *AppContext {
	return &AppContext{
		Config: cf,
		Fiber:  fiber.New(),
		DB:     database,
		Log:    log,
	}
}
