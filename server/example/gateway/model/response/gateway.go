// @Title
// @Description
// @Author  Wangwengang  2023/12/12 19:05
// @Update  Wangwengang  2023/12/12 19:05
package response

import (
	"github.com/gin-gonic/gin"
	"github.com/wwengg/simple/proto/pbgateway"
	"net/http"
)

func GatewayResult(code int32, msg string, c *gin.Context) {
	var isJson, isProtobuf bool
	if a, exist := c.Get("isJson"); exist {
		isJson = a.(bool)
	}
	if a, exist := c.Get("isProtobuf"); exist {
		isProtobuf = a.(bool)
	}
	if isJson {
		c.JSON(http.StatusOK, Response{
			code,
			msg,
		})
	} else if isProtobuf {
		reply := &pbgateway.Response{
			Code: code,
			Msg:  msg,
		}
		c.ProtoBuf(http.StatusOK, reply)

	} else {

	}

}

func Ok(c *gin.Context) {
	GatewayResult(200, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	GatewayResult(200, message, c)
}

func Fail(code int32, c *gin.Context) {
	GatewayResult(code, "error", c)
}
