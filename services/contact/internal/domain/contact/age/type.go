package age

import (
	"github.com/pkg/errors"
	"strconv"
)

var (
	ErrWrongLength = errors.Errorf(
		"age must be less than or equal to %d",
		MaxLength,
	)
)

const (
	MaxLength uint8 = 200
)

type Age uint8

func (a Age) String() string {
	return strconv.FormatUint(
		uint64(a),
		10,
	)
}

func New(age uint8) (
	*Age,
	error,
) {
	if age > MaxLength {
		return nil, ErrWrongLength
	}
	a := Age(age)
	return &a, nil
}

func (a Age) Equal(age Age) bool {
	return a == age
}

func (a Age) Less(age Age) bool {
	return a < age
}

func (a Age) More(age Age) bool {
	return a > age
}
