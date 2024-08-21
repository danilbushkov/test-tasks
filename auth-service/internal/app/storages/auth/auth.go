package auth

import (
	"context"
	"errors"

	ctx "github.com/danilbushkov/test-tasks/internal/app/context"
	"github.com/danilbushkov/test-tasks/internal/app/db"
	e "github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/google/uuid"
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

func (as *AuthStorage) AddRefreshTokenSignature(uuid *uuid.UUID, hash []byte) error {
	_, err := as.db.Pool.Exec(context.Background(),
		"INSERT INTO auth(uuid, refresh_token_signature) VALUES ($1, $2)",
		uuid.String(),
		hash,
	)
	if err != nil {
		return e.DatabaseError(err)
	}

	return nil
}

func (as *AuthStorage) DeleteRefreshTokenSignature(hash []byte) error {
	commandTag, err := as.db.Pool.Exec(context.Background(),
		"DELETE FROM auth WHERE refresh_token_signature=$1",
		hash,
	)
	if err != nil {
		return e.DatabaseError(err)
	}
	if commandTag.RowsAffected() != 1 {
		return errors.New("Signature not found")
	}

	return nil
}

func (as *AuthStorage) GetRefreshTokenSignature(uuid string) ([]byte, error) {
	signature := []byte{}
	err := as.db.Pool.QueryRow(context.Background(),
		"SELECT refresh_token_signature FROM auth WHERE uuid=$1",
		uuid,
	).Scan(signature)
	if err != nil {
		return nil, e.DatabaseError(err)
	}

	return signature, nil
}
