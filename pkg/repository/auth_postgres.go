package repository

import (
	"fmt"
	todo "github.com/TelitsynNikita"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (full_name, email, password) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.FullName, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(userData todo.User) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT tl.id, tl.full_name, tl.role, tl.email FROM %s tl WHERE email=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, userData.Email, userData.Password)

	return user, err
}
