package postgres

import (
	"github.com/google/uuid"
	"slurm-clean-arch/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContactIntoGroup(
	groupID uuid.UUID,
	in ...*contact.Contact,
) ([]*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) AddContactToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}
