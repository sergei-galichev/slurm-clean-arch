package postgres

import (
	"github.com/google/uuid"
	"slurm-clean-arch/pkg/queryparameter"
	"slurm-clean-arch/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateContact(
	ID uuid.UUID,
	updateFn func(c *contact.Contact) (*contact.Contact, error),
) (*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteContact(ID uuid.UUID) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) ListContact(parameter queryparameter.QueryParameter) ([]*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) ReadContactByID(ID uuid.UUID) (response *contact.Contact, err error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) CountContact() (uint64, error) {
	// TODO implement me
	panic("implement me")
}
