package sbus

import (
	"fmt"
	"github.com/wwengg/simple/core/smsg"
	"github.com/wwengg/simple/proto/pbbase"
	"testing"
)

func TestNewNsqDataPack(t *testing.T) {
	pack := NewNsqDataPack()
	md := map[string]string{}
	md["test1"] = "test1"
	proto := pbbase.PageInfo{
		Page:     1,
		PageSize: 1,
		Keyword:  "test",
	}
	data, _ := proto.Marshal()
	msg := NewNSQMsg(1, 1, smsg.ProtoBuffer, md, data)
	if packData, err := pack.Pack(msg); err != nil {
		panic(err)
	} else {
		if msg, err := pack.Unpack(packData); err != nil {
			panic(err)
		} else {
			fmt.Println(msg.GetMeta())
			fmt.Println(msg.GetCompressType())
			fmt.Println(msg.GetSerializeType())

		}
	}
}
