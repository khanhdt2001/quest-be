package service

import (
	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/database"
)

type API interface {
	Setup(router *gin.RouterGroup)
}

func NewRouter(db *database.Database) *gin.Engine {
	router := gin.Default()

	apiRouteGroup := router.Group("/api/v1")
	NewAuthController(db).Setup(apiRouteGroup.Group("/auth"))
	return router
}
