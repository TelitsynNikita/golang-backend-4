package repository

import (
	todo "github.com/TelitsynNikita"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(user todo.User) (todo.User, error)
}

type Deals interface {
	Create(userId int, deal todo.Deal) (int, error)
	GetAllNew() ([]todo.AllNewDeals, error)
	GetAllOwnDeals(id int, role string, status string) ([]todo.AllNewDeals, error)
	GetOneDealById(id int) (todo.OneDeal, error)
	Delete(id int) error
	UpdateStatus(status string, id int) error
	UpdateDealBookkeeperId(userId int, dealId int) error
}

type Repository struct {
	Authorization
	Deals
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Deals:         newDealPostgres(db),
	}
}
