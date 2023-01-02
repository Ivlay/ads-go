package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	adsgo "github.com/Ivlay/ads-go"
	"github.com/Ivlay/ads-go/internal/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

type UserService struct {
	repo repository.User
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user adsgo.User) (adsgo.User, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) GenerateToken(claim int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		claim,
	})

	return token.SignedString([]byte(os.Getenv("SIGN_KEY")))
}

func (s *UserService) Login(input adsgo.LoginInput) (adsgo.User, error) {
	input.Password = generatePasswordHash(input.Password)
	return s.repo.Login(input)
}

func (s *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SIGN_KEY")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}
