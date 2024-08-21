package errors

import (
	"errors"
	"github.com/danilbushkov/test-tasks/internal/app/context"
	e "github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/gofiber/fiber/v2"
)

func HendleError(ac *context.AppContext, c *fiber.Ctx, err error) error {
	if err != nil {
		if errors.Is(err, e.ErrDatabase) || errors.Is(err, e.ErrToken) {
			ac.Log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		} else {
			ac.Log.Error(err)
			return c.Status(fiber.StatusUnprocessableEntity).JSON(structs.Error{
				Message: err.Error(),
			})
		}
	}
	return nil
}
