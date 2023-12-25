package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"slurm-clean-arch/pkg/tools/converter"
	"slurm-clean-arch/pkg/type/pagination"
	"slurm-clean-arch/pkg/type/phonenumber"
	"slurm-clean-arch/pkg/type/query"
	"slurm-clean-arch/pkg/type/queryparameter"
	jsonContact "slurm-clean-arch/services/contact/internal/delivery/http/contact"
	domainContact "slurm-clean-arch/services/contact/internal/domain/contact"
	"slurm-clean-arch/services/contact/internal/domain/contact/age"
	"slurm-clean-arch/services/contact/internal/domain/contact/name"
	"slurm-clean-arch/services/contact/internal/domain/contact/patronymic"
	"slurm-clean-arch/services/contact/internal/domain/contact/surname"
	"time"
)

var mappingSortsContact = query.SortsOptions{
	"name":        {},
	"surname":     {},
	"patronymic":  {},
	"phoneNumber": {},
	"email":       {},
	"gender":      {},
	"age":         {},
}

// CreateContact
// @Summary     Method allow to create contact
// @Description Method allow to create contact
// @Tags        contacts
// @Accept      json
// @Produce     json
// @Param       contact body jsonContact.ShortContact true "Contact data"
// @Success     201 {object} jsonContact.ContactResponse "Contact structure"
// @Success     200
// @Failure     400 {object} ErrorResponse
// @Failure     403 "403 Forbidden"
// @Failure     404 {object} ErrorResponse "404 Not Found"
// @Router      /contacts/ [post]
func (d *Delivery) CreateContact(c *gin.Context) {
	contact := jsonContact.ShortContact{}
	if err := c.ShouldBindJSON(&contact); err != nil {
		SetError(c, http.StatusBadRequest, fmt.Errorf("payload is not correct, Error: %w", err))
		return
	}

	contactAge, err := age.New(contact.Age)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
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

	response, err := d.ucContact.Create(dContact)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	if len(response) > 0 {
		c.JSON(http.StatusCreated, jsonContact.ToContactResponse(response[0]))
	} else {
		c.Status(http.StatusOK)
	}
}

// UpdateContact
// @Summary Method allow to update contact data
// @Description Method allow to update contact data
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path string true "Contact ID"
// @Param contact body jsonContact.ShortContact true "Contact data"
// @Success 200 {object} jsonContact.ContactResponse "Contact structure"
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /contacts/{id} [put]
func (d *Delivery) UpdateContact(c *gin.Context) {
	var id jsonContact.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contact := jsonContact.ShortContact{}
	if err := c.ShouldBindJSON(&contact); err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	contactAge, err := age.New(contact.Age)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
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

	var dContact, _ = domainContact.NewWithID(
		converter.StringToUUID(id.Value),
		time.Now().UTC(),
		time.Now().UTC(),
		*phonenumber.New(contact.PhoneNumber),
		contact.Email,
		*contactName,
		*contactSurname,
		*contactPatronymic,
		*contactAge,
		contact.Gender,
	)

	response, err := d.ucContact.Update(dContact)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, jsonContact.ToContactResponse(response))
}

// DeleteContact
// @Summary Method allow to delete contact
// @Description Method allow to delete contact
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path string true "Contact ID"
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /contacts/{id} [delete]
func (d *Delivery) DeleteContact(c *gin.Context) {
	var id jsonContact.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	if err := d.ucContact.Delete(converter.StringToUUID(id.Value)); err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// ListContact
// @Summary Method allow to get list of contacts
// @Description Method allow to get list of contacts
// @Tags contacts
// @Accept json
// @Produce json
// @Param limit query int false "Record count limit" default(10) minimum(0) maximum(100)
// @Param offset query int false "Get records offset" default(0) minimum(0)
// @Param sort query string false "Sort records by field" default(name)
// @Success 200 {object} jsonContact.ListContact "List of contacts"
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /contacts/ [get]
func (d *Delivery) ListContact(c *gin.Context) {
	params, err := query.ParseQuery(
		c, query.Options{
			Sorts: mappingSortsContact,
		},
	)

	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	contacts, err := d.ucContact.List(
		queryparameter.QueryParameter{
			Sorts: params.Sorts,
			Pagination: pagination.Pagination{
				Limit:  params.Limit,
				Offset: params.Offset,
			},
		},
	)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	count, err := d.ucContact.Count()
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	var result = jsonContact.ListContact{
		Total:  count,
		Limit:  params.Limit,
		Offset: params.Offset,
		List:   []*jsonContact.ContactResponse{},
	}

	for _, contact := range contacts {
		result.List = append(result.List, jsonContact.ToContactResponse(contact))
	}

	c.JSON(http.StatusOK, result)
}

// ReadContactByID
// @Summary Method allow to get contact by ID
// @Description Method allow to get contact by ID
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path string true "Contact ID"
// @Success 200 {object} jsonContact.ContactResponse "Contact"
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /contacts/{id} [get]
func (d *Delivery) ReadContactByID(c *gin.Context) {
	var id jsonContact.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	response, err := d.ucContact.ReadByID(converter.StringToUUID(id.Value))
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, jsonContact.ToContactResponse(response))
}
