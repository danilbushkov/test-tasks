package tokens

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type SignedRefreshToken struct {
	*RefreshToken

	token     string
	signature []byte
}

func ParseSignedRefreshToken(token string) (*SignedRefreshToken, error) {
	blocks := strings.Split(token, ".")
	if len(blocks) != 2 {
		return nil, errors.New("Invalid refresh token")
	}
	refreshToken := new(RefreshToken)

	structure, err := base64.URLEncoding.DecodeString(blocks[0])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(structure, refreshToken)

	if err != nil {
		return nil, err
	}
	sigature, err := base64.URLEncoding.DecodeString(blocks[1])
	if err != nil {
		return nil, err
	}

	return &SignedRefreshToken{
		token:        token,
		signature:    sigature,
		RefreshToken: refreshToken,
	}, nil
}

func (st *SignedRefreshToken) Check(key []byte) (bool, error) {
	structure, err := json.Marshal(st.RefreshToken)
	if err != nil {
		return false, err
	}
	data := append(structure, key...)

	hash := sha512.Sum512(data)

	if len(hash) != len(st.signature) {
		return false, nil
	}
	for i := range hash {
		if hash[i] != st.signature[i] {
			return false, nil
		}
	}
	return true, nil

}

func (st *SignedRefreshToken) Token() string {
	return st.token
}

func (st *SignedRefreshToken) Signature() []byte {
	return st.signature
}

func (st *SignedRefreshToken) SignatureHash() ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword(st.signature, bcrypt.DefaultCost)

	return hash, err
}
