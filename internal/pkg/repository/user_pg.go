package repository

import (
	"fmt"

	adsgo "github.com/Ivlay/ads-go"
	"github.com/jmoiron/sqlx"
)

type UserPg struct {
	db *sqlx.DB
}

func NewUserPG(db *sqlx.DB) *UserPg {
	return &UserPg{db: db}
}

func (r *UserPg) CreateUser(inputUser adsgo.User) (adsgo.User, error) {
	var user adsgo.User
	query := fmt.Sprintf(`
		insert into %s (username, name, password)
		values ($1, $2, $3)
		returning id, name, username`,
		usersTable,
	)
	row := r.db.QueryRow(query, inputUser.Username, inputUser.Name, inputUser.Password)
	if err := row.Scan(&user.Id, &user.Name, &user.Username); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserPg) Login(input adsgo.LoginInput) (adsgo.User, error) {
	var user adsgo.User

	query := fmt.Sprintf("select id, username, name from %s where username=$1 AND password=$2", usersTable)

	err := r.db.Get(&user, query, input.Username, input.Password)

	return user, err
}
