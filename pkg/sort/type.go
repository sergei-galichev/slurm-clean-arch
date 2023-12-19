package sort

import (
	"slurm-clean-arch/pkg/columncode"
)

type Sort struct {
	Key columncode.ColumnCode
	Direction
}

type Direction string

const (
	DirectionAsc  Direction = "ASC"
	DirectionDesc Direction = "DESC"
)

func (d Direction) String() string {
	return string(d)
}

func (s Sort) Parsing(mapping map[columncode.ColumnCode]string) string {
	column, ok := mapping[s.Key]
	if !ok {
		return ""
	}
	return column + " " + s.Direction.String()
}

type Sorts []Sort

func (s Sorts) Parsing(mapping map[columncode.ColumnCode]string) []string {
	var result []string
	for _, sort := range s {
		result = append(result, sort.Parsing(mapping))
	}
	return result
}
