package structs

import (
	"errors"
)

type RefreshToken struct {
	Token *string `json:"refreshToken"`
}

func (rt *RefreshToken) Validate() error {
	if rt.Token == nil {
		return errors.New("refreshToken field is not set")
	}
	return nil
}
