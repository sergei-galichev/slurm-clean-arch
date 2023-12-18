package surname

import (
	"github.com/pkg/errors"
)

var (
	ErrWrongLength = errors.Errorf(
		"surname must be less than or equal to %d characters",
		MaxLength,
	)
)

const (
	MaxLength = 100
)

type Surname string

func (s Surname) String() string {
	return string(s)
}

func New(surname string) (
	*Surname,
	error,
) {
	if len([]rune(surname)) > MaxLength {
		return nil, ErrWrongLength
	}
	s := Surname(surname)
	return &s, nil
}
