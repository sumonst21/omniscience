// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

package omniscience

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Service struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_4baa9ef27d785557, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServicesRequest) Reset()         { *m = ListServicesRequest{} }
func (m *ListServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListServicesRequest) ProtoMessage()    {}
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_4baa9ef27d785557, []int{1}
}
func (m *ListServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesRequest.Unmarshal(m, b)
}
func (m *ListServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesRequest.Marshal(b, m, deterministic)
}
func (dst *ListServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesRequest.Merge(dst, src)
}
func (m *ListServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListServicesRequest.Size(m)
}
func (m *ListServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesRequest proto.InternalMessageInfo

type ListServicesResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListServicesResponse) Reset()         { *m = ListServicesResponse{} }
func (m *ListServicesResponse) String() string { return proto.CompactTextString(m) }
func (*ListServicesResponse) ProtoMessage()    {}
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_search_4baa9ef27d785557, []int{2}
}
func (m *ListServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesResponse.Unmarshal(m, b)
}
func (m *ListServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesResponse.Marshal(b, m, deterministic)
}
func (dst *ListServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesResponse.Merge(dst, src)
}
func (m *ListServicesResponse) XXX_Size() int {
	return xxx_messageInfo_ListServicesResponse.Size(m)
}
func (m *ListServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesResponse proto.InternalMessageInfo

func (m *ListServicesResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "omniscience.Service")
	proto.RegisterType((*ListServicesRequest)(nil), "omniscience.ListServicesRequest")
	proto.RegisterType((*ListServicesResponse)(nil), "omniscience.ListServicesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchClient is the client API for Search service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchClient interface {
	// List all services which have documents available for searching.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
}

type searchClient struct {
	cc *grpc.ClientConn
}

func NewSearchClient(cc *grpc.ClientConn) SearchClient {
	return &searchClient{cc}
}

func (c *searchClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/omniscience.Search/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServer is the server API for Search service.
type SearchServer interface {
	// List all services which have documents available for searching.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
}

func RegisterSearchServer(s *grpc.Server, srv SearchServer) {
	s.RegisterService(&_Search_serviceDesc, srv)
}

func _Search_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omniscience.Search/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Search_serviceDesc = grpc.ServiceDesc{
	ServiceName: "omniscience.Search",
	HandlerType: (*SearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _Search_ListServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_search_4baa9ef27d785557) }

var fileDescriptor_search_4baa9ef27d785557 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xcf, 0xcd, 0xcb, 0x2c, 0x4e,
	0xce, 0x4c, 0xcd, 0x4b, 0x4e, 0x95, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c,
	0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x55,
	0x92, 0xe4, 0x62, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x51, 0x12, 0xe5, 0x12, 0xf6, 0xc9,
	0x2c, 0x2e, 0x81, 0x4a, 0x17, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x79, 0x70, 0x89,
	0xa0, 0x0a, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x19, 0x70, 0x71, 0x14, 0x43, 0xc5, 0x24,
	0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x44, 0xf4, 0x90, 0xdc, 0xa1, 0x07, 0xd5, 0x10, 0x04, 0x57,
	0x65, 0xd4, 0xc4, 0xc8, 0xc5, 0x16, 0x0c, 0x76, 0xb7, 0x50, 0x05, 0x17, 0x0f, 0xb2, 0xa1, 0x42,
	0x0a, 0x28, 0x5a, 0xb1, 0x38, 0x43, 0x4a, 0x11, 0x8f, 0x0a, 0x88, 0x8b, 0x94, 0x94, 0x9b, 0x2e,
	0x3f, 0x99, 0xcc, 0x24, 0xab, 0x24, 0xa1, 0x5f, 0x66, 0xa8, 0x8f, 0xa4, 0x5a, 0x1f, 0xe6, 0x02,
	0x2b, 0x46, 0xad, 0x24, 0x36, 0x70, 0x38, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x6a,
	0xd4, 0x8b, 0x42, 0x01, 0x00, 0x00,
}
