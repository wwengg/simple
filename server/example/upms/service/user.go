// @Title
// @Description
// @Author  Wangwengang  2023/12/17 12:22
// @Update  Wangwengang  2023/12/17 12:22
package service

import (
	"context"
	"github.com/wwengg/simple/core/store"
	"github.com/wwengg/simple/proto/pbbase"
	"github.com/wwengg/simple/proto/pbgateway"
	"github.com/wwengg/simple/server/example/upms/global"
	"github.com/wwengg/simple/server/example/upms/model"
	"github.com/wwengg/simple/server/example/upms/pb"
	"github.com/wwengg/simple/server/example/upms/service/impl"
)

type User struct {
}

func (s *User) Health(ctx context.Context, args *pbbase.Empty, reply *pbgateway.Response) error {
	global.Log.Infof("Health")
	reply.Code = 200
	reply.Msg = "success2"
	return nil
}

func (s *User) CreateUser(ctx context.Context, args *pb.CreateUserArgs, reply *pbgateway.Response) error {
	//if args.Nick == "" {
	//	reply.Code = int32(pb.Code_PARAMS_ERR)
	//	return nil
	//}
	// 在内存中创建了个user
	user := model.SimpleUser{
		Nick: "测试",
	}

	if err := impl.CreateUser(&user); err != nil {
		global.Log.Infof("CreateUser err = %s", err.Error())
		reply.Code = int32(pb.Code_ADD_ERR)
		return nil
	}
	reply.Code = 200
	reply.Msg = "Success"

	return nil
}

func (s *User) UpdateUser(ctx context.Context, args *pb.UpdateUserArgs, reply *pbgateway.Response) error {
	if args.Id == 0 {
		reply.Code = int32(pb.Code_PARAMS_ERR)
		return nil
	}
	if err := impl.UpdateUserNick(args.Id, args.Nick); err != nil {
		global.Log.Infof("UpdateUser err = %s", err.Error())
		reply.Code = int32(pb.Code_UPDATE_ERR)
		return nil
	}

	reply.Code = 200
	reply.Msg = "SUCCESS"
	return nil
}

func (s *User) DeleteUser(ctx context.Context, args *pb.UpdateUserArgs, reply *pbgateway.Response) error {
	if args.Id == 0 {
		reply.Code = int32(pb.Code_PARAMS_ERR)
		return nil
	}
	user := model.SimpleUser{BASE_MODEL: store.BASE_MODEL{
		ID: args.Id,
	}}
	if err := impl.DeleteUser(user); err != nil {
		reply.Code = int32(pb.Code_DELETE_ERR)
		return nil
	}
	reply.Code = 200
	return nil
}
