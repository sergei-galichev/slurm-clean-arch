package patronymic

import (
	"github.com/pkg/errors"
)

var (
	ErrWrongLength = errors.Errorf(
		"patronymic must be less than or equal to %d characters",
		MaxLength,
	)
)

const (
	MaxLength = 100
)

type Patronymic string

func (p Patronymic) String() string {
	return string(p)
}

func New(patronymic string) (
	*Patronymic,
	error,
) {
	if len([]rune(patronymic)) > MaxLength {
		return nil, ErrWrongLength
	}
	p := Patronymic(patronymic)
	return &p, nil
}
