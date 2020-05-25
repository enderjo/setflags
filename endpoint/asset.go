package endpoint

import (
	"net/http"

	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/model/database"
	"github.com/gin-gonic/gin"
)

//AssetEndpoint AssetEndpoint
type AssetEndpoint struct {
}

//FindAll findAll
func (assetEndpoint AssetEndpoint) FindAll(c *gin.Context) {
	var flag model.Flag
	database.ServerDb.First(&flag, "21")
	c.JSON(http.StatusOK, flag)
}
