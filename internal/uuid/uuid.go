package uuid

import (
	"github.com/hantdev/certs/errors"
	"github.com/gofrs/uuid"
)

// ErrGeneratingID indicates error in generating UUID.
var ErrGeneratingID = errors.New("failed to generate uuid")

type uuidProvider struct{}

// New instantiates a UUID provider.
func New() IDProvider {
	return &uuidProvider{}
}

// IDProvider specifies an API for generating unique identifiers.
type IDProvider interface {
	ID() (string, error)
}

func (up *uuidProvider) ID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.Wrap(ErrGeneratingID, err)
	}

	return id.String(), nil
}