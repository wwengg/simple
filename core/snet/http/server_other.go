//go:build !windows
// +build !windows

// @Title
// @Description
// @Author  Wangwengang  2023/12/16 14:18
// @Update  Wangwengang  2023/12/16 14:18
package http

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func InitServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
