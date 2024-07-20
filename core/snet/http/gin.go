// @Title
// @Description
// @Author  Wangwengang  2023/12/12 00:45
// @Update  Wangwengang  2023/12/12 00:45
package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/slog"
)

type GinEngine struct {
	engine *gin.Engine
	config *sconfig.Gateway

	PublicRouterGroup  *gin.RouterGroup
	PrivateRouterGroup *gin.RouterGroup
}

func NewGinEngine(config *sconfig.Gateway) *GinEngine {
	engine := gin.New()
	engine.UseH2C = true

	if config.PublicRouterPrefix == "" {
		config.PublicRouterPrefix = "/"
	}
	if config.PublicRouterPrefix[0] != '/' {
		config.PublicRouterPrefix = "/" + config.PublicRouterPrefix
	}

	if config.PrivateRouterPrefix == "" {
		config.PrivateRouterPrefix = "/"
	}
	if config.PrivateRouterPrefix[0] != '/' {
		config.PrivateRouterPrefix = "/" + config.PrivateRouterPrefix
	}

	return &GinEngine{
		config:             config,
		engine:             engine,
		PublicRouterGroup:  engine.Group(config.PublicRouterPrefix),
		PrivateRouterGroup: engine.Group(config.PrivateRouterPrefix),
	}
}

func (g *GinEngine) Serve(certFile string, keyFile string) {
	address := fmt.Sprintf(":%d", g.config.Addr)

	go func() {
		err := http3.ListenAndServe(address, certFile, keyFile, g.engine)
		if err != nil {
			slog.Ins().Error(err.Error())
			return
		}
	}()

	// windows or other
	s := InitServer(address, g.engine)
	slog.Ins().Error(s.ListenAndServe().Error())
}

func (g *GinEngine) AddPublicHandle(route string) {
	{

	}
}

func (g *GinEngine) AddPrivateHandle() {

}

func (g *GinEngine) GetPublicRouterGroup() *gin.RouterGroup {
	return g.PublicRouterGroup
}

func (g *GinEngine) GetPrivateRouterGroup() *gin.RouterGroup {
	return g.PrivateRouterGroup
}
