package sbus

import (
	"encoding/hex"
	"github.com/wwengg/simple/core/sbus/sface"
	"github.com/wwengg/simple/core/slog"
	"sync"
)

type MsgHandler struct {
	Apis map[uint32]sface.SRouter
	// The number of worker goroutines in the business work Worker pool
	// (业务工作Worker池的数量)
	WorkerPoolSize uint32

	// A collection of idle workers, used for zconf.WorkerModeBind
	// 空闲worker集合，用于zconf.WorkerModeBind
	freeWorkers  map[uint32]struct{}
	freeWorkerMu sync.Mutex

	// A message queue for workers to take tasks
	// (Worker负责取任务的消息队列)
	TaskQueue []chan sface.STask
}

func newMsgHandle(workPoolSize uint32) *MsgHandler {
	var freeWorkers map[uint32]struct{}

	handler := &MsgHandler{
		Apis:           make(map[uint32]sface.SRouter),
		WorkerPoolSize: workPoolSize,
		// One worker corresponds to one queue (一个worker对应一个queue)
		TaskQueue:   make([]chan sface.STask, workPoolSize),
		freeWorkers: freeWorkers,
	}
	return handler
}

// SendMsgToTaskQueue sends the message to the TaskQueue for processing by the worker
// (将消息交给TaskQueue,由worker进行处理)
func (mh *MsgHandler) SendMsgToTaskQueue(task sface.STask) {
	workerID := task.GetWorkerID()

	mh.TaskQueue[workerID] <- task
	slog.Ins().Debugf("SendMsgToTaskQueue-->%s", hex.EncodeToString(task.GetData()))
}
