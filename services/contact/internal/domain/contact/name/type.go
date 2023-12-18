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
	MaxLength = 50
)

type Name string

func (n Name) String() string {
	return string(n)
}

func New(name string) (
	*Name,
	error,
) {
	if len([]rune(name)) > MaxLength {
		return nil, ErrWrongLength
	}
	n := Name(name)
	return &n, nil
}
