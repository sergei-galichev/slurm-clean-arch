package grpc

import (
	contact "slurm-clean-arch/services/contact/internal/delivery/grpc/interface"
	"slurm-clean-arch/services/contact/internal/usecase"
)

type Delivery struct {
	contact.UnimplementedContactServiceServer
	ucContact usecase.Contact
	ucGroup   usecase.Group

	options Options
}

type Options struct{}

func New(ucContact usecase.Contact, ucGroup usecase.Group, options Options) *Delivery {
	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}

	d.SetOptions(options)

	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}
