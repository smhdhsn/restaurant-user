package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/http/helper"
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
	"github.com/smhdhsn/bookstore-user/internal/validator"
	"github.com/smhdhsn/bookstore-user/util/encryption"
	"github.com/smhdhsn/bookstore-user/util/response"

	uRequest "github.com/smhdhsn/bookstore-user/internal/request/user"
	uService "github.com/smhdhsn/bookstore-user/internal/service/user"
)

// Source contains services that can be used within user source handler.
type Source struct {
	sourceServ *uService.SourceService
	hashConf   config.HashConf
}

// NewSource creates a new user source handler.
func NewSource(sourceServ *uService.SourceService, hashConf config.HashConf) *Source {
	return &Source{
		sourceServ: sourceServ,
		hashConf:   hashConf,
	}
}

// Find is responsible for fetching user's full details from database.
func (h *Source) Find(c *gin.Context) {
	userID, err := helper.StrToUint(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	user, err := h.sourceServ.Find(userID)
	if err != nil {
		if errors.Is(err, contract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else {
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	data := user.ToInternalResp()
	c.JSON(response.NewStatusOK(data))
}

// Show is responsible for fetching user's limited details from database.
func (h *Source) Show(c *gin.Context) {
	userID, err := encryption.DecodeHashIDs(
		c.Params.ByName("userCode"),
		h.hashConf.Alphabet,
		h.hashConf.Salt,
		h.hashConf.MinLength,
	)
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on decoding userCode"))
		return
	}

	user, err := h.sourceServ.Show(userID)
	if err != nil {
		if errors.Is(err, contract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else {
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	data := user.ToExternalResp(h.hashConf)
	c.JSON(response.NewStatusOK(data))
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
	if err != nil {
		if errors.Is(err, contract.ErrDuplicateEntry) {
			c.JSON(response.NewStatusBadRequest("duplicate entry"))
		} else {
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	data := user.ToInternalResp()
	c.JSON(response.NewStatusCreated(data))
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
	if err != nil {
		if errors.Is(err, contract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else if errors.Is(err, contract.ErrDuplicateEntry) {
			c.JSON(response.NewStatusBadRequest("duplicate entry"))
		} else {
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// Destroy is responsible for deleting a user from the database.
func (h *Source) Destroy(c *gin.Context) {
	userID, err := helper.StrToUint(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("error on parsing userID"))
		return
	}

	err = h.sourceServ.Destroy(userID)
	if err != nil {
		if errors.Is(err, contract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else {
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	c.Status(http.StatusNoContent)
}
