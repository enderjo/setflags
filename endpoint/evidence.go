package endpoint

import (
	"net/http"

	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/model/database"
	"github.com/gin-gonic/gin"
)

//EvidenceEndpoint EvidenceEndpoint
type EvidenceEndpoint struct {
}

//FindAll findAll
func (evidenceEndpoint EvidenceEndpoint) FindAll(c *gin.Context) {
	var flag model.Flag
	database.ServerDb.First(&flag, "21")
	c.JSON(http.StatusOK, flag)
}
