package http

import (
	"net/http"
	"smart-stock-trading-platform-order-service/internal/domain"
	"smart-stock-trading-platform-order-service/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase usecase.OrderUsecase
}

func NewOrderHandler(usecase usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecase: usecase}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	// Get user ID from JWT token
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIDInt := userIDFromToken.(int)

	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.UserID = uint(userIDInt)
	order.Status = domain.Pending
	order.Price = float64(order.Quantity) * float64(order.Price)

	if err := h.usecase.CreateOrder(c, &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.usecase.GetOrder(c, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrdersByUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	orders, total, err := h.usecase.GetOrdersByUser(c, uint(userID), int64(offset), int64(limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       orders,
		"total":      total,
		"offset":     offset,
		"limit":      limit,
		"totalPages": (total + int64(limit) - 1) / int64(limit), // คำนวณจำนวนหน้า
	})
}
