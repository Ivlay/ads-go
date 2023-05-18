package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Ivlay/ads-go/internal/pkg/service"
	"github.com/gin-gonic/gin"
)

type User struct {
	services *service.Service
}

func NewUserHandler(services *service.Service) *User {
	return &User{
		services: services,
	}
}

func (h *User) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user id params")
		return
	}

	user, err := h.services.User.GetById(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			newErrorResponse(c, http.StatusNotFound, "User not found")
		default:
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
