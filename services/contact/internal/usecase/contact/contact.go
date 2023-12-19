package contact

import (
	"github.com/google/uuid"
	"slurm-clean-arch/pkg/queryparameter"
	"slurm-clean-arch/services/contact/internal/domain/contact"
)

func (uc *UseCase) Create(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) Update(contactUpdate *contact.Contact) (*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) List(parameter queryparameter.QueryParameter) ([]*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *contact.Contact, err error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) Count() (uint64, error) {
	// TODO implement me
	panic("implement me")
}
