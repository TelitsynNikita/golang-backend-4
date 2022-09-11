package service

import (
	todo "github.com/TelitsynNikita"
	"github.com/TelitsynNikita/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(user todo.User) (string, error)
	ParseToken(accessToken string) (int, string, error)
}

type Deals interface {
	Create(userId int, deal todo.Deal) (int, error)
	GetAllNew() ([]todo.AllNewDeals, error)
	GetAllOwnDeals(id int, role string, status string) ([]todo.AllNewDeals, error)
	GetOneDealById(id int) (todo.OneDeal, error)
	Delete(id int) error
	UpdateStatus(status string, id int) error
	UpdateDealBookkeeperId(userId int, requestId int) error
}

type Service struct {
	Authorization
	Deals
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Deals:         newTodoDealService(repos.Deals),
	}
}
