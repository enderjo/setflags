package router

import (
	"github.com/PioneerDev/setflags/config"
	"github.com/gin-gonic/gin"
)

// InitRouter init router
func InitRouter() *gin.Engine {
	gin.SetMode(config.ServerConfig.Mode)

	router := gin.Default()
	initSwagger(router)

	apiGroup := router.Group("api")
	initEndpoint(apiGroup)

	return router
}
