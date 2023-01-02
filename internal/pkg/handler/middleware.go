package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Ivlay/ads-go/internal/pkg/service"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	services *service.Service
}

func NewMiddleware(services *service.Service) *Middleware {
	return &Middleware{
		services: services,
	}
}

func (h *Middleware) UserIdentify(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	headerParts := strings.Split(header, " ")

	token := headerParts[1]
	headerName := headerParts[0]

	if headerName != "Bearer" || token == "" {
		newErrorResponse(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	userId, err := h.services.User.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id ivalid type")
		return 0, errors.New("user id invalid type")
	}

	return idInt, nil
}
