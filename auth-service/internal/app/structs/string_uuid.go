package structs

import (
	"errors"
	"github.com/google/uuid"
)

type StringUUID struct {
	UUID *string `json:"uuid"`
}

func (su *StringUUID) Validate() error {
	if su.UUID == nil {
		return errors.New("uuid field is not set")
	}

	return nil
}

func (su *StringUUID) ToUUID() (*uuid.UUID, error) {
	if err := su.Validate(); err != nil {
		return nil, err
	}
	u, err := uuid.Parse(*su.UUID)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
