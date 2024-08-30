package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quest-be/internal/repository/postgres"
	"github.com/quest-be/internal/service/controller"
	"github.com/quest-be/internal/service/handler"
)

func NewRouter(db *postgres.Database) *gin.Engine {
	router := gin.Default()

	apiRouteGroup := router.Group("/api/v1")

	userHandler := handler.NewUserHandler(db)
	OtpHandler := handler.NewOtpHandler(db)
	// setup auth router
	authHandler := handler.NewAuthHandler(db)
	authController := controller.NewAuthController()
	autherRouter := NewAuthRouter()
	authHandler.SetUserHandler(userHandler)
	authHandler.SetOtpHandler(OtpHandler)
	authController.SetHandler(authHandler)
	autherRouter.SetController(authController)
	autherRouter.Setup(apiRouteGroup.Group("/auth"))

	return router
}
