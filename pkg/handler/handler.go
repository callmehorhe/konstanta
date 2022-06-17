package handler

import (
	"github.com/callmehorhe/konstanta/pkg/service"
	"github.com/gin-gonic/gin"
) 

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	
	router.POST("/create", h.Create)
	router.PUT("/update", h.UpdateStatus)
	router.GET("/status/:tid", h.GetStatusByID)
	router.POST("/payments/id", h.GetAllTransactionsByUserID)
	router.POST("/payments/email", h.GetAllTransactionsByEmail)
	router.PUT("/cancel", h.CancelTransaction)

	return router
}
