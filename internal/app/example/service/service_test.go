package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Service_GetUsers(t *testing.T) {
	svc := NewService()

	users := svc.GetUsers()

	assert.Len(t, users, 1)

}
