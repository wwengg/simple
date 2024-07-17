package sface

type STaskHandler interface {
	AddRouter(msgID uint32, router SRouter)
	StartWorkerPool(maxWorkerTaskLen uint32) //  Start the worker pool
	SendMsgToTaskQueue(request SRouter)      // Pass the message to the TaskQueue for processing by the worker(将消息交给TaskQueue,由worker进行处理)
}
