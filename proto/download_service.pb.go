// Code generated by protoc-gen-go. DO NOT EDIT.
// source: download_service.proto

package download

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// DownloadRequest is the request type of the download.
type DownloadRequest struct {
	// File key to download from S3
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The bucket to download file from
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_download_service_9db8895138dec550, []int{0}
}
func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (dst *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(dst, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DownloadRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

// DownloadResponse is the response type of the download.
type DownloadResponse struct {
	// Raw File bytes
	File                 []byte   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadResponse) Reset()         { *m = DownloadResponse{} }
func (m *DownloadResponse) String() string { return proto.CompactTextString(m) }
func (*DownloadResponse) ProtoMessage()    {}
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_download_service_9db8895138dec550, []int{1}
}
func (m *DownloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadResponse.Unmarshal(m, b)
}
func (m *DownloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadResponse.Marshal(b, m, deterministic)
}
func (dst *DownloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadResponse.Merge(dst, src)
}
func (m *DownloadResponse) XXX_Size() int {
	return xxx_messageInfo_DownloadResponse.Size(m)
}
func (m *DownloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadResponse proto.InternalMessageInfo

func (m *DownloadResponse) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*DownloadRequest)(nil), "download.DownloadRequest")
	proto.RegisterType((*DownloadResponse)(nil), "download.DownloadResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DownloadClient is the client API for Download service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DownloadClient interface {
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Download_DownloadClient, error)
}

type downloadClient struct {
	cc *grpc.ClientConn
}

func NewDownloadClient(cc *grpc.ClientConn) DownloadClient {
	return &downloadClient{cc}
}

func (c *downloadClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Download_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Download_serviceDesc.Streams[0], "/download.Download/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &downloadDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Download_DownloadClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type downloadDownloadClient struct {
	grpc.ClientStream
}

func (x *downloadDownloadClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DownloadServer is the server API for Download service.
type DownloadServer interface {
	Download(*DownloadRequest, Download_DownloadServer) error
}

func RegisterDownloadServer(s *grpc.Server, srv DownloadServer) {
	s.RegisterService(&_Download_serviceDesc, srv)
}

func _Download_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DownloadServer).Download(m, &downloadDownloadServer{stream})
}

type Download_DownloadServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type downloadDownloadServer struct {
	grpc.ServerStream
}

func (x *downloadDownloadServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Download_serviceDesc = grpc.ServiceDesc{
	ServiceName: "download.Download",
	HandlerType: (*DownloadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _Download_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "download_service.proto",
}

func init() {
	proto.RegisterFile("download_service.proto", fileDescriptor_download_service_9db8895138dec550)
}

var fileDescriptor_download_service_9db8895138dec550 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xc9, 0x2f, 0xcf,
	0xcb, 0xc9, 0x4f, 0x4c, 0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x89, 0x2b, 0x59, 0x73, 0xf1, 0xbb, 0x40, 0xd9, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x20, 0xa6, 0x90, 0x18, 0x17, 0x5b, 0x52, 0x69, 0x72, 0x76, 0x6a, 0x89, 0x04, 0x13,
	0x58, 0x10, 0xca, 0x53, 0x52, 0xe3, 0x12, 0x40, 0x68, 0x2e, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x12, 0xe2, 0x62, 0x49, 0xcb, 0xcc, 0x49, 0x05, 0x6b, 0xe7, 0x09, 0x02, 0xb3, 0x8d, 0x02, 0xb9,
	0x38, 0x60, 0xea, 0x84, 0x5c, 0x91, 0xd8, 0x92, 0x7a, 0x30, 0x77, 0xe8, 0xa1, 0x39, 0x42, 0x4a,
	0x0a, 0x9b, 0x14, 0xc4, 0x0a, 0x25, 0x06, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0x47, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xce, 0x84, 0x74, 0xdb, 0xe2, 0x00, 0x00, 0x00,
}
