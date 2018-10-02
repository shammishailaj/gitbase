// Code generated by protoc-gen-go. DO NOT EDIT.
// source: binlogservice.proto

package binlogservice // import "gopkg.in/src-d/go-vitess.v1/vt/proto/binlogservice"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import binlogdata "gopkg.in/src-d/go-vitess.v1/vt/proto/binlogdata"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UpdateStream service

type UpdateStreamClient interface {
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(ctx context.Context, in *binlogdata.StreamKeyRangeRequest, opts ...grpc.CallOption) (UpdateStream_StreamKeyRangeClient, error)
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(ctx context.Context, in *binlogdata.StreamTablesRequest, opts ...grpc.CallOption) (UpdateStream_StreamTablesClient, error)
}

type updateStreamClient struct {
	cc *grpc.ClientConn
}

func NewUpdateStreamClient(cc *grpc.ClientConn) UpdateStreamClient {
	return &updateStreamClient{cc}
}

func (c *updateStreamClient) StreamKeyRange(ctx context.Context, in *binlogdata.StreamKeyRangeRequest, opts ...grpc.CallOption) (UpdateStream_StreamKeyRangeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStream_serviceDesc.Streams[0], c.cc, "/binlogservice.UpdateStream/StreamKeyRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamStreamKeyRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStream_StreamKeyRangeClient interface {
	Recv() (*binlogdata.StreamKeyRangeResponse, error)
	grpc.ClientStream
}

type updateStreamStreamKeyRangeClient struct {
	grpc.ClientStream
}

func (x *updateStreamStreamKeyRangeClient) Recv() (*binlogdata.StreamKeyRangeResponse, error) {
	m := new(binlogdata.StreamKeyRangeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *updateStreamClient) StreamTables(ctx context.Context, in *binlogdata.StreamTablesRequest, opts ...grpc.CallOption) (UpdateStream_StreamTablesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_UpdateStream_serviceDesc.Streams[1], c.cc, "/binlogservice.UpdateStream/StreamTables", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateStreamStreamTablesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UpdateStream_StreamTablesClient interface {
	Recv() (*binlogdata.StreamTablesResponse, error)
	grpc.ClientStream
}

type updateStreamStreamTablesClient struct {
	grpc.ClientStream
}

func (x *updateStreamStreamTablesClient) Recv() (*binlogdata.StreamTablesResponse, error) {
	m := new(binlogdata.StreamTablesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for UpdateStream service

type UpdateStreamServer interface {
	// StreamKeyRange returns the binlog transactions related to
	// the specified Keyrange.
	StreamKeyRange(*binlogdata.StreamKeyRangeRequest, UpdateStream_StreamKeyRangeServer) error
	// StreamTables returns the binlog transactions related to
	// the specified Tables.
	StreamTables(*binlogdata.StreamTablesRequest, UpdateStream_StreamTablesServer) error
}

func RegisterUpdateStreamServer(s *grpc.Server, srv UpdateStreamServer) {
	s.RegisterService(&_UpdateStream_serviceDesc, srv)
}

func _UpdateStream_StreamKeyRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.StreamKeyRangeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamServer).StreamKeyRange(m, &updateStreamStreamKeyRangeServer{stream})
}

type UpdateStream_StreamKeyRangeServer interface {
	Send(*binlogdata.StreamKeyRangeResponse) error
	grpc.ServerStream
}

type updateStreamStreamKeyRangeServer struct {
	grpc.ServerStream
}

func (x *updateStreamStreamKeyRangeServer) Send(m *binlogdata.StreamKeyRangeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _UpdateStream_StreamTables_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(binlogdata.StreamTablesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdateStreamServer).StreamTables(m, &updateStreamStreamTablesServer{stream})
}

type UpdateStream_StreamTablesServer interface {
	Send(*binlogdata.StreamTablesResponse) error
	grpc.ServerStream
}

type updateStreamStreamTablesServer struct {
	grpc.ServerStream
}

func (x *updateStreamStreamTablesServer) Send(m *binlogdata.StreamTablesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _UpdateStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binlogservice.UpdateStream",
	HandlerType: (*UpdateStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamKeyRange",
			Handler:       _UpdateStream_StreamKeyRange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamTables",
			Handler:       _UpdateStream_StreamTables_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "binlogservice.proto",
}

func init() { proto.RegisterFile("binlogservice.proto", fileDescriptor_binlogservice_0e1eb8b2f97a2dc1) }

var fileDescriptor_binlogservice_0e1eb8b2f97a2dc1 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xca, 0xcc, 0xcb,
	0xc9, 0x4f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x94, 0x12, 0x80, 0x70, 0x53, 0x12, 0x4b, 0x12, 0x21, 0x0a, 0x8c, 0x0e, 0x31,
	0x72, 0xf1, 0x84, 0x16, 0xa4, 0x24, 0x96, 0xa4, 0x06, 0x97, 0x14, 0xa5, 0x26, 0xe6, 0x0a, 0x45,
	0x73, 0xf1, 0x41, 0x58, 0xde, 0xa9, 0x95, 0x41, 0x89, 0x79, 0xe9, 0xa9, 0x42, 0x8a, 0x7a, 0x48,
	0xba, 0x50, 0xe5, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0x94, 0xf0, 0x29, 0x29, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x62, 0x30, 0x60, 0x14, 0x0a, 0xe5, 0xe2, 0x81, 0xc8, 0x86, 0x24,
	0x26, 0xe5, 0xa4, 0x16, 0x0b, 0xc9, 0x63, 0xea, 0x83, 0xc8, 0xc0, 0x0c, 0x56, 0xc0, 0xad, 0x00,
	0x61, 0xac, 0x93, 0x4e, 0x94, 0x56, 0x59, 0x66, 0x49, 0x6a, 0x71, 0xb1, 0x5e, 0x66, 0xbe, 0x3e,
	0x84, 0xa5, 0x9f, 0x9e, 0xaf, 0x5f, 0x56, 0xa2, 0x0f, 0xf6, 0xa4, 0x3e, 0x4a, 0x20, 0x24, 0xb1,
	0x81, 0x05, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xf4, 0x0a, 0x9c, 0x31, 0x01, 0x00,
	0x00,
}
