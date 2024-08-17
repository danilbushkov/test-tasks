package auth

import "github.com/danilbushkov/test-tasks/internal/app/context"

type AuthHandlers struct {
	appContext *context.AppContext
}

func New(actx *context.AppContext) *AuthHandlers {
	return &AuthHandlers{
		appContext: actx,
	}
}
