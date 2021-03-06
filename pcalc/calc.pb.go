// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pcalc/calc.proto

package pcalc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CalcRequest struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcRequest) Reset()         { *m = CalcRequest{} }
func (m *CalcRequest) String() string { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()    {}
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0855ef20a743e93c, []int{0}
}

func (m *CalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcRequest.Unmarshal(m, b)
}
func (m *CalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcRequest.Marshal(b, m, deterministic)
}
func (m *CalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcRequest.Merge(m, src)
}
func (m *CalcRequest) XXX_Size() int {
	return xxx_messageInfo_CalcRequest.Size(m)
}
func (m *CalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcRequest proto.InternalMessageInfo

func (m *CalcRequest) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CalcRequest) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type CalcResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcResponse) Reset()         { *m = CalcResponse{} }
func (m *CalcResponse) String() string { return proto.CompactTextString(m) }
func (*CalcResponse) ProtoMessage()    {}
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0855ef20a743e93c, []int{1}
}

func (m *CalcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcResponse.Unmarshal(m, b)
}
func (m *CalcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcResponse.Marshal(b, m, deterministic)
}
func (m *CalcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcResponse.Merge(m, src)
}
func (m *CalcResponse) XXX_Size() int {
	return xxx_messageInfo_CalcResponse.Size(m)
}
func (m *CalcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalcResponse proto.InternalMessageInfo

func (m *CalcResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalcRequest)(nil), "CalcRequest")
	proto.RegisterType((*CalcResponse)(nil), "CalcResponse")
}

func init() { proto.RegisterFile("pcalc/calc.proto", fileDescriptor_0855ef20a743e93c) }

var fileDescriptor_0855ef20a743e93c = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0x4e, 0xcc,
	0x49, 0xd6, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x4a, 0x9a, 0x5c, 0xdc, 0xce, 0x89,
	0x39, 0xc9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x3c, 0x5c, 0x8c, 0x15, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x8c, 0x15, 0x20, 0x5e, 0xa5, 0x04, 0x13, 0x84, 0x57, 0xa9, 0xa4,
	0xc6, 0xc5, 0x03, 0x51, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94,
	0x5a, 0x5c, 0x9a, 0x53, 0x02, 0xd5, 0x00, 0xe5, 0x19, 0x99, 0x70, 0x71, 0x81, 0xd4, 0x95, 0xe6,
	0x24, 0x96, 0xe4, 0x17, 0x09, 0xa9, 0x71, 0xb1, 0x86, 0xe4, 0x17, 0xe4, 0x24, 0x0a, 0xf1, 0xe8,
	0x21, 0x59, 0x24, 0xc5, 0xab, 0x87, 0x6c, 0x96, 0x12, 0x83, 0x13, 0x7b, 0x14, 0x2b, 0xd8, 0x71,
	0x49, 0x6c, 0x60, 0x87, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xfe, 0xfa, 0x65, 0xac,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	Topla(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Topla(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/Calculator/Topla", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	Topla(context.Context, *CalcRequest) (*CalcResponse, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) Topla(ctx context.Context, req *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Topla not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Topla_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Topla(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calculator/Topla",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Topla(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Topla",
			Handler:    _Calculator_Topla_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pcalc/calc.proto",
}
