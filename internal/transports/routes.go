package transports

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/mutsaevz/team-5-ambitious/internal/services"
)

func RegisterRoutes(
	routes *gin.Engine,
	logger *slog.Logger,
	userService services.UserService,
) {
	userHandler := NewUserHandler(userService, logger)

	userHandler.RegisterRoutes(routes)
}
