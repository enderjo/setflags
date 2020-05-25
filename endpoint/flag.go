package endpoint

import (
	"net/http"

	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/model/database"
	"github.com/gin-gonic/gin"
)

//FlagEndpoint FlagEndpoint
type FlagEndpoint struct {
}

//FindAll findAll
func (flagEndpoint FlagEndpoint) FindAll(c *gin.Context) {
	var flag model.Flag
	database.ServerDb.First(&flag, "21")
	c.JSON(http.StatusOK, flag)
}
