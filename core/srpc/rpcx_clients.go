// @Title
// @Description
// @Author  Wangwengang  2023/12/11 23:34
// @Update  Wangwengang  2023/12/11 23:34
package srpc

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/wwengg/simple/core/sconfig"
)

type OptionSRPCClients func(s *RPCXClients)

// 自由设置FailMode, 不设置将使用默认方法
func WithFailMode(failMode client.FailMode) OptionSRPCClients {
	return func(s *RPCXClients) {
		s.setFailMode(failMode)
	}
}

// 自由设置SelectMode, 不设置将使用默认方法
func WithSelectMode(selectMode client.SelectMode) OptionSRPCClients {
	return func(s *RPCXClients) {
		s.setSelectMode(selectMode)
	}
}

// 自由设置Option, 不设置将使用默认方法
func WithOption(option client.Option) OptionSRPCClients {
	return func(s *RPCXClients) {
		s.setOption(option)
	}
}

type RPCXClients struct {
	config           *sconfig.RPC
	serviceDiscovery client.ServiceDiscovery
	FailMode         client.FailMode
	SelectMode       client.SelectMode
	Option           client.Option

	mu       sync.RWMutex
	xclients map[string]client.XClient

	seq uint64
}

func TODO() {
	return
}

func NewSRPCClients(config *sconfig.RPC, opts ...OptionSRPCClients) *RPCXClients {
	register, err := CreateServiceDiscovery(config.RegisterAddr, config.RegisterType, config.BasePath)
	if err != nil {
		panic(err)
	}

	rpcxClients := &RPCXClients{
		config:           config,
		serviceDiscovery: register,
		FailMode:         client.Failover,
		SelectMode:       client.RoundRobin,
		Option:           client.Option{},
		xclients:         make(map[string]client.XClient, 0),
	}

	for _, opt := range opts {
		opt(rpcxClients)
	}

	return rpcxClients

}

func (s *RPCXClients) GetReq(servicePath string, serviceMethod string) *protocol.Message {
	req := protocol.NewMessage()
	req.SetMessageType(protocol.Request)
	// servivePath
	req.ServicePath = servicePath

	// serviceMethod
	req.ServiceMethod = serviceMethod

	seq := atomic.AddUint64(&s.seq, 1)
	req.SetSeq(seq)

	return req
}

func (s *RPCXClients) RPC(ctx context.Context, servicePath string, serviceMethod string, payload []byte, serializeType protocol.SerializeType, oneway bool) (meta map[string]string, resp []byte, err error) {
	req := protocol.NewMessage()
	req.SetMessageType(protocol.Request)

	// protobuf协议
	req.SetSerializeType(serializeType)

	// servivePath
	req.ServicePath = servicePath

	// serviceMethod
	req.ServiceMethod = serviceMethod

	req.Payload = payload

	req.SetOneway(oneway) // 不用等服务的响应结果，只管发

	seq := atomic.AddUint64(&s.seq, 1)
	req.SetSeq(seq)

	var xc client.XClient
	xc, err = s.GetXClient(servicePath)
	if err != nil {
		return nil, nil, err
	}

	return xc.SendRaw(ctx, req)
}

func (s *RPCXClients) RPCProtobuf(ctx context.Context, servicePath string, serviceMethod string, payload []byte) (meta map[string]string, resp []byte, err error) {
	return s.RPC(ctx, servicePath, serviceMethod, payload, protocol.ProtoBuffer, false)
}

func (s *RPCXClients) RPCJson(ctx context.Context, servicePath string, serviceMethod string, payload []byte) (meta map[string]string, resp []byte, err error) {
	return s.RPC(ctx, servicePath, serviceMethod, payload, protocol.JSON, false)
}

func (s *RPCXClients) GetXClient(servicePath string) (xc client.XClient, err error) {
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		if e := recover(); e != nil {
			if ee, ok := e.(error); ok {
				err = ee
				return
			}

			err = fmt.Errorf("failed to get xclient: %v", e)
		}
	}()

	if s.xclients[servicePath] == nil {
		d, err := s.serviceDiscovery.Clone(servicePath)
		if err != nil {
			return nil, err
		}
		s.xclients[servicePath] = client.NewXClient(servicePath, s.FailMode, s.SelectMode, d, s.Option)
	}
	xc = s.xclients[servicePath]

	return xc, err
}

func (s *RPCXClients) setFailMode(failMode client.FailMode) {
	s.FailMode = failMode
}

func (s *RPCXClients) setSelectMode(selectMode client.SelectMode) {
	s.SelectMode = selectMode
}

func (s *RPCXClients) setOption(option client.Option) {
	s.Option = option
}
