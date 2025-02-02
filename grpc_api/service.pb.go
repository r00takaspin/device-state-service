// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceStateRequest struct {
	DeviceId             string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceStateRequest) Reset()         { *m = DeviceStateRequest{} }
func (m *DeviceStateRequest) String() string { return proto.CompactTextString(m) }
func (*DeviceStateRequest) ProtoMessage()    {}
func (*DeviceStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *DeviceStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceStateRequest.Unmarshal(m, b)
}
func (m *DeviceStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceStateRequest.Marshal(b, m, deterministic)
}
func (m *DeviceStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceStateRequest.Merge(m, src)
}
func (m *DeviceStateRequest) XXX_Size() int {
	return xxx_messageInfo_DeviceStateRequest.Size(m)
}
func (m *DeviceStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceStateRequest proto.InternalMessageInfo

func (m *DeviceStateRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type DeviceStateResponse struct {
	State                string   `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceStateResponse) Reset()         { *m = DeviceStateResponse{} }
func (m *DeviceStateResponse) String() string { return proto.CompactTextString(m) }
func (*DeviceStateResponse) ProtoMessage()    {}
func (*DeviceStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *DeviceStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceStateResponse.Unmarshal(m, b)
}
func (m *DeviceStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceStateResponse.Marshal(b, m, deterministic)
}
func (m *DeviceStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceStateResponse.Merge(m, src)
}
func (m *DeviceStateResponse) XXX_Size() int {
	return xxx_messageInfo_DeviceStateResponse.Size(m)
}
func (m *DeviceStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceStateResponse proto.InternalMessageInfo

func (m *DeviceStateResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *DeviceStateResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*DeviceStateRequest)(nil), "service.DeviceStateRequest")
	proto.RegisterType((*DeviceStateResponse)(nil), "service.DeviceStateResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x0c, 0xb9,
	0x84, 0x5c, 0x52, 0x41, 0xac, 0xe0, 0x92, 0xc4, 0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x69, 0x2e, 0xce, 0x14, 0xb0, 0x68, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x07, 0x44, 0xc0, 0x33, 0x45, 0xc9, 0x93, 0x4b, 0x18, 0x45, 0x4b, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x31, 0x48, 0x00, 0xaa, 0x1e, 0xc2, 0x11, 0x92, 0xe1,
	0xe2, 0x2c, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x52, 0x60, 0xd4, 0x60,
	0x0e, 0x42, 0x08, 0x18, 0x25, 0xa3, 0xd8, 0x1e, 0x0c, 0x71, 0x93, 0x90, 0x2f, 0x17, 0x9f, 0x7b,
	0x6a, 0x09, 0x92, 0x84, 0x90, 0xb4, 0x1e, 0xcc, 0xf9, 0x98, 0x8e, 0x95, 0x92, 0xc1, 0x2e, 0x09,
	0x71, 0x96, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xcb, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce,
	0x33, 0x12, 0x66, 0x03, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceStateServiceClient is the client API for DeviceStateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceStateServiceClient interface {
	GetDeviceState(ctx context.Context, in *DeviceStateRequest, opts ...grpc.CallOption) (*DeviceStateResponse, error)
}

type deviceStateServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceStateServiceClient(cc *grpc.ClientConn) DeviceStateServiceClient {
	return &deviceStateServiceClient{cc}
}

func (c *deviceStateServiceClient) GetDeviceState(ctx context.Context, in *DeviceStateRequest, opts ...grpc.CallOption) (*DeviceStateResponse, error) {
	out := new(DeviceStateResponse)
	err := c.cc.Invoke(ctx, "/service.DeviceStateService/GetDeviceState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceStateServiceServer is the server API for DeviceStateService service.
type DeviceStateServiceServer interface {
	GetDeviceState(context.Context, *DeviceStateRequest) (*DeviceStateResponse, error)
}

func RegisterDeviceStateServiceServer(s *grpc.Server, srv DeviceStateServiceServer) {
	s.RegisterService(&_DeviceStateService_serviceDesc, srv)
}

func _DeviceStateService_GetDeviceState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceStateServiceServer).GetDeviceState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.DeviceStateService/GetDeviceState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceStateServiceServer).GetDeviceState(ctx, req.(*DeviceStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceStateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.DeviceStateService",
	HandlerType: (*DeviceStateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceState",
			Handler:    _DeviceStateService_GetDeviceState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
