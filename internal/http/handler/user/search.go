package user

import (
	"github.com/gin-gonic/gin"

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
}

// NewSearch creates a new user search handler.
func NewSearch(searchService *uService.SearchService) *Search {
	return &Search{
		searchService: searchService,
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

	userList, err := h.searchService.FindBy(*req)
	switch err {
	case nil:
		c.JSON(response.NewStatusOK(userList))
	case contract.ErrRecordNotFound:
		c.JSON(response.NewStatusNotFound())
	default:
		c.JSON(response.NewStatusInternalServerError())
	}
}
