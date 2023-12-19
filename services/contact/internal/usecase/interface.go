package usecase

import (
	"github.com/google/uuid"
	"slurm-clean-arch/pkg/queryparameter"
	"slurm-clean-arch/services/contact/internal/domain/contact"
	"slurm-clean-arch/services/contact/internal/domain/group"
)

type Contact interface {
	Create(contacts ...*contact.Contact) ([]*contact.Contact, error)
	Update(contact *contact.Contact) (*contact.Contact, error)
	Delete(ID uuid.UUID) error // maybe add filters to parameters
	ContactReader
}

type ContactReader interface {
	List(parameter queryparameter.QueryParameter) ([]*contact.Contact, error)
	ReadByID(ID uuid.UUID) (response *contact.Contact, err error)
	Count() (uint64, error) // maybe add filters to parameters
}

type Group interface {
	Create(domainGroup *group.Group) (*group.Group, error)
	Update(group *group.Group) (*group.Group, error)
	Delete(ID uuid.UUID) error // maybe add filters to parameters

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	List(parameter queryparameter.QueryParameter) ([]*group.Group, error)
	ReadByID(ID uuid.UUID) (*group.Group, error)
	Count() (uint64, error) // maybe add filters to parameters
}

type ContactInGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error)
	AddContactToGroup(groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
