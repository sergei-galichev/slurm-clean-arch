package group

import (
	"slurm-clean-arch/services/contact/internal/domain/group"
)

func ProtoToGroupResponse(response *group.Group) *ResponseGroup {
	return &ResponseGroup{
		ID:         response.ID().String(),
		CreatedAt:  response.CreatedAt(),
		ModifiedAt: response.ModifiedAt(),
		Group: Group{
			ShortGroup: ShortGroup{
				Name:        response.Name().String(),
				Description: response.Description().String(),
			},
			ContactsAmount: response.ContactCount(),
		},
	}
}
