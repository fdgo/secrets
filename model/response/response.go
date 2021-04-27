package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespSuccess(c *gin.Context, data interface{}, format string, a ...interface{}) {
	c.JSON(http.StatusOK, struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{Code: http.StatusOK, Msg: fmt.Sprintf(format, a...), Data: data})
}
func RespFailed(c *gin.Context, code int, data interface{}, format string, a ...interface{}) {
	c.JSON(http.StatusOK, struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}{Code: code, Msg: fmt.Sprintf(format, a...), Data: data})
}
