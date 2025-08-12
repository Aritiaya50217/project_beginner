package http

import (
	"net/http"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"
	"smart-stock-trading-platform-stock-service/internal/utils"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	usecase      port.OrderUsecase
	stockUsecase port.StockUsecase
}

func NewOrderHandler(usecase port.OrderUsecase, stockUsecase port.StockUsecase) *orderHandler {
	return &orderHandler{usecase: usecase, stockUsecase: stockUsecase}
}

func (h *orderHandler) CreateOrder(c *gin.Context) {
	// ดึง userID จาก context (เซ็ตโดย middleware)
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req utils.Order

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stock, err := h.stockUsecase.FindStockByID(c.Request.Context(), req.StockID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid stock id"})
		return
	}

	// price ต้องมาจาก stock.LastPrice * req.total
	order := domain.Order{
		UserID:     userIDFromToken.(int),
		StockID:    req.StockID,
		TotalStock: req.TotalStock,
		TotalPrice: stock.LastPrice * float64(req.TotalStock),
	}

	if err := h.usecase.CreateOrder(c.Request.Context(), &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erroror": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": order.ID})

}
