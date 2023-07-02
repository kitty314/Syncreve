// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: file_sync.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FileSyncService_AddDownloadTask_FullMethodName           = "/FileSyncService/AddDownloadTask"
	FileSyncService_AddDownloadTasksByDirPath_FullMethodName = "/FileSyncService/AddDownloadTasksByDirPath"
	FileSyncService_CancelDownloadTask_FullMethodName        = "/FileSyncService/CancelDownloadTask"
	FileSyncService_GetDownloadInfoStream_FullMethodName     = "/FileSyncService/GetDownloadInfoStream"
	FileSyncService_GetDownloadInfo_FullMethodName           = "/FileSyncService/GetDownloadInfo"
	FileSyncService_GetDownloadCountInfo_FullMethodName      = "/FileSyncService/GetDownloadCountInfo"
	FileSyncService_GetDownloadCountStream_FullMethodName    = "/FileSyncService/GetDownloadCountStream"
)

// FileSyncServiceClient is the client API for FileSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileSyncServiceClient interface {
	AddDownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (*DownloadTaskResult, error)
	AddDownloadTasksByDirPath(ctx context.Context, in *DownloadDirTaskRequest, opts ...grpc.CallOption) (*DownloadTaskResult, error)
	CancelDownloadTask(ctx context.Context, in *DownloadTaskCancelRequest, opts ...grpc.CallOption) (*DownloadTaskCancelResult, error)
	GetDownloadInfoStream(ctx context.Context, in *DownloadInfoRequest, opts ...grpc.CallOption) (FileSyncService_GetDownloadInfoStreamClient, error)
	GetDownloadInfo(ctx context.Context, in *DownloadInfoRequest, opts ...grpc.CallOption) (*DownLoadInfoResult, error)
	GetDownloadCountInfo(ctx context.Context, in *DownloadCountRequest, opts ...grpc.CallOption) (*DownloadCountResult, error)
	GetDownloadCountStream(ctx context.Context, in *DownloadCountRequest, opts ...grpc.CallOption) (FileSyncService_GetDownloadCountStreamClient, error)
}

type fileSyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileSyncServiceClient(cc grpc.ClientConnInterface) FileSyncServiceClient {
	return &fileSyncServiceClient{cc}
}

func (c *fileSyncServiceClient) AddDownloadTask(ctx context.Context, in *DownloadTaskRequest, opts ...grpc.CallOption) (*DownloadTaskResult, error) {
	out := new(DownloadTaskResult)
	err := c.cc.Invoke(ctx, FileSyncService_AddDownloadTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) AddDownloadTasksByDirPath(ctx context.Context, in *DownloadDirTaskRequest, opts ...grpc.CallOption) (*DownloadTaskResult, error) {
	out := new(DownloadTaskResult)
	err := c.cc.Invoke(ctx, FileSyncService_AddDownloadTasksByDirPath_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) CancelDownloadTask(ctx context.Context, in *DownloadTaskCancelRequest, opts ...grpc.CallOption) (*DownloadTaskCancelResult, error) {
	out := new(DownloadTaskCancelResult)
	err := c.cc.Invoke(ctx, FileSyncService_CancelDownloadTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) GetDownloadInfoStream(ctx context.Context, in *DownloadInfoRequest, opts ...grpc.CallOption) (FileSyncService_GetDownloadInfoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSyncService_ServiceDesc.Streams[0], FileSyncService_GetDownloadInfoStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSyncServiceGetDownloadInfoStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSyncService_GetDownloadInfoStreamClient interface {
	Recv() (*DownLoadInfoResult, error)
	grpc.ClientStream
}

type fileSyncServiceGetDownloadInfoStreamClient struct {
	grpc.ClientStream
}

func (x *fileSyncServiceGetDownloadInfoStreamClient) Recv() (*DownLoadInfoResult, error) {
	m := new(DownLoadInfoResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSyncServiceClient) GetDownloadInfo(ctx context.Context, in *DownloadInfoRequest, opts ...grpc.CallOption) (*DownLoadInfoResult, error) {
	out := new(DownLoadInfoResult)
	err := c.cc.Invoke(ctx, FileSyncService_GetDownloadInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) GetDownloadCountInfo(ctx context.Context, in *DownloadCountRequest, opts ...grpc.CallOption) (*DownloadCountResult, error) {
	out := new(DownloadCountResult)
	err := c.cc.Invoke(ctx, FileSyncService_GetDownloadCountInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSyncServiceClient) GetDownloadCountStream(ctx context.Context, in *DownloadCountRequest, opts ...grpc.CallOption) (FileSyncService_GetDownloadCountStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileSyncService_ServiceDesc.Streams[1], FileSyncService_GetDownloadCountStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSyncServiceGetDownloadCountStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSyncService_GetDownloadCountStreamClient interface {
	Recv() (*DownloadCountResult, error)
	grpc.ClientStream
}

type fileSyncServiceGetDownloadCountStreamClient struct {
	grpc.ClientStream
}

func (x *fileSyncServiceGetDownloadCountStreamClient) Recv() (*DownloadCountResult, error) {
	m := new(DownloadCountResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileSyncServiceServer is the server API for FileSyncService service.
// All implementations must embed UnimplementedFileSyncServiceServer
// for forward compatibility
type FileSyncServiceServer interface {
	AddDownloadTask(context.Context, *DownloadTaskRequest) (*DownloadTaskResult, error)
	AddDownloadTasksByDirPath(context.Context, *DownloadDirTaskRequest) (*DownloadTaskResult, error)
	CancelDownloadTask(context.Context, *DownloadTaskCancelRequest) (*DownloadTaskCancelResult, error)
	GetDownloadInfoStream(*DownloadInfoRequest, FileSyncService_GetDownloadInfoStreamServer) error
	GetDownloadInfo(context.Context, *DownloadInfoRequest) (*DownLoadInfoResult, error)
	GetDownloadCountInfo(context.Context, *DownloadCountRequest) (*DownloadCountResult, error)
	GetDownloadCountStream(*DownloadCountRequest, FileSyncService_GetDownloadCountStreamServer) error
	mustEmbedUnimplementedFileSyncServiceServer()
}

// UnimplementedFileSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileSyncServiceServer struct {
}

func (UnimplementedFileSyncServiceServer) AddDownloadTask(context.Context, *DownloadTaskRequest) (*DownloadTaskResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDownloadTask not implemented")
}
func (UnimplementedFileSyncServiceServer) AddDownloadTasksByDirPath(context.Context, *DownloadDirTaskRequest) (*DownloadTaskResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDownloadTasksByDirPath not implemented")
}
func (UnimplementedFileSyncServiceServer) CancelDownloadTask(context.Context, *DownloadTaskCancelRequest) (*DownloadTaskCancelResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelDownloadTask not implemented")
}
func (UnimplementedFileSyncServiceServer) GetDownloadInfoStream(*DownloadInfoRequest, FileSyncService_GetDownloadInfoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDownloadInfoStream not implemented")
}
func (UnimplementedFileSyncServiceServer) GetDownloadInfo(context.Context, *DownloadInfoRequest) (*DownLoadInfoResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadInfo not implemented")
}
func (UnimplementedFileSyncServiceServer) GetDownloadCountInfo(context.Context, *DownloadCountRequest) (*DownloadCountResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadCountInfo not implemented")
}
func (UnimplementedFileSyncServiceServer) GetDownloadCountStream(*DownloadCountRequest, FileSyncService_GetDownloadCountStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDownloadCountStream not implemented")
}
func (UnimplementedFileSyncServiceServer) mustEmbedUnimplementedFileSyncServiceServer() {}

// UnsafeFileSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileSyncServiceServer will
// result in compilation errors.
type UnsafeFileSyncServiceServer interface {
	mustEmbedUnimplementedFileSyncServiceServer()
}

func RegisterFileSyncServiceServer(s grpc.ServiceRegistrar, srv FileSyncServiceServer) {
	s.RegisterService(&FileSyncService_ServiceDesc, srv)
}

func _FileSyncService_AddDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).AddDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSyncService_AddDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).AddDownloadTask(ctx, req.(*DownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_AddDownloadTasksByDirPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadDirTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).AddDownloadTasksByDirPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSyncService_AddDownloadTasksByDirPath_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).AddDownloadTasksByDirPath(ctx, req.(*DownloadDirTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_CancelDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadTaskCancelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).CancelDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSyncService_CancelDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).CancelDownloadTask(ctx, req.(*DownloadTaskCancelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_GetDownloadInfoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSyncServiceServer).GetDownloadInfoStream(m, &fileSyncServiceGetDownloadInfoStreamServer{stream})
}

type FileSyncService_GetDownloadInfoStreamServer interface {
	Send(*DownLoadInfoResult) error
	grpc.ServerStream
}

type fileSyncServiceGetDownloadInfoStreamServer struct {
	grpc.ServerStream
}

func (x *fileSyncServiceGetDownloadInfoStreamServer) Send(m *DownLoadInfoResult) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSyncService_GetDownloadInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).GetDownloadInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSyncService_GetDownloadInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).GetDownloadInfo(ctx, req.(*DownloadInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_GetDownloadCountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSyncServiceServer).GetDownloadCountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSyncService_GetDownloadCountInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSyncServiceServer).GetDownloadCountInfo(ctx, req.(*DownloadCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSyncService_GetDownloadCountStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadCountRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSyncServiceServer).GetDownloadCountStream(m, &fileSyncServiceGetDownloadCountStreamServer{stream})
}

type FileSyncService_GetDownloadCountStreamServer interface {
	Send(*DownloadCountResult) error
	grpc.ServerStream
}

type fileSyncServiceGetDownloadCountStreamServer struct {
	grpc.ServerStream
}

func (x *fileSyncServiceGetDownloadCountStreamServer) Send(m *DownloadCountResult) error {
	return x.ServerStream.SendMsg(m)
}

// FileSyncService_ServiceDesc is the grpc.ServiceDesc for FileSyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileSyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FileSyncService",
	HandlerType: (*FileSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDownloadTask",
			Handler:    _FileSyncService_AddDownloadTask_Handler,
		},
		{
			MethodName: "AddDownloadTasksByDirPath",
			Handler:    _FileSyncService_AddDownloadTasksByDirPath_Handler,
		},
		{
			MethodName: "CancelDownloadTask",
			Handler:    _FileSyncService_CancelDownloadTask_Handler,
		},
		{
			MethodName: "GetDownloadInfo",
			Handler:    _FileSyncService_GetDownloadInfo_Handler,
		},
		{
			MethodName: "GetDownloadCountInfo",
			Handler:    _FileSyncService_GetDownloadCountInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDownloadInfoStream",
			Handler:       _FileSyncService_GetDownloadInfoStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetDownloadCountStream",
			Handler:       _FileSyncService_GetDownloadCountStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file_sync.proto",
}