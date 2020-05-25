package router

import "github.com/gin-gonic/gin"

// InitRouter init router
func InitRouter() *gin.Engine {
	r := gin.New()
	initSwagger(r)

	return r
}
