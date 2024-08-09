package handlers

import (
	"github.com/danilbushkov/test-tasks/internal/app/api/handlers/news"
	"github.com/danilbushkov/test-tasks/internal/app/context"
)

type ApiHandlers struct {
	News *news.NewsHandlers

	appContext *context.AppContext
}

func New(appContext *context.AppContext) *ApiHandlers {
	return &ApiHandlers{
		appContext: appContext,
		News:       news.New(appContext),
	}
}
