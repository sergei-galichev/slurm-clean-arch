package description

import (
	"github.com/pkg/errors"
)

var (
	ErrWrongLength = errors.Errorf(
		"description must be less than or equal to %d characters",
		MaxLength,
	)
)

const (
	MaxLength = 1000
)

type Description struct {
	value string
}

func (d Description) String() string {
	return d.value
}

func New(description string) (
	Description,
	error,
) {
	if len([]rune(description)) > MaxLength {
		return Description{}, ErrWrongLength
	}
	return Description{
		value: description,
	}, nil
}
