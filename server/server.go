// @Title
// @Description
// @Author  Wangwengang  2023/12/10 11:51
// @Update  Wangwengang  2023/12/10 11:51
package server

type Server interface {
	Start()   // 启动服务，读取之前运行的数据
	Stop()    // 停止服务，保存数据
	Restart() // 重启服务

	AddHandler()
	AddHandlerSlices()

	// Add Http authentication method
	SetHttpAuth()
}
