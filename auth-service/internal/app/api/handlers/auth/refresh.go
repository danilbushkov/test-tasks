package auth

import (
	"github.com/danilbushkov/test-tasks/internal/app/api/errors"
	auth_service "github.com/danilbushkov/test-tasks/internal/app/services/auth"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/gofiber/fiber/v2"
)

func (ah *AuthHandlers) Refresh(c *fiber.Ctx) error {
	refresh := new(structs.RefreshToken)

	if err := c.BodyParser(refresh); err != nil {
		ah.appContext.Log.Info(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := refresh.Validate()
	if err != nil {
		ah.appContext.Log.Info(err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			structs.Error{
				Message: err.Error(),
			},
		)
	}

	authService := auth_service.NewFromAppContext(ah.appContext)

	tokens, err := authService.Refresh(*refresh.Token, c.IP())
	if err != nil {
		if err = errors.HendleError(ah.appContext, c, err); err != nil {
			return err
		}
	}
	return c.JSON(tokens)

}
