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
}

func NewFromAppContext(actx *context.AppContext) *AuthService {
	return &AuthService{
		authStorage: auth_storage.NewFromAppContext(actx),
		tokenConfig: actx.Config.Token,
	}
}

func (as *AuthService) Get(uuid *uuid.UUID, ip string) (*structs.Tokens, error) {
	accessToken, err := tokens.NewAccess(ip, uuid, time.Duration(as.tokenConfig.AccessLifeTime)*time.Second).Signed([]byte(as.tokenConfig.AccessKey))

	if err != nil {
		return nil, err
	}
	//hash, err := bcrypt.GenerateFromPassword(rt.Signature, bcrypt.DefaultCost)

	err = as.authStorage.AddRefreshToken([]byte{1, 1, 1})
	if err != nil {
		return nil, err
	}

	return &structs.Tokens{
		AccessToken:  accessToken,
		RefreshToken: "",
	}, nil
}

func (as *AuthService) Refresh(refreshToken *jwt.Token) (*structs.Tokens, error) {

	return nil, nil
}
