package router

import (
	"github.com/PioneerDev/setflags/config"
	"github.com/PioneerDev/setflags/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter init router
func InitRouter() *gin.Engine {
	gin.SetMode(config.ServerConfig.Mode)

	router := gin.Default()

	router.Use(middleware.Cors())

	initSwagger(router)

	apiGroup := router.Group("api")
	initEndpoint(apiGroup)

	return router
}
