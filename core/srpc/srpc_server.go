// @Title
// @Description
// @Author  Wangwengang  2023/12/12 21:43
// @Update  Wangwengang  2023/12/12 21:43
package srpc

import (
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
	"github.com/wwengg/simple/core/sconfig"
	"time"
)

type SRPCServer interface {
	RegisterName(name string, rcvr interface{}, metadata string) error
	Serve(network, address string) (err error)
}

func AddRegistryPlugin(s *server.Server, addr string) {

	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		EtcdServers:    sconfig.S_CONF.RPC.RegisterAddr,
		BasePath:       sconfig.S_CONF.RPC.BasePath,
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		panic(err)
	}
	s.Plugins.Add(r)
}
