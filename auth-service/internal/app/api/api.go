package api

import (
	"github.com/danilbushkov/test-tasks/internal/app/api/handlers"
	"github.com/danilbushkov/test-tasks/internal/app/context"
	"github.com/gofiber/fiber/v2"
)

type Api struct {
	Handlers   *handlers.ApiHandlers
	appContext *context.AppContext
}

func New(actx *context.AppContext) *Api {
	return &Api{
		Handlers:   handlers.New(actx),
		appContext: actx,
	}
}

func (api *Api) registrar() *fiber.App {
	return api.appContext.Fiber
}

func (api *Api) Reg() {
	api.regNewsApi()
}
