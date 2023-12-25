package group

import (
	"time"
)

type ResponseGroup struct {
	// Group ID
	ID string `json:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	// Group create date
	CreatedAt time.Time `json:"createdAt" binding:"required"`
	// Group last update date
	ModifiedAt time.Time `json:"modifiedAt" binding:"required"`

	Group
}

// Group contains information about group
type Group struct {
	ShortGroup
	// Contacts amount in group
	ContactsAmount uint64 `json:"contactsAmount" default:"10" binding:"min=0" minimum:"0"`
}

type ShortGroup struct {
	// Group name
	Name string `json:"name" binding:"required,max=100" maxLength:"100" example:"Название группы"`
	// Group description
	Description string `json:"description" binding:"max=1000" maxLength:"1000" example:"Описание группы"`
}

type ListGroup struct {
	// Total count
	Total uint64 `json:"total" example:"10" default:"0" binding:"min=0" minimum:"0"`
	// Records limit
	Limit uint64 `json:"limit" example:"10" default:"10" binding:"min=0" minimum:"0"`
	// Offset get records
	Offset uint64 `json:"offset" example:"20" default:"0" binding:"min=0" minimum:"0"`
	// List of groups
	List []*ResponseGroup `json:"list" binding:"min=0" minimum:"0"`
}

type ID struct {
	// Group ID
	Value string `json:"id" uri:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
}

type ContactID struct {
	// Contact ID
	Value string `json:"id" uri:"contactId" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
}
