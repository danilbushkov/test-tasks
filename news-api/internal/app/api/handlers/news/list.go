package news

import (
	"errors"

	app_err "github.com/danilbushkov/test-tasks/internal/app/errors"
	ns "github.com/danilbushkov/test-tasks/internal/app/storages/news"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/gofiber/fiber/v2"
)

func (nh *NewsHandlers) List(c *fiber.Ctx) error {
	newsStorage := ns.NewFromAppContext(nh.appContext)
	list, err := newsStorage.List()
	if err != nil {
		if errors.Is(err, app_err.ErrDatabase) {
			nh.appContext.Log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		} else {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(structs.Error{
				Message: err.Error(),
			})
		}
	}
	return c.JSON(map[string]any{
		"Success": true,
		"News":    list,
	})
}
