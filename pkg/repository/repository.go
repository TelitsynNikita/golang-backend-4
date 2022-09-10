package repository

import (
	todo "github.com/TelitsynNikita"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(user todo.User) (todo.User, error)
}

type Deals interface{}

type Repository struct {
	Authorization
	Deals
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
