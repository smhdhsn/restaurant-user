package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/smhdhsn/restaurant-user/internal/config"
	"github.com/smhdhsn/restaurant-user/internal/http/helper"
	"github.com/smhdhsn/restaurant-user/internal/repository/contract"
	"github.com/smhdhsn/restaurant-user/internal/validator"
	"github.com/smhdhsn/restaurant-user/util/response"

	log "github.com/smhdhsn/restaurant-user/internal/logger"

	uRequest "github.com/smhdhsn/restaurant-user/internal/request/user"
	uService "github.com/smhdhsn/restaurant-user/internal/service/user"
)

// Search contains services that can be used within user search's handler.
type Search struct {
	searchService *uService.SearchService
	hashConf      config.HashConf
}

// NewSearch creates a new user search handler.
func NewSearch(searchService *uService.SearchService, hashConf config.HashConf) *Search {
	return &Search{
		searchService: searchService,
		hashConf:      hashConf,
	}
}

// List is responsible for fetching user's full details with a given status from the database.
func (h *Search) List(c *gin.Context) {
	req := new(uRequest.SearchListReq)
	req.Status = c.Query("status")

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		data := helper.ParseValidationErr(err)
		c.JSON(response.NewStatusUnprocessableEntity(data))
		return
	}

	uList, err := h.searchService.FindBy(*req)
	if err != nil {
		if errors.Is(err, contract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else {
			c.JSON(response.NewStatusInternalServerError())
			log.Error(err)
		}

		return
	}

	data := uList.ToInternalResp()
	c.JSON(response.NewStatusOK(data))
}

// Index is responsible for fetching limited user data for every user with a given status from the database.
func (h *Search) Index(c *gin.Context) {
	req := new(uRequest.SearchListReq)
	req.Status = c.Query("status")

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		data := helper.ParseValidationErr(err)
		c.JSON(response.NewStatusUnprocessableEntity(data))
		return
	}

	uList, err := h.searchService.FindBy(*req)
	if err != nil {
		if errors.Is(err, contract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else {
			c.JSON(response.NewStatusInternalServerError())
			log.Error(err)
		}

		return
	}

	data := uList.ToExternalResp(h.hashConf)
	c.JSON(response.NewStatusOK(data))
}
