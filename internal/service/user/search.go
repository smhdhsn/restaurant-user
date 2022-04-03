package user

import (
	"github.com/smhdhsn/bookstore-user/internal/repository/contract"

	uRequest "github.com/smhdhsn/bookstore-user/internal/request/user"
	uTransform "github.com/smhdhsn/bookstore-user/internal/transoform/user"
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

// FindBy searches for records matching a given value and fetches them from database.
func (s *SearchService) FindBy(req uRequest.SearchListReq) ([]*uTransform.SearchList, error) {
	filter := contract.FilterBy{
		Field: "status",
		Value: req.Status,
	}

	uList, err := s.uRepo.FindBy(filter)
	if err != nil {
		return nil, err
	}

	if len(uList) == 0 {
		return nil, contract.ErrRecordNotFound
	}

	transform := make([]*uTransform.SearchList, 0)
	for _, u := range uList {
		transform = append(transform, &uTransform.SearchList{
			ID:        u.ID,
			Status:    u.Status,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return transform, nil
}
