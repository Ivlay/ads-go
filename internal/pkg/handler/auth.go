package handler

import (
	"database/sql"
	"errors"
	"net/http"

	adsgo "github.com/Ivlay/ads-go"
	"github.com/Ivlay/ads-go/internal/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

type Auth struct {
	services *service.Service
}

func NewAuthHandler(services *service.Service) *Auth {
	return &Auth{
		services: services,
	}
}

func (h *Auth) SignUp(c *gin.Context) {
	var input adsgo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.Create(input)
	if err != nil {
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				newErrorResponse(c, http.StatusBadRequest, pgErr.Message)
			default:
				newErrorResponse(c, http.StatusBadRequest, err.Error())
			}
		} else {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	token, err := h.services.User.GenerateToken(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":  user,
		"token": token,
	})
}

func (h *Auth) SingIn(c *gin.Context) {
	var input adsgo.LoginInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.User.Login(input)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			newErrorResponse(c, http.StatusBadRequest, "Wrong credentials")
		default:
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	token, err := h.services.User.GenerateToken(user.Id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"user":  user,
	})
}
