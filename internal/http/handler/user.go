package handler

import "github.com/gin-gonic/gin"

// UserHandler contains services that can be used within user handler.
type UserHandler struct {
}

// NewUserHandler creates a new user handler.
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// Create is responsible for storing a user in the database.
func (h *UserHandler) Create(c *gin.Context) {
	//
}

// Get is responsible for getting a user from the database.
func (h *UserHandler) Get(c *gin.Context) {
	//
}

// Update is responsible for updating user's information in the database.
func (h *UserHandler) Update(c *gin.Context) {
	//
}

// Delete is responsible for deleting a user from the database.
func (h *UserHandler) Delete(c *gin.Context) {
	//
}

// Search is responsible for searching for user in the database.
func (h *UserHandler) Search(c *gin.Context) {
	//
}

// Login is responsible for user's authentication.
func (h *UserHandler) Login(c *gin.Context) {
	//
}
