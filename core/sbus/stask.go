package sbus

type HandleStep int

type SFuncTask interface {
	CallFunc()
}

type STask interface {
	GetConnection() SConnection // Get the connection information of the request(获取请求连接信息)

	GetData() []byte // Get the data of the request message(获取请求消息的数据)
	GetMsgID() int32 // Get the message ID of the request(获取请求的消息ID)
	GetCmd() uint16
	GetMessage() SMsg // Get the raw data of the request message (获取请求消息的原始数据 add by uuxia 2023-03-10)

	BindRouter(router SRouter) // Bind which router handles this request(绑定这次请求由哪个路由处理)
	// Move on to the next handler to start execution, but the function that calls this method will execute in reverse order of their order
	// (转进到下一个处理器开始执行 但是调用此方法的函数会根据先后顺序逆序执行)
	Call()

	//erminate the execution of the processing function, but the function that calls this method will be executed until completion
	// 终止处理函数的运行 但调用此方法的函数会执行完毕
	Abort()

	//Set 在 Request 中存放一个上下文
	Set(key string, value interface{})
	//Get 从 Request 中获取一个上下文信息
	Get(key string) (value interface{}, exists bool)
}

type BaseRequest struct{}

func (br *BaseRequest) GetConnection() SConnection { return nil }
func (br *BaseRequest) GetData() []byte            { return nil }
func (br *BaseRequest) GetMsgID() int32            { return 0 }
func (br *BaseRequest) GetCmd() uint16             { return 0 }
func (br *BaseRequest) GetMessage() SMsg           { return nil }
func (br *BaseRequest) BindRouter(router SRouter)  {}
func (br *BaseRequest) Call()                      {}
func (br *BaseRequest) Abort()                     {}

func (br *BaseRequest) Set(key string, value interface{}) {}

func (br *BaseRequest) Get(key string) (value interface{}, exists bool) { return nil, false }
