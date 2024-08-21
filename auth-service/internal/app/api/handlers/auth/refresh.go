package auth

import (
	"errors"

	e "github.com/danilbushkov/test-tasks/internal/app/errors"
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
