package user

import (
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"
)

// SearchService contains repositories that will be used within this service.
type SearchService struct {
	uRepo contract.UserRepository
}

// NewSearchService creates a user's search service with it's dependencies.
func NewSearchService(uRepo contract.UserRepository) *SearchService {
	return &SearchService{
		uRepo: uRepo,
	}
}
