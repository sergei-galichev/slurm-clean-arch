package grpc

import (
	"context"
	contact "slurm-clean-arch/services/contact/internal/delivery/grpc/interface"
)

func (d *Delivery) CreateGroup(
	ctx context.Context,
	req *contact.CreateGroupRequest,
) (*contact.CreateGroupResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (d *Delivery) UpdateGroup(
	ctx context.Context,
	req *contact.UpdateGroupRequest,
) (*contact.UpdateGroupResponse, error) {
	// TODO implement me
	panic("implement me")
}

func (d *Delivery) DeleteGroup(
	ctx context.Context,
	req *contact.DeleteGroupRequest,
) (*contact.DeleteGroupResponse, error) {
	// TODO implement me
	panic("implement me")
}
