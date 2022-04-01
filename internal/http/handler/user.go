package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/smhdhsn/bookstore-user/internal/http/helper"
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
	"github.com/smhdhsn/bookstore-user/internal/request"
	"github.com/smhdhsn/bookstore-user/internal/service"
	"github.com/smhdhsn/bookstore-user/util/response"
)

// validate holds the validator's functionalities.
var validate *validator.Validate

// UserHandler contains services that can be used within user handler.
type UserHandler struct {
	uServ *service.UserService
}

// init will be executed when this package is imported.
func init() {
	validate = validator.New()
}

// NewUserHandler creates a new user handler.
func NewUserHandler(uServ *service.UserService) *UserHandler {
	return &UserHandler{
		uServ: uServ,
	}
}

// Find is responsible for finding a user inside the database.
func (h *UserHandler) Find(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Params.ByName("userID"), 10, 32)
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	user, err := h.uServ.Find(uint(userID))
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
func (h *UserHandler) Store(c *gin.Context) {
	req := new(request.StoreUserReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("error on binding JSON"))
		return
	}

	if err := validate.Struct(req); err != nil {
		data := helper.ParseValidationErr(err)
		c.JSON(response.NewStatusUnprocessableEntity(data))
		return
	}

	user, err := h.uServ.Store(req)
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
func (h *UserHandler) Update(c *gin.Context) {
	req := new(request.UpdateUserReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("error on binding JSON"))
		return
	}

	if err := validate.Struct(req); err != nil {
		data := helper.ParseValidationErr(err)
		c.JSON(response.NewStatusUnprocessableEntity(data))
		return
	}

	userID, err := strconv.ParseUint(c.Params.ByName("userID"), 10, 32)
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	err = h.uServ.Update(req, uint(userID))
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
func (h *UserHandler) Destroy(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Params.ByName("userID"), 10, 32)
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	err = h.uServ.Destroy(uint(userID))
	switch err {
	case nil:
		c.Status(http.StatusNoContent)
	case contract.ErrRecordNotFound:
		c.JSON(response.NewStatusNotFound())
	default:
		c.JSON(response.NewStatusInternalServerError())
	}
}
