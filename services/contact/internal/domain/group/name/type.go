package name

import (
	"github.com/pkg/errors"
)

var (
	ErrWrongLength = errors.Errorf(
		"name must be less than or equal to %d characters",
		MaxLength,
	)
)

const (
	MaxLength = 250
)

type Name struct {
	value string
}

func (n Name) String() string {
	return n.value
}

func New(name string) (
	Name,
	error,
) {
	if len([]rune(name)) > MaxLength {
		return Name{}, ErrWrongLength
	}
	return Name{
		value: name,
	}, nil
}
