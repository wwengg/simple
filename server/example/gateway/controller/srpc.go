// @Title
// @Description
// @Author  Wangwengang  2023/12/12 15:16
// @Update  Wangwengang  2023/12/12 15:16
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wwengg/simple/core/slog"
	"github.com/wwengg/simple/server/example/gateway/global"
	"github.com/wwengg/simple/server/example/gateway/model/response"
	"io"
	"strconv"
)

func Http2RpcxPost(c *gin.Context) {
	var isJson, isProtobuf bool
	if a, exist := c.Get("isJson"); exist {
		isJson = a.(bool)
	}
	if a, exist := c.Get("isProtobuf"); exist {
		isProtobuf = a.(bool)
	}
	servicePath := c.Param("servicePath")
	if servicePath == "" {
		global.SLog.Error("empty servicepath")
	}

	serviceMethod := c.Param("serviceMethod")
	if serviceMethod == "" {
		global.SLog.Error("empty servicemethod")
	}

	payload, err := io.ReadAll(c.Request.Body)
	c.Request.Body.Close()
	if err != nil {
		global.SLog.Error(err.Error())
	}
	meta := make(map[string]string, 0)
	var resp []byte
	if isJson {
		meta, resp, err = global.SRPC.RPCJson(servicePath, serviceMethod, payload)
	} else if isProtobuf {
		meta, resp, err = global.SRPC.RPCProtobuf(servicePath, serviceMethod, payload)
	}
	if err != nil {
		global.SLog.Debug("code != 0", slog.String("code", err.Error()))
		if code, err := strconv.Atoi(err.Error()); err != nil {
			response.GatewayResult(int32(code), "error", c)
		} else {
			global.SLog.Errorf("SRPC SendRaw error=%s", err.Error())
			response.GatewayResult(500, "error", c)
		}
		return
	}
	if meta["X-RPCX-MessageStatusType"] == "Error" {
		global.SLog.Errorf("RPCX error = %s", meta["X-RPCX-ErrorMessage"])
		response.GatewayResult(500, meta["X-RPCX-ErrorMessage"], c)
		return
	}

	response.SrpcResult(resp, c)
}
