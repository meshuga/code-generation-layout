package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Api struct {
}

// NewApi example
func NewApi() *Api {
	return &Api{}
}

// User example
type User struct {
	ID       int64
	Email    string
	Password string
}

// UsersCollection example
type UsersCollection []User

// Error example
type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

// ListUsers example
// @Summary List users from the store
// @Tags    admin
// @Accept  json
// @Produce json
// @Success 200 {array} api.UsersCollection "ok"
// @Router  /admin/user/ [get]
func (a *Api) ListUsers(ctx *gin.Context) {
	// write your code
}
