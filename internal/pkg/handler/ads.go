package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	adsgo "github.com/Ivlay/ads-go"
	"github.com/Ivlay/ads-go/internal/pkg/service"
	"github.com/gin-gonic/gin"
)

type Ads struct {
	services *service.Service
}

func NewAdsHandler(services *service.Service) *Ads {
	return &Ads{
		services: services,
	}
}

func (h *Ads) GetAll(c *gin.Context) {
	order := c.Query("order")
	orderBy := c.Query("orderBy")

	adsList, err := h.services.Ads.GetAll(order, orderBy)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, adsList)
}

func (h *Ads) GetByUserId(c *gin.Context) {
	order := c.Query("order")
	orderBy := c.Query("orderBy")

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	adsList, err := h.services.Ads.GetByUserId(userId, order, orderBy)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, adsList)
}

func (h *Ads) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ads id params")
		return
	}

	ads, err := h.services.Ads.GetById(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			newErrorResponse(c, http.StatusNotFound, "Ads not found")
		default:
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, ads)
}

func (h *Ads) Create(c *gin.Context) {
	var input adsgo.Advertisement

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.UserId = userId

	ads, err := h.services.Ads.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, ads)
}

func (h *Ads) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ads id params")
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Ads.Delete(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
}
