package handler

import (
	"net/http"

	"github.com/AndrewMislyuk/payments-api/internal/domain"
	"github.com/gin-gonic/gin"
)

type getSubscribeURL struct {
	URL string `json:"url"`
}

// @Summary Product Subscribe
// @Tags Subscribe
// @Description product-subscribe
// @ID product-subscribe
// @Accept  json
// @Produce  json
// @Param input body domain.Product true "Product ID"
// @Success 200 {object} getSubscribeURL
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/subscribe [post]
func (h *Handler) productSubscribe(c *gin.Context) {
	var input domain.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	url, err := h.paymentService.ProductSubscription(input.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, getSubscribeURL{
		URL: url,
	})
}
