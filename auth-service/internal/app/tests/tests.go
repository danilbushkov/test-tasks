package tests

import (
	"testing"

	"github.com/danilbushkov/test-tasks/internal/app"
	"github.com/danilbushkov/test-tasks/internal/app/db"
	"github.com/pashagolub/pgxmock/v4"
)

func getTestApp() (*app.App, pgxmock.PgxPoolIface, error) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		return nil, nil, err
	}
	return app.NewWithDB(&db.DB{Pool: mock}), mock, nil
}

func setEnv(t *testing.T) {

	t.Setenv("ACCESS_LIFE_TIME", "300")
	t.Setenv("ACCESS_KEY", "a_key")
	t.Setenv("REFRESH_LIFE_TIME", "3600")
	t.Setenv("REFRESH_KEY", "r_key")
}
