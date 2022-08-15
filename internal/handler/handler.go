package handler

import (
	_ "github.com/AndrewMislyuk/payments-api/docs"
	"github.com/AndrewMislyuk/payments-api/internal/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Payments interface {
	ProductSubscription(productId string) (string, error)
}

type Handler struct {
	paymentService Payments
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		paymentService: services.Payments,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api")
	{
		api.POST("/subscribe", h.productSubscribe)
	}

	return router
}
