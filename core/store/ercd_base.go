// @Title
// @Description
// @Author  Wangwengang  2023/12/23 19:14
// @Update  Wangwengang  2023/12/23 19:14
package store

import clientv3 "go.etcd.io/etcd/client/v3"

type EtcdBase struct {
	EtcdCli *clientv3.Client
}
