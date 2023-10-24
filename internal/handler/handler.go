package handler

import (
	"github.com/dkshi/banklocsrv/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	departments := router.Group("/departments", h.userIdentity)
	{
		departments.POST("/atms", h.atms)
		departments.POST("/offices", h.offices)

	}

	return router
}
