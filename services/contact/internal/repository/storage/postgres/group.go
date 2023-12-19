package postgres

import (
	"github.com/google/uuid"
	"slurm-clean-arch/pkg/queryparameter"
	"slurm-clean-arch/services/contact/internal/domain/group"
)

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateGroup(
	ID uuid.UUID,
	updateFn func(group *group.Group) (*group.Group, error),
) (*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteGroup(ID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) ListGroup(parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) ReadGroupByID(ID uuid.UUID) (*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) CountGroup() (uint64, error) {
	// TODO implement me
	panic("implement me")
}
