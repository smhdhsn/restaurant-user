package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/smhdhsn/bookstore-user/internal/db"
	"github.com/smhdhsn/bookstore-user/internal/request"
	"github.com/smhdhsn/bookstore-user/internal/service"
	"github.com/smhdhsn/bookstore-user/util/response"
	"gorm.io/gorm"
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

// Store is responsible for storing a user in the database.
func (h *UserHandler) Store(c *gin.Context) {
	req := new(request.StoreUserReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("error on binding JSON"))
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(response.NewStatusUnprocessableEntity(err))
		return
	}

	user, err := h.uServ.Store(req)
	if err != nil {
		if mErr, ok := err.(*mysql.MySQLError); ok {
			switch mErr.Number {
			case db.ErrDuplicateEntry:
				c.JSON(response.NewStatusBadRequest("email already exists"))
				return
			}
		}

		c.JSON(response.NewStatusInternalServerError("error on storing user into database"))
		return
	}

	c.JSON(response.NewStatusCreated(user))
}

// Find is responsible for finding a user inside the database.
func (h *UserHandler) Find(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Params.ByName("userID"), 10, 32)
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	user, err := h.uServ.Find(uint(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound("record not found"))
			return
		}

		c.JSON(response.NewStatusInternalServerError("error on finding user inside database"))
		return
	}

	c.JSON(response.NewStatusOK(user))
}

// Update is responsible for updating user's information inside database.
func (h *UserHandler) Update(c *gin.Context) {
	req := new(request.UpdateUserReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("error on binding JSON"))
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(response.NewStatusUnprocessableEntity(err))
		return
	}

	userID, err := strconv.ParseUint(c.Params.ByName("userID"), 10, 32)
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	err = h.uServ.Update(req, uint(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound("record not found"))
			return
		}

		if mErr, ok := err.(*mysql.MySQLError); ok {
			switch mErr.Number {
			case db.ErrDuplicateEntry:
				c.JSON(response.NewStatusBadRequest("email already exists"))
				return
			}
		}

		c.JSON(response.NewStatusInternalServerError("error on updating user's information inside database"))
		return
	}

	c.Status(http.StatusNoContent)
}

// Destroy is responsible for deleting a user from the database.
func (h *UserHandler) Destroy(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Params.ByName("userID"), 10, 32)
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	err = h.uServ.Destroy(uint(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound("record not found"))
			return
		}

		c.JSON(response.NewStatusInternalServerError("error on deleting user from database"))
		return
	}

	c.Status(http.StatusNoContent)
}
