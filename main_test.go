package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/meshuga/code-generation-playground/internal/app/example/model"
	mocks "github.com/meshuga/code-generation-playground/internal/mocks/app/example/service"
)

func TestPingRoute(t *testing.T) {

	svc := mocks.NewService(t)
	svc.On("GetUsers").Return(model.UsersCollection{
		{
			ID: 1,
		},
	})

	router := setupRouter(svc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
