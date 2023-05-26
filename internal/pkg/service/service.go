package service

import (
	adsgo "github.com/Ivlay/ads-go"
	"github.com/Ivlay/ads-go/internal/pkg/repository"
)

type User interface {
	Create(user adsgo.User) (adsgo.User, error)
	GenerateToken(claim int) (string, error)
	ParseToken(accessToken string) (int, error)
	Login(input adsgo.LoginInput) (adsgo.User, error)
	GetById(id int) (adsgo.User, error)
}

type Ads interface {
	GetAll(order, orderBy string) ([]adsgo.Advertisement, error)
	GetById(id int) (adsgo.Advertisement, error)
	Create(adsInput adsgo.Advertisement) (int, error)
	Delete(id, userId int) error
	GetByUserId(id int, order, orderBy string) ([]adsgo.Advertisement, error)
}

type Service struct {
	User
	Ads
}

func New(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
		Ads:  NewAdsService(repos.Ads),
	}
}
