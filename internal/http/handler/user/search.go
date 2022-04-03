package user

import (
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
