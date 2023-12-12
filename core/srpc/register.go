// @Title
// @Description
// @Author  Wangwengang  2023/12/10 18:54
// @Update  Wangwengang  2023/12/10 18:54
package srpc

import (
	"fmt"
	etcdclient "github.com/rpcxio/rpcx-etcd/client"
	"github.com/smallnest/rpcx/client"
)

func CreateServiceDiscovery(regAddr []string, regType string, basePath string) (client.ServiceDiscovery, error) {
	switch regType {
	case "peer2peer": // peer2peer://127.0.0.1:8972
		return client.NewPeer2PeerDiscovery("tcp@"+regAddr[0], "")
	case "multiple":
		var pairs []*client.KVPair
		for _, v := range regAddr {
			pairs = append(pairs, &client.KVPair{Key: v})
		}
		return client.NewMultipleServersDiscovery(pairs)
	case "etcd":
		return etcdclient.NewEtcdDiscoveryTemplate(basePath, regAddr, true, nil)
	case "etcdv3":
		return etcdclient.NewEtcdV3DiscoveryTemplate(basePath, regAddr, true, nil)
	default:
		return nil, fmt.Errorf("wrong registry type %s. only support peer2peer,multiple,etcd and etcdv3,", regType)
	}
}
