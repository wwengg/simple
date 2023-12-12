// @Title
// @Description
// @Author  Wangwengang  2023/12/12 15:15
// @Update  Wangwengang  2023/12/12 15:15
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wwengg/simple/server/example/gateway/controller"
)

func InitSRPCRouter(Router *gin.RouterGroup) {
	Http2SRPCRouter := Router
	{
		Http2SRPCRouter.POST(":servicePath/:serviceMethod", controller.Http2RpcxPost)
	}
}
