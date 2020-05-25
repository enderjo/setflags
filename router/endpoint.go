package router

import (
	"github.com/PioneerDev/setflags/endpoint"
	"github.com/gin-gonic/gin"
)

func initEndpoint(router *gin.RouterGroup) {
	//flags
	router.GET("/flags", endpoint.FlagEndpoint{}.FindAll)
	router.POST("/flags/:id", endpoint.FlagEndpoint{}.FindAll)
	router.PUT("/flags/:id/:op", endpoint.FlagEndpoint{}.FindAll)
	router.GET("/myflags/:id", endpoint.FlagEndpoint{}.FindAll)

	//evidences
	router.POST("/flags/:flag_id/evidences/:id/:op", endpoint.EvidenceEndpoint{}.FindAll)
	router.GET("/flags/:flag_id/evidences", endpoint.EvidenceEndpoint{}.FindAll)

	//users
	router.GET("/users/:user_id/rewards/:flag_id", endpoint.UserEndpoint{}.FindAll)

	//assets
	router.GET("/assets/:id", endpoint.AssetEndpoint{}.FindAll)
}
