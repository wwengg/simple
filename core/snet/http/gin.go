// @Title
// @Description
// @Author  Wangwengang  2023/12/12 00:45
// @Update  Wangwengang  2023/12/12 00:45
package http

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/wwengg/simple/core/sconfig"
)

type GinEngine struct {
	engine *gin.Engine
	config *sconfig.Gateway

	PublicRouterGroup  *gin.RouterGroup
	PrivateRouterGroup *gin.RouterGroup
}

func NewGinEngine(config *sconfig.Gateway) *GinEngine {
	engine := gin.New()

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
	address := g.config.Addr

	// windows or other
	s := InitServer(address, g.engine)
	ln, err := net.Listen("tcp4", address)
	if err != nil {
		panic(err)
	}
	type tcpKeepAliveListener struct {
		*net.TCPListener
	}
	erred := s.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
	if erred != nil {
		panic(erred)
	}
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
