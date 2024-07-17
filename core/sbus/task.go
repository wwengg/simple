package sbus

import (
	"github.com/wwengg/simple/core/sbus/sface"
	"github.com/wwengg/simple/core/snet"
	"sync"
)

const (
	PRE_HANDLE  sface.HandleStep = iota // PreHandle for pre-processing
	HANDLE                              // Handle for processing
	POST_HANDLE                         // PostHandle for post-processing

	HANDLE_OVER
)

var TaskPool = new(sync.Pool)

func init() {
	TaskPool.New = func() interface{} {
		return allocateTask()
	}
}

func allocateTask() sface.STask {
	task := new(Task)
	task.steps = PRE_HANDLE
	task.needNext = true
	task.index = -1
	return task
}

type Task struct {
	sface.BaseRequest
	conn     snet.SConnection
	msg      sface.SMsg
	router   sface.SRouter          // the router that handles this request(请求处理的函数)
	steps    sface.HandleStep       // used to control the execution of router functions(用来控制路由函数执行)
	stepLock sync.RWMutex           // concurrency lock(并发互斥)
	needNext bool                   // whether to execute the next router function(是否需要执行下一个路由函数)
	index    int8                   // router function slice index(路由函数切片索引)
	keys     map[string]interface{} // keys 路由处理时可能会存取的上下文信息
}

func (r *Task) Reset(conn snet.SConnection, msg sface.SMsg) {
	r.steps = PRE_HANDLE
	r.conn = conn
	r.msg = msg
	r.needNext = true
	r.index = -1
	r.keys = nil
}

func GetTask(conn snet.SConnection, msg sface.SMsg) sface.STask {

	// 根据当前模式判断是否使用对象池

	// 从对象池中取得一个 Request 对象,如果池子中没有可用的 Request 对象则会调用 allocateRequest 函数构造一个新的对象分配
	r := TaskPool.Get().(*Task)
	// 因为取出的 Request 对象可能是已存在也可能是新构造的,无论是哪种情况都应该初始化再返回使用
	r.Reset(conn, msg)
	return r
}

func PutTask(task sface.STask) {
	TaskPool.Put(task)
}

func (r *Task) GetMessage() sface.SMsg {
	return r.msg
}

func (r *Task) GetConnection() snet.SConnection {
	return r.conn
}

func (r *Task) GetData() []byte {
	return r.msg.GetData()
}

func (r *Task) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}

func (r *Task) BindRouter(router sface.SRouter) {
	r.router = router
}

func (r *Task) next() {
	if r.needNext == false {
		r.needNext = true
		return
	}

	r.stepLock.Lock()
	r.steps++
	r.stepLock.Unlock()
}

func (r *Task) Call() {

	if r.router == nil {
		return
	}

	for r.steps < HANDLE_OVER {
		switch r.steps {
		case PRE_HANDLE:
			r.router.PreHandle(r)
		case HANDLE:
			r.router.Handle(r)
		case POST_HANDLE:
			r.router.PostHandle(r)
		}

		r.next()
	}

	r.steps = PRE_HANDLE
}
