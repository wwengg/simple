// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.3
// source: pb/upms.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// upms的错误码 1000 ~ 1999
type Code int32

const (
	Code_None       Code = 0
	Code_PARAMS_ERR Code = 1001 // 参数错误
	Code_ADD_ERR    Code = 1011 // 创建失败
	Code_UPDATE_ERR Code = 1012 // 更新失败
	Code_DELETE_ERR Code = 1013 // 删除失败
	Code_GET_ERR    Code = 1014 // 查询失败
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:    "None",
		1001: "PARAMS_ERR",
		1011: "ADD_ERR",
		1012: "UPDATE_ERR",
		1013: "DELETE_ERR",
		1014: "GET_ERR",
	}
	Code_value = map[string]int32{
		"None":       0,
		"PARAMS_ERR": 1001,
		"ADD_ERR":    1011,
		"UPDATE_ERR": 1012,
		"DELETE_ERR": 1013,
		"GET_ERR":    1014,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_upms_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_pb_upms_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_pb_upms_proto_rawDescGZIP(), []int{0}
}

// http   /user/createUser[post]
// CreateUserReply 用Empty
type CreateUserArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nick string `protobuf:"bytes,1,opt,name=nick,proto3" json:"nick,omitempty"`
}

func (x *CreateUserArgs) Reset() {
	*x = CreateUserArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_upms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserArgs) ProtoMessage() {}

func (x *CreateUserArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pb_upms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserArgs.ProtoReflect.Descriptor instead.
func (*CreateUserArgs) Descriptor() ([]byte, []int) {
	return file_pb_upms_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserArgs) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

// http /user/updateUser [post]
// Cre
type UpdateUserArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nick  string `protobuf:"bytes,2,opt,name=nick,proto3" json:"nick,omitempty"`
	Phone int64  `protobuf:"varint,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *UpdateUserArgs) Reset() {
	*x = UpdateUserArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_upms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserArgs) ProtoMessage() {}

func (x *UpdateUserArgs) ProtoReflect() protoreflect.Message {
	mi := &file_pb_upms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserArgs.ProtoReflect.Descriptor instead.
func (*UpdateUserArgs) Descriptor() ([]byte, []int) {
	return file_pb_upms_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUserArgs) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateUserArgs) GetNick() string {
	if x != nil {
		return x.Nick
	}
	return ""
}

func (x *UpdateUserArgs) GetPhone() int64 {
	if x != nil {
		return x.Phone
	}
	return 0
}

var File_pb_upms_proto protoreflect.FileDescriptor

var file_pb_upms_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x75, 0x70, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x22, 0x4a, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x69, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x69, 0x63, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2a, 0x5f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0a, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x53, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xe9, 0x07, 0x12, 0x0c, 0x0a, 0x07, 0x41, 0x44, 0x44, 0x5f,
	0x45, 0x52, 0x52, 0x10, 0xf3, 0x07, 0x12, 0x0f, 0x0a, 0x0a, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x45, 0x52, 0x52, 0x10, 0xf4, 0x07, 0x12, 0x0f, 0x0a, 0x0a, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xf5, 0x07, 0x12, 0x0c, 0x0a, 0x07, 0x47, 0x45, 0x54, 0x5f,
	0x45, 0x52, 0x52, 0x10, 0xf6, 0x07, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_upms_proto_rawDescOnce sync.Once
	file_pb_upms_proto_rawDescData = file_pb_upms_proto_rawDesc
)

func file_pb_upms_proto_rawDescGZIP() []byte {
	file_pb_upms_proto_rawDescOnce.Do(func() {
		file_pb_upms_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_upms_proto_rawDescData)
	})
	return file_pb_upms_proto_rawDescData
}

var file_pb_upms_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_upms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_upms_proto_goTypes = []interface{}{
	(Code)(0),              // 0: pb.Code
	(*CreateUserArgs)(nil), // 1: pb.CreateUserArgs
	(*UpdateUserArgs)(nil), // 2: pb.UpdateUserArgs
}
var file_pb_upms_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_upms_proto_init() }
func file_pb_upms_proto_init() {
	if File_pb_upms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_upms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_upms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_upms_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_upms_proto_goTypes,
		DependencyIndexes: file_pb_upms_proto_depIdxs,
		EnumInfos:         file_pb_upms_proto_enumTypes,
		MessageInfos:      file_pb_upms_proto_msgTypes,
	}.Build()
	File_pb_upms_proto = out.File
	file_pb_upms_proto_rawDesc = nil
	file_pb_upms_proto_goTypes = nil
	file_pb_upms_proto_depIdxs = nil
}
