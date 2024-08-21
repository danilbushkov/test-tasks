package auth

import (
	"errors"
	e "github.com/danilbushkov/test-tasks/internal/app/errors"
	auth_service "github.com/danilbushkov/test-tasks/internal/app/services/auth"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/gofiber/fiber/v2"
)

func (ah *AuthHandlers) Get(c *fiber.Ctx) error {
	stringUUID := new(structs.StringUUID)
	if err := c.BodyParser(stringUUID); err != nil {
		ah.appContext.Log.Info(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	uuid, err := stringUUID.ToUUID()
	if err != nil {
		ah.appContext.Log.Info(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			structs.Error{
				Message: err.Error(),
			},
		)
	}

	authService := auth_service.NewFromAppContext(ah.appContext)

	tokens, err := authService.Get(uuid, c.IP())

	if err != nil {
		if errors.Is(err, e.ErrDatabase) || errors.Is(err, e.ErrToken) {
			ah.appContext.Log.Error(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		} else {
			ah.appContext.Log.Info(err)
			return c.Status(fiber.StatusUnprocessableEntity).JSON(structs.Error{
				Message: err.Error(),
			})
		}
	}

	return c.JSON(tokens)
}
