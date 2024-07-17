package sbus

import (
	"encoding/hex"
	"fmt"
	"github.com/wwengg/simple/core/sbus/sface"
	"github.com/wwengg/simple/core/slog"
)

type TaskHandler struct {
	Apis map[int32]sface.SRouter
	// The number of worker goroutines in the business work Worker pool
	// (业务工作Worker池的数量)
	WorkerPoolSize uint32
	// A message queue for workers to take tasks
	// (Worker负责取任务的消息队列)
	TaskQueue chan sface.STask
}

func newTaskHandler(workPoolSize, maxTaskQueueLen uint32) *TaskHandler {
	var freeWorkers map[uint32]struct{}
	freeWorkers = make(map[uint32]struct{}, workPoolSize)
	for i := uint32(0); i < workPoolSize; i++ {
		freeWorkers[i] = struct{}{}
	}
	handler := &TaskHandler{
		Apis:           make(map[int32]sface.SRouter),
		WorkerPoolSize: workPoolSize,
		TaskQueue:      make(chan sface.STask, maxTaskQueueLen),
	}
	return handler
}

func (mh *TaskHandler) AddRouter(msgID int32, router sface.SRouter) {
	// 1. Check whether the current API processing method bound to the msgID already exists
	// (判断当前msg绑定的API处理方法是否已经存在)
	if _, ok := mh.Apis[msgID]; ok {
		msgErr := fmt.Sprintf("repeated api , msgID = %+v\n", msgID)
		panic(msgErr)
	}
	// 2. Add the binding relationship between msg and API
	// (添加msg与api的绑定关系)
	mh.Apis[msgID] = router
	slog.Ins().Infof("Add Router msgID = %d", msgID)
}

// SendMsgToTaskQueue sends the message to the TaskQueue for processing by the worker
// (将消息交给TaskQueue,由worker进行处理)
func (mh *TaskHandler) SendMsgToTaskQueue(task sface.STask) {

	mh.TaskQueue <- task
	slog.Ins().Debugf("SendMsgToTaskQueue-->%s", hex.EncodeToString(task.GetData()))
}

// doFuncHandler handles functional requests (执行函数式请求)
func (mh *TaskHandler) doFuncHandler(task sface.SFuncTask, workerID int) {
	defer func() {
		if err := recover(); err != nil {
			slog.Ins().Errorf("workerID: %d doFuncRequest panic: %v", workerID, err)
		}
	}()
	// Execute the functional request (执行函数式请求)
	task.CallFunc()
}

// doMsgHandler immediately handles messages in a non-blocking manner
// (立即以非阻塞方式处理消息)
func (mh *TaskHandler) doMsgHandler(task sface.STask, workerID int) {
	defer func() {
		if err := recover(); err != nil {
			slog.Ins().Errorf("workerID: %d doMsgHandler panic: %v", workerID, err)
		}
	}()

	msgId := task.GetMsgID()
	handler, ok := mh.Apis[msgId]

	if !ok {
		slog.Ins().Errorf("api msgID = %d is not FOUND!", task.GetMsgID())
		return
	}

	// Bind the Request request to the corresponding Router relationship
	// (Request请求绑定Router对应关系)
	task.BindRouter(handler)

	// Execute the corresponding processing method
	task.Call()

	// 执行完成后回收 Request 对象回对象池
	PutTask(task)
}

// StartOneWorker starts a worker workflow
// (启动一个Worker工作流程)
func (mh *TaskHandler) StartOneWorker(workerID int, taskQueue chan sface.STask) {
	slog.Ins().Debugf("Worker ID = %d is started.", workerID)
	// Continuously wait for messages in the queue
	// (不断地等待队列中的消息)
	for {
		select {
		// If there is a message, take out the Request from the queue and execute the bound business method
		// (有消息则取出队列的Request，并执行绑定的业务方法)
		case task := <-taskQueue:

			switch req := task.(type) {

			case sface.SFuncTask:
				// Internal function call request (内部函数调用request)

				mh.doFuncHandler(req, workerID)

			case sface.STask: // Client message request
				mh.doMsgHandler(req, workerID)
			}
		}
	}
}

// StartWorkerPool starts the worker pool
func (mh *TaskHandler) StartWorkerPool() {
	// Iterate through the required number of workers and start them one by one
	// (遍历需要启动worker的数量，依此启动)
	for i := 0; i < int(mh.WorkerPoolSize); i++ {

		// Start the current worker, blocking and waiting for messages to be passed in the corresponding task queue
		// (启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来)
		go mh.StartOneWorker(i, mh.TaskQueue)
	}
}
