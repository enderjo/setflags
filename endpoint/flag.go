package endpoint

import (
	"github.com/PioneerDev/setflags/endpoint/response"
	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/service"
	"github.com/gin-gonic/gin"
)

//FlagEndpoint FlagEndpoint
type FlagEndpoint struct {
}

// List get all flaglist
// @Tags Flags
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.Flag true "用户注册接口"
// @Success 200 {string} string "{"data":{}}"
// @Router /flags [get]
func (flagEndpoint FlagEndpoint) List(c *gin.Context) {
	flags := service.NewFlagService().List()
	response.Success(c, flags)
}

//Save create a flag
func (flagEndpoint FlagEndpoint) Save(c *gin.Context) {
	var flag model.Flag
	err := c.ShouldBindJSON(&flag)
	if err != nil {
		response.Fail(c, 1, "param error")
		return
	}

	save := service.NewFlagService().Save(&flag)
	if !save {
		response.Fail(c, 1, "save error")
		return
	}

	response.Success(c, flag)
}

//Update update flag status
func (flagEndpoint FlagEndpoint) Update(c *gin.Context) {
	var flag model.Flag
	err := c.ShouldBindJSON(&flag)
	if err != nil {
		response.Fail(c, 1, "param error")
		return
	}

	save := service.NewFlagService().Update(&flag)
	if !save {
		response.Fail(c, 1, "update error")
		return
	}

	response.Success(c, flag)
}

//MyflagList list my flags
func (flagEndpoint FlagEndpoint) MyflagList(c *gin.Context) {
	id := c.Param("id")
	flags := service.NewFlagService().MyList(id)
	response.Success(c, flags)
}
