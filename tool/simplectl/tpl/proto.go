// @Title
// @Description
// @Author  Wangwengang  2023/12/25 20:14
// @Update  Wangwengang  2023/12/25 20:14
package tpl

func CommonProtoTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
*/
syntax = "proto3";

package pbcommon;

option go_package = "{{ .PkgName }}/proto/pbcommon";

enum ErrCode {
  ErrCodeNone = 0;
  ErrCodeSuccess = 200; // 操作成功
  ErrCodeFail = 500; // 操作失败
  ErrCodeUnknown = 501; // 未知错误
  ErrCodeInternal = 502; // 内部错误
  ErrCodeInvalid = 503; // 无效数据
  ErrCodeInvalidParam = 504; // 无效参数
  ErrCodeParamError = 505; // 参数错误

  ErrCodeFindError = 1001; // 查询失败
  ErrCodeCreateError = 1002; // 创建失败
  ErrCodeDeleteError = 1003; // 删除失败
  ErrCodeUpdateError = 1004; // 更新失败
}

message CommonResult{
  ErrCode code = 1;
  string msg = 2;
}

message IdArgs {
  int64 id = 1;
  string idStr = 2;
  repeated int64 ids = 3;
  repeated string idStrs = 4;
}

message PageInfo{
  int32 page = 1;
  int32 pageSize = 2;
}`)
}

func NewProtoTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
*/
syntax = "proto3";

package pb{{ .CmdName }};

option go_package = "{{ .PkgName }}/proto/pb{{ .CmdName }}";

import "pbcommon/pbcommon.proto";


message {{ .CmdParent }}Model{
  int64 id = 1;
  string createdAt = 2;
  string updatedAt = 3;

}

message Find{{ .CmdParent }}Args{
  pbcommon.PageInfo pageInfo = 1;
  {{ .CmdParent }}Model query = 2;
}

message Find{{ .CmdParent }}Replay{
  pbcommon.ErrCode code = 1;
  string msg = 2;
  {{ .CmdParent }}Model data =3;
  repeated {{ .CmdParent }}Model list = 4;
}

service {{ .CmdParent }} {
  rpc Create{{ .CmdParent }}({{ .CmdParent }}Model) returns(pbcommon.CommonResult){}
  rpc Update{{ .CmdParent }}({{ .CmdParent }}Model) returns(pbcommon.CommonResult){}
  rpc Delete{{ .CmdParent }}(pbcommon.IdArgs) returns(pbcommon.CommonResult){}
  rpc Find{{ .CmdParent }}ById(pbcommon.IdArgs) returns(Find{{ .CmdParent }}Replay){}
  rpc Find{{ .CmdParent }}List(Find{{ .CmdParent }}Args) returns(Find{{ .CmdParent }}Replay){}
}`)
}
