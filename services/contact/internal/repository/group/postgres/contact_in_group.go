package postgres

import (
	"github.com/google/uuid"
	"slurm-clean-arch/services/contact/internal/domain/contact"
)

func (r *Repository) CreateContactIntoGroup(
	groupID uuid.UUID,
	contacts ...*contact.Contact,
) ([]*contact.Contact, error) {
	// TODO implement me
	panic("implement me")
}
