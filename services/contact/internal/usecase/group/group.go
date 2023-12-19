package group

import (
	"github.com/google/uuid"
	"slurm-clean-arch/pkg/queryparameter"
	"slurm-clean-arch/services/contact/internal/domain/group"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) Update(groupUpdate *group.Group) (*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) List(parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (*group.Group, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) Count() (uint64, error) {
	// TODO implement me
	panic("implement me")
}
