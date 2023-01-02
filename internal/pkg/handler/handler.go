package handler

import (
	"github.com/Ivlay/ads-go/internal/pkg/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	SignUp(c *gin.Context)
	SingIn(c *gin.Context)
}

type AdsHandler interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
}

type MiddlewareHandler interface {
	UserIdentify(c *gin.Context)
}

type Handler struct {
	AuthHandler
	AdsHandler
	MiddlewareHandler
}

func New(service *service.Service) *Handler {
	return &Handler{
		AuthHandler:       NewAuthHandler(service),
		AdsHandler:        NewAdsHandler(service),
		MiddlewareHandler: NewMiddleware(service),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/sign-up", h.AuthHandler.SignUp)
			v1.POST("/sign-in", h.AuthHandler.SingIn)
		}

		ads := v1.Group("/ads")
		{
			ads.GET("", h.AdsHandler.GetAll)
			ads.GET("/:id", h.AdsHandler.GetById)
			ads.POST("", h.MiddlewareHandler.UserIdentify, h.AdsHandler.Create)
			ads.DELETE("/:id", h.MiddlewareHandler.UserIdentify, h.AdsHandler.Delete)
		}

	}

	return router
}
