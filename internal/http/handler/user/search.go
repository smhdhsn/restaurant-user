package user

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/smhdhsn/bookstore-user/internal/config"
	"github.com/smhdhsn/bookstore-user/internal/http/helper"
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
	"github.com/smhdhsn/bookstore-user/internal/validator"
	"github.com/smhdhsn/bookstore-user/util/response"

	uRequest "github.com/smhdhsn/bookstore-user/internal/request/user"
	uService "github.com/smhdhsn/bookstore-user/internal/service/user"
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

// List is responsible for finding every user with a given status from the database.
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
		}
		return
	}

	data := uList.ToExternalResp(h.hashConf)
	c.JSON(response.NewStatusOK(data))
}
