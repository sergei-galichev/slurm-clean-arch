package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"slurm-clean-arch/pkg/columncode"
	"slurm-clean-arch/pkg/queryparameter"
	"slurm-clean-arch/pkg/tools/transaction"
	"slurm-clean-arch/services/contact/internal/domain/contact"
)

var mappingSortContact = map[columncode.ColumnCode]string{
	"id":          "id",
	"fullName":    "full_name",
	"phoneNumber": "phone_number",
	"name":        "name",
	"surname":     "surname",
	"patronymic":  "patronymic",
	"email":       "email",
	"gender":      "gender",
	"age":         "age",
}

func (r *Repository) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, tx, err)
	}(ctx, tx)

	return nil, nil

}

func (r *Repository) createContactTx(
	ctx context.Context,
	tx pgx.Tx,
	contacts ...*contact.Contact,
) ([]*contact.Contact, error) {
	if len(contacts) == 0 {
		return []*contact.Contact{}, nil
	}

	//_, err := tx.CopyFrom(
	//	ctx,
	//	pgx.Identifier{"slurm", "contact"},
	//	dao.CreateColumnContact,
	//	r.toCopyFromSource(contacts...),
	//)
	//if err != nil {
	//	return nil, err
	//}

	return contacts, nil
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
