package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smhdhsn/bookstore-user/internal/service"
	"github.com/smhdhsn/bookstore-user/util/response"
	"gorm.io/gorm"
)

// UserHandler contains services that can be used within user handler.
type UserHandler struct {
	uServ *service.UserService
}

// NewUserHandler creates a new user handler.
func NewUserHandler(uServ *service.UserService) *UserHandler {
	return &UserHandler{
		uServ: uServ,
	}
}

// Store is responsible for storing a user in the database.
func (h *UserHandler) Store(c *gin.Context) {
	req := new(service.StoreUserReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("error on binding JSON"))
		return
	}

	switch user, err := h.uServ.Store(req); err {
	case nil:
		c.JSON(response.NewStatusCreated(user))

	default:
		c.JSON(response.NewStatusInternalServerError("error on storing user into database"))
	}
}

// Find is responsible for finding a user inside the database.
func (h *UserHandler) Find(c *gin.Context) {
	switch user, err := h.uServ.Find(c.Params.ByName("userID")); err {
	case nil:
		c.JSON(response.NewStatusOK(user))

	case service.ErrParseUint:
		c.JSON(response.NewStatusBadRequest(err.Error()))

	case gorm.ErrRecordNotFound:
		c.JSON(response.NewStatusNotFound("record not found"))

	default:
		c.JSON(response.NewStatusInternalServerError("error on finding user inside database"))
	}
}

// Update is responsible for updating user's information inside database.
func (h *UserHandler) Update(c *gin.Context) {
	req := new(service.UpdateUserReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("error on binding JSON"))
		return
	}

	switch err := h.uServ.Update(req, c.Params.ByName("userID")); err {
	case nil:
		c.Status(http.StatusNoContent)

	case service.ErrParseUint:
		c.JSON(response.NewStatusBadRequest(err.Error()))

	case gorm.ErrRecordNotFound:
		c.JSON(response.NewStatusNotFound("record not found"))

	default:
		c.JSON(response.NewStatusInternalServerError("error on updating user's information inside database"))
	}
}

// Destroy is responsible for deleting a user from the database.
func (h *UserHandler) Destroy(c *gin.Context) {
	switch err := h.uServ.Destroy(c.Params.ByName("userID")); err {
	case nil:
		c.Status(http.StatusNoContent)

	case service.ErrParseUint:
		c.JSON(response.NewStatusBadRequest(err.Error()))

	case gorm.ErrRecordNotFound:
		c.JSON(response.NewStatusNotFound("record not found"))

	default:
		c.JSON(response.NewStatusInternalServerError("error on deleting user from database"))
	}
}
