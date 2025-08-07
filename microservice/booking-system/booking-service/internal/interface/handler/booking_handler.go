package handler

import (
	"booking-system-booking-service/internal/app"
	"booking-system-booking-service/internal/domain"
	"booking-system-booking-service/internal/utils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	usecase app.BookingUsecase
}

func NewBookingHandler(usecase app.BookingUsecase) *BookingHandler {
	return &BookingHandler{usecase: usecase}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req utils.BookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}

	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}

	// Get user ID from JWT token
	userIDFromToken, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var userIDInt int
	switch v := userIDFromToken.(type) {
	case float64:
		userIDInt = int(v)
	case int:
		userIDInt = v
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id type"})
		return
	}

	booking := domain.Booking{
		UserID:    userIDInt,
		ItemID:    req.ItemID,
		StartTime: startTime,
		EndTime:   endTime,
	}

	if err := h.usecase.Create(c, &booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "booking created successfully"})
}

func (h *BookingHandler) GetByID(c *gin.Context) {
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var userIDInt int
	switch v := userIDFromToken.(type) {
	case float64:
		userIDInt = int(v)
	case int:
		userIDInt = v
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id type"})
		return
	}

	bookingIDParam := c.Param("id")
	bookingID, err := strconv.Atoi(bookingIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	booking, err := h.usecase.GetBookingByID(c.Request.Context(), bookingID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if booking.UserID != userIDInt {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	// user-service
	token := c.GetHeader("Authorization")

	if !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}

	userIDstr := strconv.Itoa(booking.UserID)
	user, err := utils.GetUserInfo(os.Getenv("USER_SERVICE_URL"), userIDstr, token)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch user info"})
		return
	}

	result := domain.Booking{
		ID:     booking.ID,
		UserID: user.ID,
		ItemID: booking.ItemID,
	}

	c.JSON(http.StatusOK, result)
}
