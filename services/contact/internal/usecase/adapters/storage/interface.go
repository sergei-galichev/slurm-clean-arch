package storage

import (
	"github.com/google/uuid"
	"slurm-clean-arch/pkg/queryparameter"
	"slurm-clean-arch/services/contact/internal/domain/contact"
	"slurm-clean-arch/services/contact/internal/domain/group"
)

type Storage interface {
	Contact
	Group
}

type Contact interface {
	CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error)
	UpdateContact(ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error)
	DeleteContact(ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	ListContact(parameter queryparameter.QueryParameter) ([]*contact.Contact, error)
	ReadContactByID(ID uuid.UUID) (response *contact.Contact, err error)
	CountContact() (uint64, error) // maybe add filters to parameters
}

type Group interface {
	CreateGroup(group *group.Group) (*group.Group, error)
	UpdateGroup(ID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error)
	DeleteGroup(ID uuid.UUID) error // maybe add filters to parameters

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	ListGroup(parameter queryparameter.QueryParameter) ([]*group.Group, error)
	ReadGroupByID(ID uuid.UUID) (*group.Group, error)
	CountGroup() (uint64, error) // maybe add filters to parameters
}

type ContactInGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, in ...*contact.Contact) ([]*contact.Contact, error)
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
	AddContactToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error
}
