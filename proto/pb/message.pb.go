// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: proto/message.proto

package pb

import (
	context "context"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users   []uint64 `protobuf:"varint,1,rep,packed,name=Users,proto3" json:"Users,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Type    string   `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetUsers() []uint64 {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Request) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Request) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users   []uint64 `protobuf:"varint,1,rep,packed,name=Users,proto3" json:"Users,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetUsers() []uint64 {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_message_proto protoreflect.FileDescriptor

var file_proto_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x46, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x08,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_proto_rawDescOnce sync.Once
	file_proto_message_proto_rawDescData = file_proto_message_proto_rawDesc
)

func file_proto_message_proto_rawDescGZIP() []byte {
	file_proto_message_proto_rawDescOnce.Do(func() {
		file_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_proto_rawDescData)
	})
	return file_proto_message_proto_rawDescData
}

var file_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_message_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: message.Request
	(*Response)(nil), // 1: message.Response
}
var file_proto_message_proto_depIdxs = []int32{
	0, // 0: message.DeliveryService.Delivery:input_type -> message.Request
	1, // 1: message.DeliveryService.Delivery:output_type -> message.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_message_proto_init() }
func file_proto_message_proto_init() {
	if File_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_message_proto_goTypes,
		DependencyIndexes: file_proto_message_proto_depIdxs,
		MessageInfos:      file_proto_message_proto_msgTypes,
	}.Build()
	File_proto_message_proto = out.File
	file_proto_message_proto_rawDesc = nil
	file_proto_message_proto_goTypes = nil
	file_proto_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeliveryServiceClient is the client API for DeliveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeliveryServiceClient interface {
	Delivery(ctx context.Context, opts ...grpc.CallOption) (DeliveryService_DeliveryClient, error)
}

type deliveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryServiceClient(cc grpc.ClientConnInterface) DeliveryServiceClient {
	return &deliveryServiceClient{cc}
}

func (c *deliveryServiceClient) Delivery(ctx context.Context, opts ...grpc.CallOption) (DeliveryService_DeliveryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DeliveryService_serviceDesc.Streams[0], "/message.DeliveryService/Delivery", opts...)
	if err != nil {
		return nil, err
	}
	x := &deliveryServiceDeliveryClient{stream}
	return x, nil
}

type DeliveryService_DeliveryClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type deliveryServiceDeliveryClient struct {
	grpc.ClientStream
}

func (x *deliveryServiceDeliveryClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *deliveryServiceDeliveryClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DeliveryServiceServer is the server API for DeliveryService service.
type DeliveryServiceServer interface {
	Delivery(DeliveryService_DeliveryServer) error
}

// UnimplementedDeliveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeliveryServiceServer struct {
}

func (*UnimplementedDeliveryServiceServer) Delivery(DeliveryService_DeliveryServer) error {
	return status.Errorf(codes.Unimplemented, "method Delivery not implemented")
}

func RegisterDeliveryServiceServer(s *grpc.Server, srv DeliveryServiceServer) {
	s.RegisterService(&_DeliveryService_serviceDesc, srv)
}

func _DeliveryService_Delivery_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeliveryServiceServer).Delivery(&deliveryServiceDeliveryServer{stream})
}

type DeliveryService_DeliveryServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type deliveryServiceDeliveryServer struct {
	grpc.ServerStream
}

func (x *deliveryServiceDeliveryServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *deliveryServiceDeliveryServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DeliveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.DeliveryService",
	HandlerType: (*DeliveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Delivery",
			Handler:       _DeliveryService_Delivery_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/message.proto",
}