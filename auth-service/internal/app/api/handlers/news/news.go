package news

import "github.com/danilbushkov/test-tasks/internal/app/context"

type NewsHandlers struct {
	appContext *context.AppContext
}

func New(actx *context.AppContext) *NewsHandlers {
	return &NewsHandlers{
		appContext: actx,
	}
}
