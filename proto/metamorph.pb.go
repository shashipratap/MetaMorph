// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: proto/metamorph.proto

package proto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID    string `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	NodeSpec  []byte `protobuf:"bytes,2,opt,name=NodeSpec,proto3" json:"NodeSpec,omitempty"`
	Task      []byte `protobuf:"bytes,3,opt,name=Task,proto3" json:"Task,omitempty"`
	NodeState string `protobuf:"bytes,4,opt,name=NodeState,proto3" json:"NodeState,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metamorph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metamorph_proto_msgTypes[0]
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
	return file_proto_metamorph_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *Request) GetNodeSpec() []byte {
	if x != nil {
		return x.NodeSpec
	}
	return nil
}

func (x *Request) GetTask() []byte {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *Request) GetNodeState() string {
	if x != nil {
		return x.NodeState
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Res    []byte `protobuf:"bytes,2,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_metamorph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_metamorph_proto_msgTypes[1]
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
	return file_proto_metamorph_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *Response) GetRes() []byte {
	if x != nil {
		return x.Res
	}
	return nil
}

var File_proto_metamorph_proto protoreflect.FileDescriptor

var file_proto_metamorph_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x70,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22,
	0x34, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x72, 0x65, 0x73, 0x32, 0xa2, 0x03, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x08, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x55, 0x55, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x48, 0x57,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x57, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa4, 0x01, 0x0a, 0x0c, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_metamorph_proto_rawDescOnce sync.Once
	file_proto_metamorph_proto_rawDescData = file_proto_metamorph_proto_rawDesc
)

func file_proto_metamorph_proto_rawDescGZIP() []byte {
	file_proto_metamorph_proto_rawDescOnce.Do(func() {
		file_proto_metamorph_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_metamorph_proto_rawDescData)
	})
	return file_proto_metamorph_proto_rawDescData
}

var file_proto_metamorph_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_metamorph_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: proto.Request
	(*Response)(nil), // 1: proto.Response
}
var file_proto_metamorph_proto_depIdxs = []int32{
	0,  // 0: proto.NodeService.Create:input_type -> proto.Request
	0,  // 1: proto.NodeService.Describe:input_type -> proto.Request
	0,  // 2: proto.NodeService.Update:input_type -> proto.Request
	0,  // 3: proto.NodeService.Delete:input_type -> proto.Request
	0,  // 4: proto.NodeService.List:input_type -> proto.Request
	0,  // 5: proto.NodeService.Deploy:input_type -> proto.Request
	0,  // 6: proto.NodeService.GetNodeUUID:input_type -> proto.Request
	0,  // 7: proto.NodeService.GetHWStatus:input_type -> proto.Request
	0,  // 8: proto.NodeService.UpdateHWStatus:input_type -> proto.Request
	0,  // 9: proto.AgentService.GetTasks:input_type -> proto.Request
	0,  // 10: proto.AgentService.UpdateTaskStatus:input_type -> proto.Request
	0,  // 11: proto.AgentService.UpdateNodeState:input_type -> proto.Request
	1,  // 12: proto.NodeService.Create:output_type -> proto.Response
	1,  // 13: proto.NodeService.Describe:output_type -> proto.Response
	1,  // 14: proto.NodeService.Update:output_type -> proto.Response
	1,  // 15: proto.NodeService.Delete:output_type -> proto.Response
	1,  // 16: proto.NodeService.List:output_type -> proto.Response
	1,  // 17: proto.NodeService.Deploy:output_type -> proto.Response
	1,  // 18: proto.NodeService.GetNodeUUID:output_type -> proto.Response
	1,  // 19: proto.NodeService.GetHWStatus:output_type -> proto.Response
	1,  // 20: proto.NodeService.UpdateHWStatus:output_type -> proto.Response
	1,  // 21: proto.AgentService.GetTasks:output_type -> proto.Response
	1,  // 22: proto.AgentService.UpdateTaskStatus:output_type -> proto.Response
	1,  // 23: proto.AgentService.UpdateNodeState:output_type -> proto.Response
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_metamorph_proto_init() }
func file_proto_metamorph_proto_init() {
	if File_proto_metamorph_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_metamorph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_metamorph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_metamorph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_metamorph_proto_goTypes,
		DependencyIndexes: file_proto_metamorph_proto_depIdxs,
		MessageInfos:      file_proto_metamorph_proto_msgTypes,
	}.Build()
	File_proto_metamorph_proto = out.File
	file_proto_metamorph_proto_rawDesc = nil
	file_proto_metamorph_proto_goTypes = nil
	file_proto_metamorph_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Describe(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	List(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Deploy(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetNodeUUID(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetHWStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateHWStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Describe(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Describe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Update(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Delete(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) List(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Deploy(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetNodeUUID(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/GetNodeUUID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetHWStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/GetHWStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) UpdateHWStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.NodeService/UpdateHWStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	Create(context.Context, *Request) (*Response, error)
	Describe(context.Context, *Request) (*Response, error)
	Update(context.Context, *Request) (*Response, error)
	Delete(context.Context, *Request) (*Response, error)
	List(context.Context, *Request) (*Response, error)
	Deploy(context.Context, *Request) (*Response, error)
	GetNodeUUID(context.Context, *Request) (*Response, error)
	GetHWStatus(context.Context, *Request) (*Response, error)
	UpdateHWStatus(context.Context, *Request) (*Response, error)
}

// UnimplementedNodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (*UnimplementedNodeServiceServer) Create(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNodeServiceServer) Describe(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Describe not implemented")
}
func (*UnimplementedNodeServiceServer) Update(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedNodeServiceServer) Delete(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedNodeServiceServer) List(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedNodeServiceServer) Deploy(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (*UnimplementedNodeServiceServer) GetNodeUUID(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeUUID not implemented")
}
func (*UnimplementedNodeServiceServer) GetHWStatus(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHWStatus not implemented")
}
func (*UnimplementedNodeServiceServer) UpdateHWStatus(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHWStatus not implemented")
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Describe(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Update(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Delete(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).List(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Deploy(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetNodeUUID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNodeUUID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/GetNodeUUID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNodeUUID(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetHWStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetHWStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/GetHWStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetHWStatus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_UpdateHWStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).UpdateHWStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/UpdateHWStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).UpdateHWStatus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NodeService_Create_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _NodeService_Describe_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NodeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NodeService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NodeService_List_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _NodeService_Deploy_Handler,
		},
		{
			MethodName: "GetNodeUUID",
			Handler:    _NodeService_GetNodeUUID_Handler,
		},
		{
			MethodName: "GetHWStatus",
			Handler:    _NodeService_GetHWStatus_Handler,
		},
		{
			MethodName: "UpdateHWStatus",
			Handler:    _NodeService_UpdateHWStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/metamorph.proto",
}

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	GetTasks(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateTaskStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateNodeState(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) GetTasks(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.AgentService/GetTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) UpdateTaskStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.AgentService/UpdateTaskStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) UpdateNodeState(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.AgentService/UpdateNodeState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	GetTasks(context.Context, *Request) (*Response, error)
	UpdateTaskStatus(context.Context, *Request) (*Response, error)
	UpdateNodeState(context.Context, *Request) (*Response, error)
}

// UnimplementedAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (*UnimplementedAgentServiceServer) GetTasks(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (*UnimplementedAgentServiceServer) UpdateTaskStatus(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskStatus not implemented")
}
func (*UnimplementedAgentServiceServer) UpdateNodeState(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeState not implemented")
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/GetTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).GetTasks(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_UpdateTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).UpdateTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/UpdateTaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).UpdateTaskStatus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_UpdateNodeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).UpdateNodeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/UpdateNodeState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).UpdateNodeState(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTasks",
			Handler:    _AgentService_GetTasks_Handler,
		},
		{
			MethodName: "UpdateTaskStatus",
			Handler:    _AgentService_UpdateTaskStatus_Handler,
		},
		{
			MethodName: "UpdateNodeState",
			Handler:    _AgentService_UpdateNodeState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/metamorph.proto",
}
