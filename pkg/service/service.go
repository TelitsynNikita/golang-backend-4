package service

import "github.com/TelitsynNikita/pkg/repository"

type Authorization interface{}

type Deals interface{}

type Service struct {
	Authorization
	Deals
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
