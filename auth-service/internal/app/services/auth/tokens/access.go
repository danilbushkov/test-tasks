package tokens

import (
	"time"

	"github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessTokenClaims struct {
	UUID string `json:"uuid"`
	Ip   string `json:"ip"`
	jwt.RegisteredClaims
}

type AccessToken struct {
	token  *jwt.Token
	claims *AccessTokenClaims
}

func NewAccess(ip string, uuid *uuid.UUID, exp time.Time) *AccessToken {
	claims := &AccessTokenClaims{
		Ip:   ip,
		UUID: uuid.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return &AccessToken{
		token:  accessToken,
		claims: claims,
	}
}

func NewAccessFromString(s string, key string) (*AccessToken, error) {
	accessClaims := new(AccessTokenClaims)
	token, err := jwt.ParseWithClaims(s, accessClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, errors.TokenError(err)
	}
	return &AccessToken{
		token:  token,
		claims: accessClaims,
	}, nil
}

func (at *AccessToken) Signed(key []byte) (string, error) {
	refreshToken, err := at.token.SignedString(key)
	if err != nil {
		return "", errors.TokenError(err)
	}
	return refreshToken, nil
}

func (at *AccessToken) UUID() string {
	return at.claims.UUID
}
