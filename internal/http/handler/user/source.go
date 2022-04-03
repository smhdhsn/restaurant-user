package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smhdhsn/bookstore-user/internal/http/helper"
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
	"github.com/smhdhsn/bookstore-user/internal/validator"
	"github.com/smhdhsn/bookstore-user/util/response"

	uRequest "github.com/smhdhsn/bookstore-user/internal/request/user"
	uService "github.com/smhdhsn/bookstore-user/internal/service/user"
)

// Source contains services that can be used within user source handler.
type Source struct {
	sourceServ *uService.SourceService
}

// NewSource creates a new user source handler.
func NewSource(sourceServ *uService.SourceService) *Source {
	return &Source{
		sourceServ: sourceServ,
	}
}

// Find is responsible for finding a user inside the database.
func (h *Source) Find(c *gin.Context) {
	userID, err := helper.StrToUint(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	user, err := h.sourceServ.Find(userID)
	switch err {
	case nil:
		c.JSON(response.NewStatusOK(user))
	case contract.ErrRecordNotFound:
		c.JSON(response.NewStatusNotFound())
	default:
		c.JSON(response.NewStatusInternalServerError())
	}
}

// Store is responsible for storing a user in the database.
func (h *Source) Store(c *gin.Context) {
	req := new(uRequest.SourceStoreReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("error on binding JSON"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		data := helper.ParseValidationErr(err)
		c.JSON(response.NewStatusUnprocessableEntity(data))
		return
	}

	user, err := h.sourceServ.Store(*req)
	switch err {
	case nil:
		c.JSON(response.NewStatusCreated(user))
	case contract.ErrDuplicateEntry:
		c.JSON(response.NewStatusBadRequest("duplicate entry"))
	default:
		c.JSON(response.NewStatusInternalServerError())
	}
}

// Update is responsible for updating user's information inside database.
func (h *Source) Update(c *gin.Context) {
	req := new(uRequest.SourceUpdateReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("error on binding JSON"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		data := helper.ParseValidationErr(err)
		c.JSON(response.NewStatusUnprocessableEntity(data))
		return
	}

	userID, err := helper.StrToUint(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	err = h.sourceServ.Update(*req, userID)
	switch err {
	case nil:
		c.Status(http.StatusNoContent)
	case contract.ErrRecordNotFound:
		c.JSON(response.NewStatusNotFound())
	case contract.ErrDuplicateEntry:
		c.JSON(response.NewStatusBadRequest("duplicate entry"))
	default:
		c.JSON(response.NewStatusInternalServerError())
	}
}

// Destroy is responsible for deleting a user from the database.
func (h *Source) Destroy(c *gin.Context) {
	userID, err := helper.StrToUint(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	err = h.sourceServ.Destroy(userID)
	switch err {
	case nil:
		c.Status(http.StatusNoContent)
	case contract.ErrRecordNotFound:
		c.JSON(response.NewStatusNotFound())
	default:
		c.JSON(response.NewStatusInternalServerError())
	}
}
