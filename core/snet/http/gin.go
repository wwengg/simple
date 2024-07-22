// @Title
// @Description
// @Author  Wangwengang  2023/12/12 00:45
// @Update  Wangwengang  2023/12/12 00:45
package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/slog"
	"net/http"
)

type GinEngine struct {
	engine *gin.Engine
	config *sconfig.Gateway
	ln     *quic.EarlyListener

	PublicRouterGroup  *gin.RouterGroup
	PrivateRouterGroup *gin.RouterGroup
}

func NewGinEngine(config *sconfig.Gateway, ln *quic.EarlyListener) *GinEngine {
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
		ln:                 ln,
		PublicRouterGroup:  engine.Group(config.PublicRouterPrefix),
		PrivateRouterGroup: engine.Group(config.PrivateRouterPrefix),
	}
}

func (g *GinEngine) Serve() {
	s := http3.Server{
		Handler: g.engine,
	}
	// 同时开启Tcp和Udp
	hErr := make(chan error, 1)
	qErr := make(chan error, 1)
	go func() {
		hErr <- http.ListenAndServeTLS(fmt.Sprintf(":%d", g.config.Addr), g.config.CertPath, g.config.KeyPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			s.SetQUICHeaders(w.Header())
			g.engine.ServeHTTP(w, r)
		}))
	}()
	go func() {
		qErr <- s.ServeListener(g.ln)
	}()
	select {
	case err := <-hErr:
		s.Close()
		slog.Ins().Error(err.Error())
	case err := <-qErr:
		// Cannot close the HTTP server or wait for requests to complete properly :/
		slog.Ins().Error(err.Error())
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
