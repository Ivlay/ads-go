package handler

import "github.com/Ivlay/ads-go/internal/pkg/service"

type Comments struct {
	services *service.Service
}

func NewCommentsService(services *service.Service) *Comments {
	return &Comments{
		services: services,
	}
}

func (h *Comments) createComment() {

}
