package service

import (
	adsgo "github.com/Ivlay/ads-go"
	"github.com/Ivlay/ads-go/internal/pkg/repository"
)

type AdsService struct {
	repo repository.Ads
}

func NewAdsService(repo repository.Ads) *AdsService {
	return &AdsService{repo: repo}
}

func (s *AdsService) Create(adsInput adsgo.Advertisement) (int, error) {
	return s.repo.Create(adsInput)
}

func (s *AdsService) GetAll(order, orderBy string) ([]adsgo.Advertisement, error) {
	return s.repo.GetAll(order, orderBy)
}

func (s *AdsService) GetByUserId(id int, order, orderBy string) ([]adsgo.Advertisement, error) {
	return s.repo.GetByUserId(id, order, orderBy)
}

func (s *AdsService) GetById(id int) (adsgo.Advertisement, error) {
	return s.repo.GetById(id)
}

func (s *AdsService) Delete(id, userId int) error {
	return s.repo.Delete(id, userId)
}
