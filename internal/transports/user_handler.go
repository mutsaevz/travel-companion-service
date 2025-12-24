package transports

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mutsaevz/team-5-ambitious/internal/dto"
	"github.com/mutsaevz/team-5-ambitious/internal/models"
	"github.com/mutsaevz/team-5-ambitious/internal/services"
)

type UserHandler struct {
	service services.UserService
	logger  *slog.Logger
}

func NewUserHandler(service services.UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (h UserHandler) RegisterRoutes(ctx *gin.Engine) {
	api := ctx.Group("/users")
	{
		api.POST("/", h.Create)
		api.GET("/", h.List)
		api.GET("/:id", h.GetByID)
		api.PATCH("/:id", h.Update)
		api.DELETE("/:id", h.Delete)
	}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	h.logger.Info("handler called",
		slog.String("method", ctx.Request.Method),
		slog.String("path", ctx.FullPath()),
	)

	var input dto.UserCreateRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warn("invalid JSON",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.FullPath()),
			slog.Any("error", err),
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	user, err := h.service.Create(&input)
	if err != nil {
		h.logger.Error("error adding user",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.FullPath()),
			slog.Any("error", err.Error()),
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error adding user"})
		return
	}

	h.logger.Info("user added successfully")

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) List(ctx *gin.Context) {
	h.logger.Info("handler called",
		slog.String("method", ctx.Request.Method),
		slog.String("path", ctx.FullPath()),
	)

	var filter models.Page

	if pageStr := ctx.Query("page"); pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err == nil {
			filter.Page = page
		}
	}

	if pageSizeStr := ctx.Query("pageSize"); pageSizeStr != "" {
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err == nil {
			filter.PageSize = pageSize
		}
	}

	users, err := h.service.List(filter)
	if err != nil {
		h.logger.Error("user list error",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.FullPath()),
			slog.Any("error", err.Error()),
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user list error"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetByID(ctx *gin.Context) {

	h.logger.Info("handler called",
		slog.String("method", ctx.Request.Method),
		slog.String("path", ctx.FullPath()),
	)

	idStr := ctx.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Error("handler called",
			slog.String("error", "invalid id"),
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"eror": "invalid id"})
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		h.logger.Error("handler called",
			slog.String("method", ctx.Request.Method),
			slog.String("error", "user not found"),
			slog.Any("error", err),
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	h.logger.Info("handler called",
		slog.String("method", ctx.Request.Method),
		slog.String("path", ctx.FullPath()),
	)

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Warn("invalid id",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.FullPath()),
			slog.Any("error", err),
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input dto.UserUpdateRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warn("invalid JSON",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.FullPath()),
			slog.Any("error", err),
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	updated, err := h.service.Update(uint(id), input)

	if err != nil {
		h.logger.Error("error saving changes",
			slog.Uint64("user_id", uint64(id)),
			slog.Any("error", err),
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error saving changes"})
		return
	}

	h.logger.Info("changes saved")
	ctx.JSON(http.StatusOK, updated)
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	h.logger.Info("handler called",
		slog.String("method", ctx.Request.Method),
		slog.String("path", ctx.FullPath()),
	)

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.logger.Warn("invalid id",
			slog.String("method", ctx.Request.Method),
			slog.String("path", ctx.FullPath()),
			slog.Any("error", err),
		)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		h.logger.Error("failed to delete user",
			slog.String("method", ctx.Request.Method),
			slog.Uint64("user_id", uint64(id)),
			slog.Any("error", err),
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	h.logger.Info("user deleted successfully")
	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
