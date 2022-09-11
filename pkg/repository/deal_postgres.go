package repository

import (
	"fmt"
	todo "github.com/TelitsynNikita"
	"github.com/jmoiron/sqlx"
)

type DealPostgres struct {
	db *sqlx.DB
}

func newDealPostgres(db *sqlx.DB) *DealPostgres {
	return &DealPostgres{db: db}
}

func (r *DealPostgres) Create(userId int, deal todo.Deal) (int, error) {
	var id int
	createRequestQuery := fmt.Sprintf("INSERT INTO %s (purpose, description, amount, user_id) values ($1, $2, $3, $4) RETURNING id", dealsTable)
	row := r.db.QueryRow(createRequestQuery, deal.Purpose, deal.Description, deal.Amount, userId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *DealPostgres) GetAllNew() ([]todo.AllNewDeals, error) {
	var deals []todo.AllNewDeals
	query := fmt.Sprintf("SELECT tl.id, tl.purpose, tl.amount, tl.status, tl.bookkeeper_id, ul.full_name FROM %s tl JOIN %s ul on tl.user_id = ul.id WHERE tl.status='NEW'",
		dealsTable, usersTable)
	err := r.db.Select(&deals, query)

	return deals, err
}

func (r *DealPostgres) GetOneDealById(dealId int) (todo.OneDeal, error) {
	var deal todo.OneDeal
	query := fmt.Sprintf("SELECT tl.id, tl.purpose, tl.amount, tl.status, tl.description, tl.user_id, tl.created_at, ul.full_name FROM %s tl JOIN %s ul on tl.user_id = ul.id WHERE tl.id = $1",
		dealsTable, usersTable)
	err := r.db.Get(&deal, query, dealId)

	return deal, err
}

func (r *DealPostgres) Delete(dealId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.id=$1", dealsTable)
	_, err := r.db.Exec(query, dealId)

	return err
}

func (r *DealPostgres) UpdateStatus(status string, id int) error {
	query := fmt.Sprintf("UPDATE %s tl SET status=$1  WHERE id=$2", dealsTable)
	_, err := r.db.Exec(query, status, id)

	return err
}

func (r *DealPostgres) UpdateDealBookkeeperId(userId int, requestId int) error {
	query := fmt.Sprintf("UPDATE %s tl SET bookkeeper_id=$1  WHERE id=$2", dealsTable)
	_, err := r.db.Exec(query, userId, requestId)

	return err
}

func (r *DealPostgres) GetAllOwnDeals(userId int, role string, status string) ([]todo.AllNewDeals, error) {
	var deals []todo.AllNewDeals
	var query string

	if role == "BOOKKEEPER" {
		query = fmt.Sprintf("SELECT tl.id, tl.purpose, tl.amount, tl.status, tl.bookkeeper_id, ul.full_name FROM %s tl JOIN %s ul on tl.user_id=ul.id WHERE tl.bookkeeper_id=$1 AND tl.status=$2",
			dealsTable, usersTable)
	} else if role == "USER" {
		query = fmt.Sprintf("SELECT tl.id, tl.purpose, tl.amount, tl.status, tl.bookkeeper_id, ul.full_name FROM %s tl JOIN %s ul on tl.user_id=ul.id WHERE tl.user_id=$1 AND tl.status=$2",
			dealsTable, usersTable)
	}

	err := r.db.Select(&deals, query, userId, status)

	fmt.Println(deals)

	return deals, err
}
