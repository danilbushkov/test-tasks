package tests

import (
	"github.com/danilbushkov/test-tasks/internal/app"
	"github.com/danilbushkov/test-tasks/internal/app/db"
	"github.com/pashagolub/pgxmock/v4"
)

func getTestApp() (*app.App, error) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		return nil, err
	}
	return app.NewWithDB(&db.DB{Pool: mock}), nil
}
