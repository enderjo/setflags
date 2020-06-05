package endpoint

import (
	"github.com/PioneerDev/setflags/endpoint/response"
	"github.com/PioneerDev/setflags/service"
	"github.com/gin-gonic/gin"
)

//AssetEndpoint AssetEndpoint
type AssetEndpoint struct {
}

//Get get the asset
func (assetEndpoint AssetEndpoint) Get(c *gin.Context) {
	id := c.Param("id")
	asset := service.NewAssetService().Get(id)

	response.Success(c, asset)
}
