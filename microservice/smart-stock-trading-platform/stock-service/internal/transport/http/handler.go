package http

import (
	"net/http"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	usecase port.StockUsecase
}

func NewStockHandler(usecase port.StockUsecase) *StockHandler {
	return &StockHandler{usecase: usecase}
}

func (h *StockHandler) GetQuote(c *gin.Context) {
	symbol := c.Param("symbol")
	quote, err := h.usecase.FetchQuote(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quote)
}

func (h *StockHandler) GetAllQuote(c *gin.Context) {
	exchange := c.DefaultQuery("exchange", "US")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	quotes, err := h.usecase.FetchAllQuotes(exchange, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}

// เพิ่ม stock ใหม่ พร้อมสร้าง blockchain block
func (h *StockHandler) CreateStockWithBlock(c *gin.Context) {
	var stock domain.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if stock.Symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	err := h.usecase.AddStockWithBlock(stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "stock added with blockchain block"})
}
