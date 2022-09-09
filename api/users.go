package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/meshuga/code-generation-playground/internal/service"
)

type Api struct {
	svc service.Service
}

// NewApi example
func NewApi(svc service.Service) *Api {
	return &Api{
		svc: svc,
	}
}

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
// @Success 200 {array} model.UsersCollection "ok"
// @Router  /admin/user/ [get]
func (a *Api) ListUsers(ctx *gin.Context) {
	users := a.svc.GetUsers()
	ctx.JSON(200, users)
}
