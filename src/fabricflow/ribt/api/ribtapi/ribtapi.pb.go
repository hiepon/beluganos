// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ribtapi.proto

package ribtapi

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

type TunnelRoute struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Nexthop              string   `protobuf:"bytes,2,opt,name=nexthop,proto3" json:"nexthop,omitempty"`
	Family               uint32   `protobuf:"varint,3,opt,name=family,proto3" json:"family,omitempty"`
	TunnelType           int32    `protobuf:"varint,4,opt,name=tunnel_type,json=tunnelType,proto3" json:"tunnel_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TunnelRoute) Reset()         { *m = TunnelRoute{} }
func (m *TunnelRoute) String() string { return proto.CompactTextString(m) }
func (*TunnelRoute) ProtoMessage()    {}
func (*TunnelRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0ddecee513f04d2, []int{0}
}

func (m *TunnelRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TunnelRoute.Unmarshal(m, b)
}
func (m *TunnelRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TunnelRoute.Marshal(b, m, deterministic)
}
func (m *TunnelRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TunnelRoute.Merge(m, src)
}
func (m *TunnelRoute) XXX_Size() int {
	return xxx_messageInfo_TunnelRoute.Size(m)
}
func (m *TunnelRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TunnelRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TunnelRoute proto.InternalMessageInfo

func (m *TunnelRoute) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *TunnelRoute) GetNexthop() string {
	if m != nil {
		return m.Nexthop
	}
	return ""
}

func (m *TunnelRoute) GetFamily() uint32 {
	if m != nil {
		return m.Family
	}
	return 0
}

func (m *TunnelRoute) GetTunnelType() int32 {
	if m != nil {
		return m.TunnelType
	}
	return 0
}

type GetTunnelsRequest struct {
	KeyType              string   `protobuf:"bytes,1,opt,name=key_type,json=keyType,proto3" json:"key_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTunnelsRequest) Reset()         { *m = GetTunnelsRequest{} }
func (m *GetTunnelsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTunnelsRequest) ProtoMessage()    {}
func (*GetTunnelsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0ddecee513f04d2, []int{1}
}

func (m *GetTunnelsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTunnelsRequest.Unmarshal(m, b)
}
func (m *GetTunnelsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTunnelsRequest.Marshal(b, m, deterministic)
}
func (m *GetTunnelsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTunnelsRequest.Merge(m, src)
}
func (m *GetTunnelsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTunnelsRequest.Size(m)
}
func (m *GetTunnelsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTunnelsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTunnelsRequest proto.InternalMessageInfo

func (m *GetTunnelsRequest) GetKeyType() string {
	if m != nil {
		return m.KeyType
	}
	return ""
}

type GetTunnelsReply struct {
	Id                   uint32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 int32                   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Remote               string                  `protobuf:"bytes,3,opt,name=remote,proto3" json:"remote,omitempty"`
	Local                string                  `protobuf:"bytes,4,opt,name=local,proto3" json:"local,omitempty"`
	Routes               map[string]*TunnelRoute `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GetTunnelsReply) Reset()         { *m = GetTunnelsReply{} }
func (m *GetTunnelsReply) String() string { return proto.CompactTextString(m) }
func (*GetTunnelsReply) ProtoMessage()    {}
func (*GetTunnelsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0ddecee513f04d2, []int{2}
}

func (m *GetTunnelsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTunnelsReply.Unmarshal(m, b)
}
func (m *GetTunnelsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTunnelsReply.Marshal(b, m, deterministic)
}
func (m *GetTunnelsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTunnelsReply.Merge(m, src)
}
func (m *GetTunnelsReply) XXX_Size() int {
	return xxx_messageInfo_GetTunnelsReply.Size(m)
}
func (m *GetTunnelsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTunnelsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTunnelsReply proto.InternalMessageInfo

func (m *GetTunnelsReply) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetTunnelsReply) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *GetTunnelsReply) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

func (m *GetTunnelsReply) GetLocal() string {
	if m != nil {
		return m.Local
	}
	return ""
}

func (m *GetTunnelsReply) GetRoutes() map[string]*TunnelRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

func init() {
	proto.RegisterType((*TunnelRoute)(nil), "ribtapi.TunnelRoute")
	proto.RegisterType((*GetTunnelsRequest)(nil), "ribtapi.GetTunnelsRequest")
	proto.RegisterType((*GetTunnelsReply)(nil), "ribtapi.GetTunnelsReply")
	proto.RegisterMapType((map[string]*TunnelRoute)(nil), "ribtapi.GetTunnelsReply.RoutesEntry")
}

func init() { proto.RegisterFile("ribtapi.proto", fileDescriptor_f0ddecee513f04d2) }

var fileDescriptor_f0ddecee513f04d2 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xdd, 0xb4, 0x69, 0xec, 0x84, 0xfa, 0x67, 0x28, 0xb2, 0xf6, 0x62, 0x28, 0x1e, 0x82,
	0x87, 0x20, 0xf5, 0x22, 0xe2, 0x45, 0x41, 0xc5, 0x93, 0xb8, 0xf4, 0x2e, 0xa9, 0x9d, 0xe2, 0xd2,
	0x34, 0x59, 0x93, 0x8d, 0x74, 0xbf, 0xb6, 0x9f, 0x40, 0xb2, 0x9b, 0x6a, 0x11, 0xbd, 0xcd, 0x9b,
	0xcc, 0x9b, 0xf9, 0xe5, 0x25, 0x30, 0x28, 0xe5, 0x4c, 0xa7, 0x4a, 0x26, 0xaa, 0x2c, 0x74, 0x81,
	0x41, 0x2b, 0xc7, 0x6b, 0x08, 0xa7, 0x75, 0x9e, 0x53, 0x26, 0x8a, 0x5a, 0x13, 0x1e, 0x41, 0x4f,
	0x95, 0xb4, 0x90, 0x6b, 0xce, 0x22, 0x16, 0xf7, 0x45, 0xab, 0x90, 0x43, 0x90, 0xd3, 0x5a, 0xbf,
	0x15, 0x8a, 0x7b, 0xf6, 0xc1, 0x46, 0x36, 0x8e, 0x45, 0xba, 0x92, 0x99, 0xe1, 0x9d, 0x88, 0xc5,
	0x03, 0xd1, 0x2a, 0x3c, 0x81, 0x50, 0xdb, 0xc5, 0x2f, 0xda, 0x28, 0xe2, 0xdd, 0x88, 0xc5, 0xbe,
	0x00, 0xd7, 0x9a, 0x1a, 0x45, 0xe3, 0x04, 0x0e, 0x1f, 0x48, 0xbb, 0xe3, 0x95, 0xa0, 0xf7, 0x9a,
	0x2a, 0x8d, 0xc7, 0xb0, 0xbb, 0x24, 0xe3, 0x2c, 0x8e, 0x20, 0x58, 0x92, 0xb1, 0xf3, 0x9f, 0x0c,
	0xf6, 0xb7, 0x0d, 0x2a, 0x33, 0xb8, 0x07, 0x9e, 0x9c, 0xdb, 0xc1, 0x81, 0xf0, 0xe4, 0x1c, 0x11,
	0xba, 0xd6, 0xea, 0xd9, 0x6b, 0xb6, 0x6e, 0x00, 0x4b, 0x5a, 0x15, 0x9a, 0x2c, 0x60, 0x5f, 0xb4,
	0x0a, 0x87, 0xe0, 0x67, 0xc5, 0x6b, 0x9a, 0x59, 0xb4, 0xbe, 0x70, 0x02, 0xaf, 0xa1, 0x57, 0x36,
	0x49, 0x54, 0xdc, 0x8f, 0x3a, 0x71, 0x38, 0x39, 0x4d, 0x36, 0xc1, 0xfd, 0xba, 0x9d, 0xd8, 0xc0,
	0xaa, 0xbb, 0x5c, 0x97, 0x46, 0xb4, 0x9e, 0xd1, 0x13, 0x84, 0x5b, 0x6d, 0x3c, 0x80, 0xce, 0x92,
	0x4c, 0xfb, 0x22, 0x4d, 0x89, 0x67, 0xe0, 0x7f, 0xa4, 0x59, 0xed, 0x08, 0xc3, 0xc9, 0xf0, 0x7b,
	0xfb, 0xd6, 0x47, 0x10, 0x6e, 0xe4, 0xca, 0xbb, 0x64, 0x93, 0x67, 0x08, 0xc4, 0xe3, 0xed, 0xf4,
	0x46, 0x49, 0xbc, 0x07, 0xf8, 0x41, 0xc0, 0xd1, 0x9f, 0x5c, 0x36, 0xc4, 0x11, 0xff, 0x8f, 0x79,
	0xbc, 0x73, 0xce, 0x66, 0x3d, 0xfb, 0x07, 0x5c, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x35, 0xd2,
	0x0a, 0x6c, 0x12, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RIBTApiClient is the client API for RIBTApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RIBTApiClient interface {
	GetTunnels(ctx context.Context, in *GetTunnelsRequest, opts ...grpc.CallOption) (RIBTApi_GetTunnelsClient, error)
}

type rIBTApiClient struct {
	cc *grpc.ClientConn
}

func NewRIBTApiClient(cc *grpc.ClientConn) RIBTApiClient {
	return &rIBTApiClient{cc}
}

func (c *rIBTApiClient) GetTunnels(ctx context.Context, in *GetTunnelsRequest, opts ...grpc.CallOption) (RIBTApi_GetTunnelsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RIBTApi_serviceDesc.Streams[0], "/ribtapi.RIBTApi/GetTunnels", opts...)
	if err != nil {
		return nil, err
	}
	x := &rIBTApiGetTunnelsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RIBTApi_GetTunnelsClient interface {
	Recv() (*GetTunnelsReply, error)
	grpc.ClientStream
}

type rIBTApiGetTunnelsClient struct {
	grpc.ClientStream
}

func (x *rIBTApiGetTunnelsClient) Recv() (*GetTunnelsReply, error) {
	m := new(GetTunnelsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RIBTApiServer is the server API for RIBTApi service.
type RIBTApiServer interface {
	GetTunnels(*GetTunnelsRequest, RIBTApi_GetTunnelsServer) error
}

// UnimplementedRIBTApiServer can be embedded to have forward compatible implementations.
type UnimplementedRIBTApiServer struct {
}

func (*UnimplementedRIBTApiServer) GetTunnels(req *GetTunnelsRequest, srv RIBTApi_GetTunnelsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTunnels not implemented")
}

func RegisterRIBTApiServer(s *grpc.Server, srv RIBTApiServer) {
	s.RegisterService(&_RIBTApi_serviceDesc, srv)
}

func _RIBTApi_GetTunnels_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTunnelsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RIBTApiServer).GetTunnels(m, &rIBTApiGetTunnelsServer{stream})
}

type RIBTApi_GetTunnelsServer interface {
	Send(*GetTunnelsReply) error
	grpc.ServerStream
}

type rIBTApiGetTunnelsServer struct {
	grpc.ServerStream
}

func (x *rIBTApiGetTunnelsServer) Send(m *GetTunnelsReply) error {
	return x.ServerStream.SendMsg(m)
}

var _RIBTApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ribtapi.RIBTApi",
	HandlerType: (*RIBTApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTunnels",
			Handler:       _RIBTApi_GetTunnels_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ribtapi.proto",
}
