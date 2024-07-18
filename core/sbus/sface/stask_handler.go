package sface

type STaskHandler interface {
	AddRouter(msgID int32, router SRouter)
	StartWorkerPool()               //  Start the worker pool
	SendTaskToTaskQueue(task STask) // Pass the message to the TaskQueue for processing by the worker(将消息交给TaskQueue,由worker进行处理)
	Stop()
}
