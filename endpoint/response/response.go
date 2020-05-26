package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Response response
type Response struct {
	Error *Error      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

//Error response error
type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

//Success success
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Data: data})
}

//Fail fail
func Fail(c *gin.Context, code int, desc string) {
	c.JSON(http.StatusOK, Response{Error: &Error{code, desc}})
}
