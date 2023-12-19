package queryparameter

import (
	"slurm-clean-arch/pkg/pagination"
	"slurm-clean-arch/pkg/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	// TODO: add filters
}
