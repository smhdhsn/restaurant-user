package request

// StoreUserReq is responsible for holding user's information to be stored into database.
type StoreUserReq struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Status    string `json:"status" validate:"required,oneof=active inactive"`
}

// UpdateUserReq is responsible for holding user's data to be updated inside database.
type UpdateUserReq struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Status    string `json:"status" validate:"required,oneof=active inactive"`
}
