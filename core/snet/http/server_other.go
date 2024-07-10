//go:build !windows
// +build !windows

// @Title
// @Description
// @Author  Wangwengang  2023/12/16 14:18
// @Update  Wangwengang  2023/12/16 14:18
package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer(address string, router *gin.Engine) server {
	//s := endless.NewServer(address, router)
	//s.ReadHeaderTimeout = 20 * time.Second
	//s.WriteTimeout = 20 * time.Second
	//s.MaxHeaderBytes = 1 << 20
	//return s
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
