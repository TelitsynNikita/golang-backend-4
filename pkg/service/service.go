package service

import (
	todo "github.com/TelitsynNikita"
	"github.com/TelitsynNikita/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(user todo.User) (string, error)
}

type Deals interface{}

type Service struct {
	Authorization
	Deals
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
