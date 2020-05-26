package endpoint

import (
	"github.com/PioneerDev/setflags/endpoint/response"
	"github.com/PioneerDev/setflags/model"
	"github.com/PioneerDev/setflags/model/database"
	"github.com/gin-gonic/gin"
)

//FlagEndpoint FlagEndpoint
type FlagEndpoint struct {
}

// List get all flaglist
// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.Flag true "用户注册接口"
// @Success 200 {string} string "{"data":{}}"
// @Router /flags [get]
func (flagEndpoint FlagEndpoint) List(c *gin.Context) {
	var flags []model.Flag
	database.ServerDb.Find(&flags).Order("createdAt desc")

	response.Success(c, flags)
}

//Save create a flag
func (flagEndpoint FlagEndpoint) Save(c *gin.Context) {
	response.Fail(c, 1, "error")
}

//Update update flag status
func (flagEndpoint FlagEndpoint) Update(c *gin.Context) {

}

//MyflagList list my flags
func (flagEndpoint FlagEndpoint) MyflagList(c *gin.Context) {

}
