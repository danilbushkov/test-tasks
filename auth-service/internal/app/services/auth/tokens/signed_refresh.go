package tokens

import "golang.org/x/crypto/bcrypt"

type SignedRefreshToken struct {
	*RefreshToken

	token     string
	signature []byte
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
