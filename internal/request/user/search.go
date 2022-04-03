package request

// SearchListReq holds user search's request schema for finding users from database.
type SearchListReq struct {
	Status string `form:"status" validate:"required,oneof=active inactive"`
}
