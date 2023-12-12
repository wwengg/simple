// @Title
// @Description
// @Author  Wangwengang  2023/12/12 00:45
// @Update  Wangwengang  2023/12/12 00:45
package http

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/wwengg/simple/core/sconfig"
	"time"
)

type GinEngine struct {
	engine *gin.Engine
	config *sconfig.Gateway

	PublicRouterGroup  *gin.RouterGroup
	PrivateRouterGroup *gin.RouterGroup
}

func NewGinEngine(config *sconfig.Config) *GinEngine {
	engine := gin.New()

	if config.Gateway.PublicRouterPrefix == "" {
		config.Gateway.PublicRouterPrefix = "/"
	}
	if config.Gateway.PublicRouterPrefix[0] != '/' {
		config.Gateway.PublicRouterPrefix = "/" + config.Gateway.PublicRouterPrefix
	}

	if config.Gateway.PrivateRouterPrefix == "" {
		config.Gateway.PrivateRouterPrefix = "/"
	}
	if config.Gateway.PrivateRouterPrefix[0] != '/' {
		config.Gateway.PrivateRouterPrefix = "/" + config.Gateway.PrivateRouterPrefix
	}

	return &GinEngine{
		config:             &config.Gateway,
		engine:             engine,
		PublicRouterGroup:  engine.Group(config.Gateway.PublicRouterPrefix),
		PrivateRouterGroup: engine.Group(config.Gateway.PrivateRouterPrefix),
	}
}

func (g *GinEngine) Serve() {
	address := fmt.Sprintf(":%d", g.config.Addr)
	s := endless.NewServer(address, g.engine)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20

	_ = fmt.Errorf(s.ListenAndServe().Error())
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
