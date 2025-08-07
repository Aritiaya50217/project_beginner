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
	// ดึง Authorization header มา
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	// Get user ID from JWT token
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIDInt := userIDFromToken.(int)

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

	booking := domain.Booking{
		UserID:    uint(userIDInt),
		ItemID:    uint(req.ItemID),
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
	// ดึง Authorization header มา
	authHeader := c.GetHeader("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	token = strings.TrimSpace(token)

	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIDInt := userIDFromToken.(int)

	// ดึง booking ID
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

	if booking.UserID != uint(userIDInt) {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	userIDStr := strconv.FormatUint(uint64(booking.UserID), 10)

	user, err := utils.GetUserInfo(os.Getenv("USER_SERVICE_URL"), userIDStr, token)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to fetch user info"})
		return
	}

	var userInfo utils.UserResponse
	userInfo.ID = user.ID
	userInfo.Name = user.Name
	userInfo.Email = user.Email

	result := utils.BookingResponse{
		ID:        booking.ID,
		UserID:    userInfo,
		StartTime: utils.FormatDate(booking.StartTime),
		EndTime:   utils.FormatDate(booking.EndTime),
		CreateAt:  utils.FormatDate(booking.CreatedAt),
		UpdateAt:  utils.FormatDate(booking.UpdatedAt),
	}

	c.JSON(http.StatusOK, result)
}
