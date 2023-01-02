package repository

import (
	"fmt"

	adsgo "github.com/Ivlay/ads-go"
	"github.com/jmoiron/sqlx"
)

type AdsPg struct {
	db *sqlx.DB
}

func NewAdsPg(db *sqlx.DB) *AdsPg {
	return &AdsPg{db: db}
}

func (r *AdsPg) GetAll(order, orderBy string) ([]adsgo.Advertisement, error) {
	var aa []adsgo.Advertisement
	if order == "" {
		order = "desc"
	}

	if orderBy == "" {
		orderBy = "created_at"
	}

	query := fmt.Sprintf(`
		select * from %s
		order by %s %s
	`, advertisementsTable, orderBy, order)

	err := r.db.Select(&aa, query)

	return aa, err
}

func (r *AdsPg) GetById(id int) (adsgo.Advertisement, error) {
	var ads adsgo.Advertisement

	query := fmt.Sprintf(`
		select * from %s
		where id = $1
	`, advertisementsTable)

	err := r.db.Get(&ads, query, id)
	if err != nil {
		return ads, err
	}

	return ads, nil
}

func (r *AdsPg) Create(adsInput adsgo.Advertisement) (int, error) {
	var id int

	query := fmt.Sprintf(`
		insert into %s (title, description, user_id, images)
		values ($1, $2, $3, $4)
		returning id
	`, advertisementsTable)

	row := r.db.QueryRow(query, adsInput.Title, adsInput.Description, adsInput.UserId, adsInput.Images)
	if err := row.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}

func (r *AdsPg) Delete(id int, userId int) error {
	query := fmt.Sprintf(`
		delete from %s
		where id = $1 and user_id = $2
	`, advertisementsTable)

	_, err := r.db.Exec(query, id, userId)
	return err
}
