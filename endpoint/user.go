package endpoint

import (
	"net/http"

	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/model/database"
	"github.com/gin-gonic/gin"
)

//UserEndpoint UserEndpoint
type UserEndpoint struct {
}

//FindAll findAll
func (userEndpoint UserEndpoint) FindAll(c *gin.Context) {
	var flag model.Flag
	database.ServerDb.First(&flag, "21")
	c.JSON(http.StatusOK, flag)
}
