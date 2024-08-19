package auth

import (
	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/danilbushkov/test-tasks/internal/app/context"
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
	claims := jwt.MapClaims{
		"ip":   ip,
		"uuid": uuid.String(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodNone, claims)

	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *AuthService) Refresh(refreshToken *jwt.Token) (*Tokens, error) {

	return nil, nil
}
