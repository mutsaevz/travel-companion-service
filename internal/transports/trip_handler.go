package transports

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mutsaevz/team-5-ambitious/internal/dto"
	"github.com/mutsaevz/team-5-ambitious/internal/repository"
	"github.com/mutsaevz/team-5-ambitious/internal/services"
)

type TripHandler struct {
	service services.TripService
	logger  *slog.Logger
}

func NewTripHandler(service services.TripService, logger *slog.Logger) *TripHandler {
	return &TripHandler{
		service: service,
		logger:  logger,
	}
}

func (h *TripHandler) RegisterRoutes(ctx *gin.Engine) {
	api := ctx.Group("/trips")
	{
		api.POST("/driver/:driverID", h.Create)
		api.GET("/", h.List)
		api.GET("/:id", h.GetByID)
		api.PUT("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}
}

func (h *TripHandler) Create(ctx *gin.Context) {
	var req dto.TripCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	idStr := ctx.Param("driverID")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	trip, err := h.service.Create(uint(id), &req)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "driver not found"})
			return
		}
		h.logger.Error("failed to create trip", slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, trip)
}

func (h *TripHandler) List(ctx *gin.Context) {
	var filter dto.TripFilter

	if from := ctx.Query("fromCity"); from != "" {
		filter.FromCity = &from
	}

	if to := ctx.Query("toCity"); to != "" {
		filter.ToCity = &to
	}

	if timeStr := ctx.Query("startTime"); timeStr != "" {
		time, err := time.Parse(time.RFC3339, timeStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid time format"})
			return
		}
		filter.StartTime = &time
	}

	if pageStr := ctx.Query("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil {
			filter.Page = page
		}
	}

	if pagesizeStr := ctx.Query("pageSize"); pagesizeStr != "" {
		if pageSize, err := strconv.Atoi(pagesizeStr); err == nil {
			filter.PageSize = pageSize
		}
	}

	list, err := h.service.List(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *TripHandler) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	trip, err := h.service.GetByID(uint(id))
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
			return
		}
		h.logger.Error("failed to get trip", slog.Uint64("trip_id", uint64(id)), slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, trip)
}

func (h *TripHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.TripUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	trip, err := h.service.Update(uint(id), req)
	if err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
			return
		}
		h.logger.Error("failed to update trip", slog.Uint64("trip_id", uint64(id)), slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, trip)
}

func (h *TripHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		if err == repository.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "trip not found"})
			return
		}
		h.logger.Error("failed to delete trip", slog.Uint64("trip_id", uint64(id)), slog.Any("error", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
