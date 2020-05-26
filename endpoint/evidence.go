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

//List List
func (evidenceEndpoint EvidenceEndpoint) List(c *gin.Context) {
	var flag model.Flag
	database.ServerDb.First(&flag, "21")
	c.JSON(http.StatusOK, flag)
}

//Save save evidence
func (evidenceEndpoint EvidenceEndpoint) Save(c *gin.Context) {

}
