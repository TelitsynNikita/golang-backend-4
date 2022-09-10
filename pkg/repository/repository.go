package repository

import "github.com/jmoiron/sqlx"

type Authorization interface{}

type Deals interface{}

type Repository struct {
	Authorization
	Deals
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
