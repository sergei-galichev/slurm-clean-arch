package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slurm-clean-arch/pkg/tools/converter"
	"slurm-clean-arch/pkg/type/pagination"
	"slurm-clean-arch/pkg/type/query"
	"slurm-clean-arch/pkg/type/queryparameter"
	jsonGroup "slurm-clean-arch/services/contact/internal/delivery/http/group"
	domainGroup "slurm-clean-arch/services/contact/internal/domain/group"
	"slurm-clean-arch/services/contact/internal/domain/group/description"
	"slurm-clean-arch/services/contact/internal/domain/group/name"
	"time"
)

var mappingSortsGroup = query.SortsOptions{
	"id":           {},
	"name":         {},
	"description":  {},
	"contactCount": {},
}

// CreateGroup
// @Summary Method allow to create contact group
// @Description Method allow to create contact group
// @Tags groups
// @Accept json
// @Produce json
// @Param group body jsonGroup.ShortGroup true "Group data"
// @Success 200 {object} jsonGroup.ResponseGroup "Group structure"
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /groups/ [post]
func (d *Delivery) CreateGroup(c *gin.Context) {
	var group = &jsonGroup.ShortGroup{}

	if err := c.ShouldBindJSON(group); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groupName, err := name.New(group.Name)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groupDescription, err := description.New(group.Description)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	newGroup, err := d.ucGroup.Create(
		domainGroup.New(
			groupName,
			groupDescription,
		),
	)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	response := jsonGroup.ResponseGroup{
		ID:         newGroup.ID().String(),
		CreatedAt:  newGroup.CreatedAt(),
		ModifiedAt: newGroup.ModifiedAt(),
		Group: jsonGroup.Group{
			ShortGroup: jsonGroup.ShortGroup{
				Name:        newGroup.Name().String(),
				Description: newGroup.Description().String(),
			},
			ContactsAmount: newGroup.ContactCount(),
		},
	}

	c.JSON(http.StatusOK, response)
}

// UpdateGroup
// @Summary Method allow to update contact group
// @Description Method allow to update contact group
// @Tags groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Param group body jsonGroup.ShortGroup true "Group data"
// @Success 200 {object} jsonGroup.ResponseGroup "Group structure"
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /groups/{id} [put]
func (d *Delivery) UpdateGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	group := jsonGroup.ShortGroup{}
	if err := c.ShouldBindJSON(&group); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groupName, err := name.New(group.Name)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groupDescription, err := description.New(group.Description)
	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	response, err := d.ucGroup.Update(
		domainGroup.NewWithID(
			converter.StringToUUID(id.Value),
			time.Now().UTC(),
			time.Now().UTC(),
			groupName,
			groupDescription,
			0,
		),
	)
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, jsonGroup.ProtoToGroupResponse(response))
}

// DeleteGroup
// @Summary Method allow to delete contact group
// @Description Method allow to delete contact group
// @Tags groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} string
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /groups/{id} [delete]
func (d *Delivery) DeleteGroup(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	if err := d.ucGroup.Delete(converter.StringToUUID(id.Value)); err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// ListGroup
// @Summary Method allow to list contact groups
// @Description Method allow to list contact groups
// @Tags groups
// @Accept json
// @Produce json
// @Param limit query int false "Records count limit" default(10) minimum(0) maximum(100)
// @Param offset query int false "Get records offset" default(0) minimum(0)
// @Param sort query string false "Sort by field" default(name)
// @Success 200 {object} jsonGroup.ListGroup "List of groups"
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /groups/ [get]
func (d *Delivery) ListGroup(c *gin.Context) {
	params, err := query.ParseQuery(
		c, query.Options{
			Sorts: mappingSortsGroup,
		},
	)

	if err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	groups, err := d.ucGroup.List(
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

	count, err := d.ucGroup.Count()
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	var list = make([]*jsonGroup.ResponseGroup, len(groups))

	for i, elem := range groups {
		list[i] = jsonGroup.ProtoToGroupResponse(elem)
	}

	response := jsonGroup.ListGroup{
		Total:  count,
		Limit:  params.Limit,
		Offset: params.Offset,
		List:   list,
	}

	c.JSON(http.StatusOK, response)
}

// ReadGroupByID
// @Summary Method allow to read contact group by id
// @Description Method allow to read contact group by id
// @Tags groups
// @Accept json
// @Produce json
// @Param id path string true "Group ID"
// @Success 200 {object} jsonGroup.ResponseGroup "Group structure"
// @Failure 400 {object} ErrorResponse
// @Failure 403 "403 Forbidden"
// @Failure 404 {object} ErrorResponse "404 Not Found"
// @Router /groups/{id} [get]
func (d *Delivery) ReadGroupByID(c *gin.Context) {
	var id jsonGroup.ID
	if err := c.ShouldBindUri(&id); err != nil {
		SetError(c, http.StatusBadRequest, err)
		return
	}

	response, err := d.ucGroup.ReadByID(converter.StringToUUID(id.Value))
	if err != nil {
		SetError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, jsonGroup.ProtoToGroupResponse(response))
}
