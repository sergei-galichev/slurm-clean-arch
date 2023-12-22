package queryparameter

import (
	"slurm-clean-arch/pkg/type/pagination"
	"slurm-clean-arch/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	// TODO: add filters
}
