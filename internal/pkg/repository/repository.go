package repository

import (
	adsgo "github.com/Ivlay/ads-go"
	"github.com/jmoiron/sqlx"
)

const (
	advertisementsTable = "advertisements"
	usersTable          = "users"
)

type User interface {
	CreateUser(user adsgo.User) (adsgo.User, error)
	Login(input adsgo.LoginInput) (adsgo.User, error)
	GetById(id int) (adsgo.User, error)
}

type Comments interface {
	CreateComment() (adsgo.Comments, error)
	DeleteComment(id int) error
	GetCommentsByAdsId(id int) ([]adsgo.Comments, error)
}

type Ads interface {
	GetAll(order, orderBy string) ([]adsgo.Advertisement, error)
	GetById(id int) (adsgo.Advertisement, error)
	Create(adsInput adsgo.Advertisement) (int, error)
	Delete(id, userId int) error
}

type Repository struct {
	User
	Ads
	Comments
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		User:     NewUserPG(db),
		Ads:      NewAdsPg(db),
		Comments: NewCommentsPg(db),
	}
}
