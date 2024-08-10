package news

import (
	"errors"

	app_err "github.com/danilbushkov/test-tasks/internal/app/errors"
	ns "github.com/danilbushkov/test-tasks/internal/app/storages/news"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/gofiber/fiber/v2"
)

func (nh *NewsHandlers) Add(c *fiber.Ctx) error {

	n := new(structs.News)

	if err := c.BodyParser(n); err != nil {
		nh.appContext.Log.Info(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	newsStorage := ns.NewFromAppContext(nh.appContext)
	if err := newsStorage.Add(n); err != nil {
		if errors.Is(err, app_err.ErrDatabase) {
			nh.appContext.Log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		} else {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(structs.Error{
				Message: err.Error(),
			})
		}
	}
	return c.SendStatus(fiber.StatusCreated)
}
