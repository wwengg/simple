package sbus

import (
	"context"
	"errors"
	"github.com/wwengg/simple/core/slog"
	"net"
	"sync"
	"syscall"
	"time"
)

type MsgData struct {
	MsgID uint32
	data  []byte
}

var MsgDataPool = new(sync.Pool)

func init() {
	MsgDataPool.New = func() interface{} {
		return allocateMsgData()
	}
}

func allocateMsgData() *MsgData {
	msgData := new(MsgData)
	return msgData
}
func (md *MsgData) Reset(msgId uint32, data []byte) {
	md.MsgID = msgId
	md.data = data
}

func GetMsgData(msgId uint32, data []byte) *MsgData {

	// 根据当前模式判断是否使用对象池

	// 从对象池中取得一个 Request 对象,如果池子中没有可用的 Request 对象则会调用 allocateRequest 函数构造一个新的对象分配
	r := MsgDataPool.Get().(*MsgData)
	// 因为取出的 Request 对象可能是已存在也可能是新构造的,无论是哪种情况都应该初始化再返回使用
	r.Reset(msgId, data)
	return r
}

func PutMsgData(msgData *MsgData) {
	MsgDataPool.Put(msgData)
}

type BaseConnection struct {
	// The ID of the current connection, also known as SessionID, globally unique, used by server Connection
	// uint64 range: 0~18,446,744,073,709,551,615
	// This is the maximum number of connID theoretically supported by the process
	// (当前连接的ID 也可以称作为SessionID，ID全局唯一 ，服务端Connection使用
	// uint64 取值范围：0 ~ 18,446,744,073,709,551,615
	// 这个是理论支持的进程connID的最大数量)
	ConnID uint64
	// connection id for string
	// (字符串的连接id)
	ConnIdStr string
	// The message management module that manages MsgID and the corresponding processing method
	// (消息管理MsgID和对应处理方法的消息管理模块)
	TaskHandler STaskHandler
	// onConnStart is the Hook function when the current connection is created.
	// (当前连接创建时Hook函数)
	OnConnStart func(conn SConnection)
	// onConnStop is the Hook function when the current connection is created.
	// (当前连接断开时的Hook函数)
	OnConnStop func(conn SConnection)
	// ctx and cancel are used to notify that the connection has exited/stopped.
	// (告知该链接已经退出/停止的channel)
	ctx    context.Context
	cancel context.CancelFunc

	// frameDecoder is the decoder for splitting or splicing data packets.
	// (断粘包解码器)
	FrameDecoder SFrameDecoder

	Datapack SDataPack

	SendFunc func([]byte) error
	ReadFunc func(conn SConnection, buffer []byte) (n int, err error)

	// msgLock is used for locking when users send and receive messages.
	// (用户收发消息的Lock)
	msgLock sync.RWMutex

	// property is the connection attribute. (链接属性)
	Property map[string]string

	// propertyLock protects the current property lock. (保护当前property的锁)
	propertyLock sync.Mutex

	IOReadBuffSize uint32
}

func (bc *BaseConnection) callOnConnStart() {
	if bc.OnConnStart != nil {
		slog.Ins().Infof("CallOnConnStart....")
		bc.OnConnStart(bc)
	}
}

func (bc *BaseConnection) callOnConnStop() {
	if bc.OnConnStop != nil {
		slog.Ins().Infof("CallOnConnStart....")
		bc.OnConnStop(bc)
	}
}

func (bc *BaseConnection) isClosed() bool {
	return bc.ctx == nil || bc.ctx.Err() != nil
}

func (bc *BaseConnection) StartReader() {
	slog.Ins().Infof("[Reader Goroutine is running]")
	defer slog.Ins().Infof("%s [conn Reader exit!]", bc.ConnIdStr)
	defer bc.Stop()
	defer func() {
		if err := recover(); err != nil {
			slog.Ins().Errorf("Reader connID=%d, panic err=%v", bc.GetConnID(), err)
		}
	}()
	//Reduce buffer allocation times to improve efficiency
	// add by ray 2023-02-03
	buffer := make([]byte, bc.IOReadBuffSize)

	for {
		select {
		case <-bc.ctx.Done():
			// 停止循环 不读了，连接断开啦！！
			return
		default:
			if n, err := bc.ReadFunc(bc, buffer); err != nil {
				slog.Ins().Error(err.Error())
				return
			} else {
				if n == 0 {
					continue
				}
				//if n > 0 && c.hc != nil {
				//	c.updateActivity()
				//}
				// Deal with the custom protocol fragmentation problem, added by uuxia 2023-03-21
				// (处理自定义协议断粘包问题)
				if bc.FrameDecoder != nil {
					// Decode the 0-n bytes of data read
					// (为读取到的0-n个字节的数据进行解码)
					bufArrays, err2 := bc.FrameDecoder.Decode(buffer[0:n])
					if bufArrays == nil {
						continue
					}
					for _, bytes := range bufArrays {
						msg, err := bc.Datapack.Unpack(bytes)
						if err != nil {
							slog.Ins().Error(err.Error())
							continue
						}
						// Get the current client's Request data
						// (得到当前客户端请求的Request数据)
						task := GetTask(bc, msg)
						bc.TaskHandler.SendTaskToTaskQueue(task)
					}
					if err2 != nil {
						slog.Ins().Error(err2.Error())
						return // 发送过长数据包或协议错误，断开连接
					}
				} else {
					msg, err := bc.Datapack.Unpack(buffer[0:n])
					if err != nil {
						slog.Ins().Error(err.Error())
						continue
					}
					// Get the current client's Request data
					// (得到当前客户端请求的Request数据)
					task := GetTask(bc, msg)
					bc.TaskHandler.SendTaskToTaskQueue(task)
				}
			}

		}
	}
}

// Start()
func (bc *BaseConnection) Start() {
	bc.ctx, bc.cancel = context.WithCancel(context.Background())

	bc.callOnConnStart()

	// Start the Goroutine for reading data from the client
	// (开启用户从客户端读取数据流程的Goroutine)
	go bc.StartReader()

	select {
	case <-bc.ctx.Done():
		// If the user has registered a close callback for the connection, it should be called explicitly at this moment.
		// (如果用户注册了该链接的	关闭回调业务，那么在此刻应该显示调用)
		bc.callOnConnStop()
		return
	}
}
func (bc *BaseConnection) Stop() {
	bc.cancel()
}
func (bc *BaseConnection) Context() context.Context {
	return bc.ctx
}
func (bc *BaseConnection) GetConnID() uint64 {
	return bc.ConnID
}
func (bc *BaseConnection) GetConnIdStr() string {
	return bc.ConnIdStr
}
func (bc *BaseConnection) GetTaskHandler() STaskHandler {
	return bc.TaskHandler
}
func (bc *BaseConnection) RemoteAddr() net.Addr     { return nil }
func (bc *BaseConnection) LocalAddr() net.Addr      { return nil }
func (bc *BaseConnection) LocalAddrString() string  { return "" }
func (bc *BaseConnection) RemoteAddrString() string { return "" }
func (bc *BaseConnection) SendData(data []byte) error {
	bc.msgLock.RLock()
	defer bc.msgLock.RUnlock()
	defer func() {
		if err := recover(); err != nil {
			slog.Ins().Errorf("SendData connID=%d, panic err=%v", bc.GetConnID(), err)
		}
	}()
	if bc.isClosed() == true {
		return errors.New("Connection closed when send Data")
	}
	err := bc.SendFunc(data)
	if err != nil {
		slog.Ins().Errorf("SendMsg err data = %+v, err = %+v", data, err)
		return err
	}
	return nil
}
func (bc *BaseConnection) SendMsg(msgID uint32, data []byte) error { return nil }
func (bc *BaseConnection) SetProperty(key string, value string) {
	bc.propertyLock.Lock()
	defer bc.propertyLock.Unlock()
	if bc.Property == nil {
		bc.Property = make(map[string]string)
	}

	bc.Property[key] = value
}
func (bc *BaseConnection) GetProperty(key string) (string, error) {
	bc.propertyLock.Lock()
	defer bc.propertyLock.Unlock()

	if value, ok := bc.Property[key]; ok {
		return value, nil
	}

	return "", errors.New("no property found")
}
func (bc *BaseConnection) RemoveProperty(key string) {
	bc.propertyLock.Lock()
	defer bc.propertyLock.Unlock()

	delete(bc.Property, key)
}
func (bc *BaseConnection) IsAlive() bool                          { return true }
func (bc *BaseConnection) SetHeartBeat(checker SHeartbeatChecker) {}

// func (bc *BaseConnection) AddCloseCallback(handler, key interface{}, callback func()) {}
// func (bc *BaseConnection) RemoveCloseCallback(handler, key interface{})               {}
// func (bc *BaseConnection) InvokeCloseCallbacks()
//
// Read implements the Conn Read method.
func (bc *BaseConnection) Read(b []byte) (int, error) {
	return 0, nil
}

// Write implements the Conn Write method.
func (bc *BaseConnection) Write(b []byte) (int, error) {
	return 0, nil
}

// Close closes the connection.
func (bc *BaseConnection) Close() error { return nil }

// SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
func (bc *BaseConnection) SetDeadline(t time.Time) error { return nil }

// SetReadDeadline implements the Conn SetReadDeadline method.
func (bc *BaseConnection) SetReadDeadline(t time.Time) error { return nil }

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (bc *BaseConnection) SetWriteDeadline(t time.Time) error { return nil }

// SetReadBuffer sets the size of the operating system's receive buffer associated with the connection.
func (bc *BaseConnection) SetReadBuffer(bytes int) error { return nil }

// SetWriteBuffer sets the size of the operating system's transmit buffer associated with the connection.
func (bc *BaseConnection) SetWriteBuffer(bytes int) error { return nil }

// SyscallConn returns a raw network connection. This implements the syscall.Conn interface.
func (bc *BaseConnection) SyscallConn() (syscall.RawConn, error) { return nil, nil }
