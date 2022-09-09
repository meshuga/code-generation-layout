package service

import "github.com/meshuga/code-generation-playground/internal/app/example/model"

type Service interface {
	GetUsers() model.UsersCollection
}

type service struct {
}

func (s *service) GetUsers() model.UsersCollection {
	return model.UsersCollection{
		{
			ID:     1,
			Email:  "user@test.com",
			Status: model.Verified,
		},
	}
}

func NewService() Service {
	return &service{}
}
