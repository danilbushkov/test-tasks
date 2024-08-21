package tokens

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/danilbushkov/test-tasks/internal/app/errors"
	"github.com/google/uuid"
)

type RefreshToken struct {
	Ip   string    `json:"ip"`
	UUID string    `json:"uuid"`
	Exp  time.Time `json:"exp"`
}

func NewRefresh(ip string, uuid *uuid.UUID, exp time.Time) *RefreshToken {
	return &RefreshToken{
		Ip:   ip,
		UUID: uuid.String(),
		Exp:  exp,
	}
}

func (rt *RefreshToken) Sign(key []byte) (*SignedRefreshToken, error) {
	structure, err := json.Marshal(rt)
	if err != nil {
		return nil, errors.TokenError(err)
	}
	data := append(structure, key...)

	hash := sha512.Sum512(data)

	eData := base64.URLEncoding.EncodeToString(structure)

	eHash := base64.URLEncoding.EncodeToString(hash[:])

	token := eData + "." + eHash

	return &SignedRefreshToken{
		token:        token,
		signature:    hash[:],
		RefreshToken: rt,
	}, nil
}
