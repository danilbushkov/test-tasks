package auth

import (
	"time"

	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/danilbushkov/test-tasks/internal/app/context"
	"github.com/danilbushkov/test-tasks/internal/app/errors"
	auth_storage "github.com/danilbushkov/test-tasks/internal/app/storages/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	authStorage *auth_storage.AuthStorage
	tokenConfig *config.TokenConfig
}

func NewFromAppContext(actx *context.AppContext) *AuthService {
	return &AuthService{
		authStorage: auth_storage.NewFromAppContext(actx),
		tokenConfig: actx.Config.Token,
	}
}

func (as *AuthService) Get(uuid *uuid.UUID, ip string) (*Tokens, error) {
	claims := &TokenClaims{
		Ip:   ip,
		UUID: uuid.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(as.tokenConfig.AccessLifeTime) * time.Second)),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(as.tokenConfig.AccessKey))
	if err != nil {
		return nil, errors.TokenError(err)
	}
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(as.tokenConfig.RefreshLifeTime) * time.Second)),
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(as.tokenConfig.RefreshKey))
	if err != nil {
		return nil, errors.TokenError(err)
	}

	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *AuthService) Refresh(refreshToken *jwt.Token) (*Tokens, error) {

	return nil, nil
}
