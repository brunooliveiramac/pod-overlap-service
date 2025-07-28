package httpapi

import (
	"github.com/brunooliveiramac/pod-overlap-service/internal/overlap"
	dto "github.com/brunooliveiramac/pod-overlap-service/internal/platform/http/dto"
	"github.com/brunooliveiramac/pod-overlap-service/internal/platform/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterRoutes(router *gin.Engine) {
	router.Use(RequestLogger())
	router.POST("/api/overlap", overlapHandler)
}

// RequestLogger returns a Gin middleware for logging HTTP requests.
func RequestLogger() gin.HandlerFunc {
	return gin.Logger()
}

func overlapHandler(c *gin.Context) {
	var req dto.OverlapRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse DTOs into domain types
	startRange, err := parseDateRangeDTO(req.StartRange)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_range: " + err.Error()})
		return
	}
	endRange, err := parseDateRangeDTO(req.EndRange)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_range: " + err.Error()})
		return
	}

	logger.Log.Info("[overlap] Checking: start_range={start: %s, end: %s}, end_range={start: %s, end: %s}", startRange.Start.Format(time.RFC3339), startRange.End.Format(time.RFC3339), endRange.Start.Format(time.RFC3339), endRange.End.Format(time.RFC3339))
	result := overlap.Overlaps(startRange, endRange)
	logger.Log.Info("[overlap] Result: %v", result)
	c.JSON(http.StatusOK, dto.OverlapResponseDTO{Overlap: result})
}

func parseDateRangeDTO(d dto.DateRangeDTO) (overlap.DateRange, error) {
	start, err := time.Parse(time.RFC3339, d.Start)
	if err != nil {
		return overlap.DateRange{}, err
	}
	end, err := time.Parse(time.RFC3339, d.End)
	if err != nil {
		return overlap.DateRange{}, err
	}
	return overlap.DateRange{Start: start, End: end}, nil
}
