package http

import (
	"net/http"
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

func (h *StockHandler) GetCompany(c *gin.Context) {
	symbol := c.Param("symbol")
	company, err := h.usecase.FetchCompayny(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid symbol"})
		return
	}
	c.JSON(http.StatusOK, company)
}
