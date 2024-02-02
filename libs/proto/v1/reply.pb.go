// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: libs/proto/v1/reply.proto

package v1

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

type DefaultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefaultReply) Reset() {
	*x = DefaultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_libs_proto_v1_reply_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultReply) ProtoMessage() {}

func (x *DefaultReply) ProtoReflect() protoreflect.Message {
	mi := &file_libs_proto_v1_reply_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultReply.ProtoReflect.Descriptor instead.
func (*DefaultReply) Descriptor() ([]byte, []int) {
	return file_libs_proto_v1_reply_proto_rawDescGZIP(), []int{0}
}

func (x *DefaultReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_libs_proto_v1_reply_proto protoreflect.FileDescriptor

var file_libs_proto_v1_reply_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6b, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x0c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x17,
	0x5a, 0x15, 0x6b, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_libs_proto_v1_reply_proto_rawDescOnce sync.Once
	file_libs_proto_v1_reply_proto_rawDescData = file_libs_proto_v1_reply_proto_rawDesc
)

func file_libs_proto_v1_reply_proto_rawDescGZIP() []byte {
	file_libs_proto_v1_reply_proto_rawDescOnce.Do(func() {
		file_libs_proto_v1_reply_proto_rawDescData = protoimpl.X.CompressGZIP(file_libs_proto_v1_reply_proto_rawDescData)
	})
	return file_libs_proto_v1_reply_proto_rawDescData
}

var file_libs_proto_v1_reply_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_libs_proto_v1_reply_proto_goTypes = []interface{}{
	(*DefaultReply)(nil), // 0: kore.v1.DefaultReply
}
var file_libs_proto_v1_reply_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_libs_proto_v1_reply_proto_init() }
func file_libs_proto_v1_reply_proto_init() {
	if File_libs_proto_v1_reply_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_libs_proto_v1_reply_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultReply); i {
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
			RawDescriptor: file_libs_proto_v1_reply_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_libs_proto_v1_reply_proto_goTypes,
		DependencyIndexes: file_libs_proto_v1_reply_proto_depIdxs,
		MessageInfos:      file_libs_proto_v1_reply_proto_msgTypes,
	}.Build()
	File_libs_proto_v1_reply_proto = out.File
	file_libs_proto_v1_reply_proto_rawDesc = nil
	file_libs_proto_v1_reply_proto_goTypes = nil
	file_libs_proto_v1_reply_proto_depIdxs = nil
}