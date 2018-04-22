// Code generated by protoc-gen-go. DO NOT EDIT.
// source: archiver.proto

package practical_grpc_v1

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

type ZipRequest struct {
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	Contents []byte `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (m *ZipRequest) Reset()                    { *m = ZipRequest{} }
func (m *ZipRequest) String() string            { return proto.CompactTextString(m) }
func (*ZipRequest) ProtoMessage()               {}
func (*ZipRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ZipRequest) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *ZipRequest) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

type ZipResponse struct {
	ZippedContents []byte `protobuf:"bytes,1,opt,name=zipped_contents,json=zippedContents,proto3" json:"zipped_contents,omitempty"`
}

func (m *ZipResponse) Reset()                    { *m = ZipResponse{} }
func (m *ZipResponse) String() string            { return proto.CompactTextString(m) }
func (*ZipResponse) ProtoMessage()               {}
func (*ZipResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ZipResponse) GetZippedContents() []byte {
	if m != nil {
		return m.ZippedContents
	}
	return nil
}

func init() {
	proto.RegisterType((*ZipRequest)(nil), "practical.grpc.v1.ZipRequest")
	proto.RegisterType((*ZipResponse)(nil), "practical.grpc.v1.ZipResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Archiver service

type ArchiverClient interface {
	Zip(ctx context.Context, opts ...grpc.CallOption) (Archiver_ZipClient, error)
}

type archiverClient struct {
	cc *grpc.ClientConn
}

func NewArchiverClient(cc *grpc.ClientConn) ArchiverClient {
	return &archiverClient{cc}
}

func (c *archiverClient) Zip(ctx context.Context, opts ...grpc.CallOption) (Archiver_ZipClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Archiver_serviceDesc.Streams[0], c.cc, "/practical.grpc.v1.Archiver/Zip", opts...)
	if err != nil {
		return nil, err
	}
	x := &archiverZipClient{stream}
	return x, nil
}

type Archiver_ZipClient interface {
	Send(*ZipRequest) error
	CloseAndRecv() (*ZipResponse, error)
	grpc.ClientStream
}

type archiverZipClient struct {
	grpc.ClientStream
}

func (x *archiverZipClient) Send(m *ZipRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *archiverZipClient) CloseAndRecv() (*ZipResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ZipResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Archiver service

type ArchiverServer interface {
	Zip(Archiver_ZipServer) error
}

func RegisterArchiverServer(s *grpc.Server, srv ArchiverServer) {
	s.RegisterService(&_Archiver_serviceDesc, srv)
}

func _Archiver_Zip_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ArchiverServer).Zip(&archiverZipServer{stream})
}

type Archiver_ZipServer interface {
	SendAndClose(*ZipResponse) error
	Recv() (*ZipRequest, error)
	grpc.ServerStream
}

type archiverZipServer struct {
	grpc.ServerStream
}

func (x *archiverZipServer) SendAndClose(m *ZipResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *archiverZipServer) Recv() (*ZipRequest, error) {
	m := new(ZipRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Archiver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "practical.grpc.v1.Archiver",
	HandlerType: (*ArchiverServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Zip",
			Handler:       _Archiver_Zip_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "archiver.proto",
}

func init() { proto.RegisterFile("archiver.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x2c, 0x4a, 0xce,
	0xc8, 0x2c, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x28, 0x4a, 0x4c,
	0x2e, 0xc9, 0x4c, 0x4e, 0xcc, 0xd1, 0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x33, 0x54, 0x72, 0xe5,
	0xe2, 0x8a, 0xca, 0x2c, 0x08, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe6, 0xe2, 0x4c,
	0xcb, 0xcc, 0x49, 0x8d, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2,
	0x00, 0x09, 0xf8, 0x25, 0xe6, 0xa6, 0x0a, 0x49, 0x71, 0x71, 0x24, 0xe7, 0xe7, 0x95, 0xa4, 0xe6,
	0x95, 0x14, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xf9, 0x4a, 0x66, 0x5c, 0xdc, 0x60,
	0x63, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xd4, 0xb9, 0xf8, 0xab, 0x32, 0x0b, 0x0a, 0x52,
	0x53, 0xe2, 0xe1, 0x3a, 0x18, 0xc1, 0x3a, 0xf8, 0x20, 0xc2, 0xce, 0x50, 0x51, 0xa3, 0x20, 0x2e,
	0x0e, 0x47, 0xa8, 0x1b, 0x85, 0xdc, 0xb8, 0x98, 0xa3, 0x32, 0x0b, 0x84, 0x64, 0xf5, 0x30, 0x5c,
	0xa9, 0x87, 0x70, 0xa2, 0x94, 0x1c, 0x2e, 0x69, 0x88, 0xd5, 0x1a, 0x8c, 0x49, 0x6c, 0x60, 0xcf,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xa2, 0xc8, 0xb9, 0xfe, 0x00, 0x00, 0x00,
}
