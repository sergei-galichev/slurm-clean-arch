package group

import (
	"github.com/google/uuid"
	"slurm-clean-arch/services/contact/internal/domain/contact"
)

func (uc *UseCase) CreateContactIntoGroup(
	groupID uuid.UUID,
	contacts ...*contact.Contact,
) ([]*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) AddContactToGroup(groupID, contactID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (uc *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}
