package http

import (
	"log"
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

func (h *StockHandler) AddStock(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid symbol"})
		return
	}

	if err := h.usecase.AddStockBySymbol(c, symbol); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add stock successfully"})
}

func (h *StockHandler) GetStockByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// ดึง userID จาก context (เซ็ตโดย middleware)
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// เช็คว่า userID ใน token ต้องตรงกับ param id
	if userIDFromToken.(int) != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: you can only access your own user data"})
		return
	}

	stock, err := h.usecase.FindStockByID(c, id)
	if err != nil {
		log.Printf("error: %+v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "id is not found."})
		return
	}
	c.JSON(http.StatusOK, stock)
}
