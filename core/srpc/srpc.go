// @Title
// @Description
// @Author  Wangwengang  2023/12/12 00:08
// @Update  Wangwengang  2023/12/12 00:08
package srpc

import (
	"context"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

type SRPC interface {
	RPC(ctx context.Context, servicePath string, serviceMethod string, payload []byte, serializeType protocol.SerializeType, oneway bool) (meta map[string]string, resp []byte, err error)
	RPCProtobuf(ctx context.Context, servicePath string, serviceMethod string, payload []byte) (meta map[string]string, resp []byte, err error)
	RPCJson(ctx context.Context, servicePath string, serviceMethod string, payload []byte) (meta map[string]string, resp []byte, err error)
	GetXClient(servicePath string) (xc client.XClient, err error)
}
