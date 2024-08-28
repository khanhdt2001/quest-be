package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/repository/postgres"
)

type API interface {
	Setup(router *gin.RouterGroup)
}

func NewRouter(db *postgres.Database) *gin.Engine {
	router := gin.Default()

	apiRouteGroup := router.Group("/api/v1")
	NewAuthRouter(db).Setup(apiRouteGroup.Group("/auth"))
	return router
}
