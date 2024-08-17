package news

import (
	"errors"
	"strconv"

	app_err "github.com/danilbushkov/test-tasks/internal/app/errors"
	ns "github.com/danilbushkov/test-tasks/internal/app/storages/news"
	"github.com/danilbushkov/test-tasks/internal/app/structs"

	"github.com/gofiber/fiber/v2"
)

func (nh *NewsHandlers) Delete(c *fiber.Ctx) error {
	p := c.Params("id")
	id, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newsStorage := ns.NewFromAppContext(nh.appContext)
	if err := newsStorage.Delete(id); err != nil {
		if errors.Is(err, app_err.ErrDatabase) {
			nh.appContext.Log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		} else {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(structs.Error{
				Message: err.Error(),
			})
		}
	}
	return c.SendStatus(fiber.StatusNoContent)
}
