package contact

import (
	"github.com/google/uuid"
	"slurm-clean-arch/pkg/type/queryparameter"
	"slurm-clean-arch/services/contact/internal/domain/contact"
	"time"
)

func (uc *UseCase) Create(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	return uc.adapterStorage.CreateContact(contacts...)
}

func (uc *UseCase) Update(contactUpdate *contact.Contact) (*contact.Contact, error) {
	return uc.adapterStorage.UpdateContact(
		contactUpdate.ID(),
		func(oldContact *contact.Contact) (*contact.Contact, error) {
			return contact.NewWithID(
				oldContact.ID(),
				oldContact.CreatedAt(),
				time.Now().UTC(),
				contactUpdate.PhoneNumber(),
				contactUpdate.Email(),
				contactUpdate.Name(),
				contactUpdate.Surname(),
				contactUpdate.Patronymic(),
				contactUpdate.Age(),
				contactUpdate.Gender(),
			)
		},
	)
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	return uc.adapterStorage.DeleteContact(ID)
}

func (uc *UseCase) List(parameter queryparameter.QueryParameter) ([]*contact.Contact, error) {
	return uc.adapterStorage.ListContact(parameter)
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *contact.Contact, err error) {
	return uc.adapterStorage.ReadContactByID(ID)
}

func (uc *UseCase) Count() (uint64, error) {
	return uc.adapterStorage.CountContact()
}
