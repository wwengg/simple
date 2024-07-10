// @Title
// @Description
// @Author  Wangwengang  2023/12/12 00:40
// @Update  Wangwengang  2023/12/12 00:40
package http

import "net"

type HttpServer interface {
	Serve()
}

type server interface {
	ListenAndServe() error
	Serve(l net.Listener) error
}
