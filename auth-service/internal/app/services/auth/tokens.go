package auth

import "github.com/golang-jwt/jwt/v5"

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenClaims struct {
	UUID string `json:"uuid"`
	Ip   string `json:"ip"`
	jwt.RegisteredClaims
}
