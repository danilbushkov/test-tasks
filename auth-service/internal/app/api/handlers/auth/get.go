package auth

import (
	"github.com/danilbushkov/test-tasks/internal/app/api/errors"
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
		if err = errors.HendleError(ah.appContext, c, err); err != nil {
			return err
		}
	}
	return c.JSON(tokens)
}
