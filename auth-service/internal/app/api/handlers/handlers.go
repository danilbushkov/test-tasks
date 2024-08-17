package handlers

import (
	"github.com/danilbushkov/test-tasks/internal/app/api/handlers/auth"
	"github.com/danilbushkov/test-tasks/internal/app/context"
)

type ApiHandlers struct {
	Auth *auth.AuthHandlers

	appContext *context.AppContext
}

func New(appContext *context.AppContext) *ApiHandlers {
	return &ApiHandlers{
		appContext: appContext,
		Auth:       auth.New(appContext),
	}
}
