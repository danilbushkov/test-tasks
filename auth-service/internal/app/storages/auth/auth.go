package auth

import (
	"context"
	"fmt"

	ctx "github.com/danilbushkov/test-tasks/internal/app/context"
	"github.com/danilbushkov/test-tasks/internal/app/db"
	e "github.com/danilbushkov/test-tasks/internal/app/errors"
)

type AuthStorage struct {
	db *db.DB
}

func New(db *db.DB) *AuthStorage {
	return &AuthStorage{
		db,
	}
}

func NewFromAppContext(actx *ctx.AppContext) *AuthStorage {
	return &AuthStorage{
		db: actx.DB,
	}
}

func (as *AuthStorage) AddRefreshToken(hash []byte) error {
	_, err := as.db.Pool.Exec(context.Background(),
		"INSERT INTO auth(refresh_token) VALUES ($1)",
		hash,
	)
	if err != nil {
		return fmt.Errorf("%w: %w", e.ErrDatabase, err)
	}

	return nil
}

// Return true if the token has been deleted
func (as *AuthStorage) DeleteRefreshToken(hash []byte) (bool, error) {
	commandTag, err := as.db.Pool.Exec(context.Background(),
		"DELETE FROM auth WHERE refresh_token=$1",
		hash,
	)
	if err != nil {
		return false, fmt.Errorf("%w: %w", e.ErrDatabase, err)
	}
	if commandTag.RowsAffected() != 1 {
		return false, nil
	}

	return true, nil
}
