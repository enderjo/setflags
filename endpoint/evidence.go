package endpoint

import (
	"github.com/PioneerDev/setflags/endpoint/response"
	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/service"
	"github.com/gin-gonic/gin"
)

//EvidenceEndpoint EvidenceEndpoint
type EvidenceEndpoint struct {
}

//List List
func (evidenceEndpoint EvidenceEndpoint) List(c *gin.Context) {
	id := c.Param("id")
	flags := service.NewEvidenceService().List(id)
	response.Success(c, flags)
}

//Save save evidence
func (evidenceEndpoint EvidenceEndpoint) Save(c *gin.Context) {
	var evidence model.Evidence
	err := c.ShouldBindJSON(&evidence)
	if err != nil {
		response.Fail(c, 1, "param error")
		return
	}

	save := service.NewEvidenceService().Save(&evidence)
	if !save {
		response.Fail(c, 1, "save error")
		return
	}

	response.Success(c, evidence)
}
