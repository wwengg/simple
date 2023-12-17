//go:build windows
// +build windows

// @Title
// @Description
// @Author  Wangwengang  2023/12/16 14:19
// @Update  Wangwengang  2023/12/16 14:19
package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
