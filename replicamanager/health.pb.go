// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: health.proto

package replicamanager

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{0}
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1}
}

func (x *HealthResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_health_proto protoreflect.FileDescriptor

var file_health_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x0f,
	0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x26, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x51, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x47, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x3b,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_proto_rawDescOnce sync.Once
	file_health_proto_rawDescData = file_health_proto_rawDesc
)

func file_health_proto_rawDescGZIP() []byte {
	file_health_proto_rawDescOnce.Do(func() {
		file_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_proto_rawDescData)
	})
	return file_health_proto_rawDescData
}

var file_health_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_health_proto_goTypes = []interface{}{
	(*HealthRequest)(nil),  // 0: replicamanager.HealthRequest
	(*HealthResponse)(nil), // 1: replicamanager.HealthResponse
}
var file_health_proto_depIdxs = []int32{
	0, // 0: replicamanager.Health.health:input_type -> replicamanager.HealthRequest
	1, // 1: replicamanager.Health.health:output_type -> replicamanager.HealthResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_health_proto_init() }
func file_health_proto_init() {
	if File_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
			RawDescriptor: file_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_proto_goTypes,
		DependencyIndexes: file_health_proto_depIdxs,
		MessageInfos:      file_health_proto_msgTypes,
	}.Build()
	File_health_proto = out.File
	file_health_proto_rawDesc = nil
	file_health_proto_goTypes = nil
	file_health_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/replicamanager.Health/health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/replicamanager.Health/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "replicamanager.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "health",
			Handler:    _Health_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}