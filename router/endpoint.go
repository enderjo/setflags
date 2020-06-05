package router

import (
	"github.com/PioneerDev/setflags/endpoint"
	"github.com/gin-gonic/gin"
)

func initEndpoint(router *gin.RouterGroup) {
	//flags
	router.GET("/flags", endpoint.FlagEndpoint{}.List)
	router.POST("/flags/:id", endpoint.FlagEndpoint{}.Save)
	router.PUT("/flags/:id/:op", endpoint.FlagEndpoint{}.Update)
	router.GET("/myflags/:id", endpoint.FlagEndpoint{}.MyflagList)

	//evidences
	router.POST("/evidences/:id", endpoint.EvidenceEndpoint{}.Save)
	router.GET("/flags/:flag_id/evidences", endpoint.EvidenceEndpoint{}.List)

	//users
	router.GET("/users/:user_id/rewards/:flag_id", endpoint.UserEndpoint{}.List)

	//assets
	router.GET("/assets/:id", endpoint.AssetEndpoint{}.Get)
}
