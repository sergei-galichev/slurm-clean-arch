package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slurm-clean-arch/pkg/tools/converter"
	"slurm-clean-arch/pkg/type/phonenumber"
	jsonContact "slurm-clean-arch/services/contact/internal/delivery/http/contact"
	jsonGroup "slurm-clean-arch/services/contact/internal/delivery/http/group"
	domainContact "slurm-clean-arch/services/contact/internal/domain/contact"
	"slurm-clean-arch/services/contact/internal/domain/contact/age"
	"slurm-clean-arch/services/contact/internal/domain/contact/name"
	"slurm-clean-arch/services/contact/internal/domain/contact/patronymic"
	"slurm-clean-arch/services/contact/internal/domain/contact/surname"
)

// CreateContactIntoGroup
// @Summary Method allow to create contact into group
// @Description Method allow to create contact into group
// @Security Cookies
// @Tags groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param contact body jsonContact.ShortContact true "Contact data"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /groups/{id}/contacts/ [post]
func (d *Delivery) CreateContactIntoGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contact := jsonContact.ShortContact{}
	if err := c.ShouldBindJSON(&contact); err != nil {
		SetError(c, http.StatusBadRequest, fmt.Errorf("payload is not correct, Error: %w", err))
		return
	}

	contactName, err := name.New(contact.Name)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contactSurname, err := surname.New(contact.Surname)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contactPatronymic, err := patronymic.New(contact.Patronymic)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contactAge, err := age.New(contact.Age)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	dContact, err := domainContact.New(
		*phonenumber.New(contact.PhoneNumber),
		contact.Email,
		*contactName,
		*contactSurname,
		*contactPatronymic,
		*contactAge,
		contact.Gender,
	)

	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contacts, err := d.ucGroup.CreateContactIntoGroup(converter.StringToUUID(id.Value), dContact)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	var list []*jsonContact.ContactResponse
	for _, item := range contacts {
		list = append(list, jsonContact.ToContactResponse(item))
	}

	c.JSON(http.StatusOK, list)
}

// AddContactToGroup
// @Summary Method allow to add contact to group
// @Description Method allow to add contact to group
// @Tags groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param contactId path string true "Contact ID"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /groups/{id}/contacts/{contactId} [post]
func (d *Delivery) AddContactToGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	var contactID jsonContact.ID
	if err := c.ShouldBindUri(&contactID); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groupUUID := converter.StringToUUID(id.Value)
	contactUUID := converter.StringToUUID(contactID.Value)
	if err := d.ucGroup.AddContactToGroup(groupUUID, contactUUID); err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// DeleteContactFromGroup
// @Summary Method allow to delete contact from group
// @Description Method allow to delete contact from group
// @Tags groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param contactId path string true "Contact ID"
// @Success 200
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /groups/{id}/contacts/{contactId} [delete]
func (d *Delivery) DeleteContactFromGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	var contactID jsonContact.ID
	if err := c.ShouldBindUri(&contactID); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}
	groupUUID := converter.StringToUUID(id.Value)
	contactUUID := converter.StringToUUID(contactID.Value)
	if err := d.ucGroup.DeleteContactFromGroup(groupUUID, contactUUID); err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}
