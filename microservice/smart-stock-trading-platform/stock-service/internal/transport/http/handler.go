package http

import (
	"net/http"
	"smart-stock-trading-platform-stock-service/internal/port"

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
	symbols := c.QueryArray("symbol")
	if len(symbols) == 0 {
		symbols = []string{"AAPL", "GOOGL", "MSFT"} // default
	}

	quotes, err := h.usecase.FetchQuotes(symbols)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}
