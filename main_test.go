package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	mocks "github.com/meshuga/code-generation-playground/internal/mocks/service"
	"github.com/meshuga/code-generation-playground/internal/model"
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
