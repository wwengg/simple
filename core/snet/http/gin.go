// @Title
// @Description
// @Author  Wangwengang  2023/12/12 00:45
// @Update  Wangwengang  2023/12/12 00:45
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/slog"
)

type GinEngine struct {
	engine *gin.Engine
	config *sconfig.Gateway
	ln     http3.QUICEarlyListener

	PublicRouterGroup  *gin.RouterGroup
	PrivateRouterGroup *gin.RouterGroup
}

func NewGinEngine(config *sconfig.Gateway, ln http3.QUICEarlyListener) *GinEngine {
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

func (g *GinEngine) Serve() {
	s := http3.Server{
		Handler: g.engine,
	}
	slog.Ins().Error(s.ServeListener(g.ln).Error())
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
