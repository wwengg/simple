// @Title
// @Description
// @Author  Wangwengang  2023/12/12 18:42
// @Update  Wangwengang  2023/12/12 18:42
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func SrpcResult(data []byte, c *gin.Context) {
	var isJson, isProtobuf bool
	if a, exist := c.Get("isJson"); exist {
		isJson = a.(bool)
	}
	if a, exist := c.Get("isProtobuf"); exist {
		isProtobuf = a.(bool)
	}
	if isJson {
		c.Data(http.StatusOK, "application/json; charset=utf-8", data)
	} else if isProtobuf {
		c.Data(http.StatusOK, "application/x-protobuf", data)
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}
