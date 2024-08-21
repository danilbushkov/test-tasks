package auth

import (
	"time"

	"github.com/danilbushkov/test-tasks/internal/app/config"
	"github.com/danilbushkov/test-tasks/internal/app/context"
	"github.com/danilbushkov/test-tasks/internal/app/services/auth/tokens"
	auth_storage "github.com/danilbushkov/test-tasks/internal/app/storages/auth"
	"github.com/danilbushkov/test-tasks/internal/app/structs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	authStorage *auth_storage.AuthStorage
	tokenConfig *config.TokenConfig
	timeNow     func() time.Time
}

func NewFromAppContext(actx *context.AppContext) *AuthService {
	return &AuthService{
		authStorage: auth_storage.NewFromAppContext(actx),
		tokenConfig: actx.Config.Token,
		timeNow:     actx.TimeNow,
	}
}

func (as *AuthService) Get(uuid *uuid.UUID, ip string) (*structs.Tokens, error) {
	accessExp := as.timeNow().Add(time.Duration(as.tokenConfig.AccessLifeTime) * time.Second)

	accessToken, err := tokens.NewAccess(ip, uuid, accessExp).Signed([]byte(as.tokenConfig.AccessKey))

	if err != nil {
		return nil, err
	}

	refreshExp := as.timeNow().Add(time.Duration(as.tokenConfig.RefreshLifeTime) * time.Second)

	signedRefresh, err := tokens.NewRefresh(ip, uuid, refreshExp).Sign([]byte(as.tokenConfig.RefreshKey))

	hash, err := signedRefresh.SignatureHash()
	if err != nil {
		return nil, err
	}
	err = as.authStorage.AddRefreshToken(hash)
	if err != nil {
		return nil, err
	}

	return &structs.Tokens{
		AccessToken:  accessToken,
		RefreshToken: signedRefresh.Token(),
	}, nil
}

func (as *AuthService) Refresh(refreshToken *jwt.Token) (*structs.Tokens, error) {

	return nil, nil
}
