package auth

import "github.com/golang-jwt/jwt/v5"

type Tokens struct {
	AccessToken  *jwt.Token `json:"accessToken"`
	RefreshToken *jwt.Token `json:"refreshToken"`
}
