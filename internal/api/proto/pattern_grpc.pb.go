// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/pattern.proto

package gRPC

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DataService_Authenticate_FullMethodName = "/gRPC.DataService/Authenticate"
	DataService_StartServer_FullMethodName  = "/gRPC.DataService/StartServer"
	DataService_StopData_FullMethodName     = "/gRPC.DataService/StopData"
)

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	StartServer(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (DataService_StartServerClient, error)
	StopData(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DataService_Authenticate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) StartServer(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (DataService_StartServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[0], DataService_StartServer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceStartServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_StartServerClient interface {
	Recv() (*DataResponse, error)
	grpc.ClientStream
}

type dataServiceStartServerClient struct {
	grpc.ClientStream
}

func (x *dataServiceStartServerClient) Recv() (*DataResponse, error) {
	m := new(DataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) StopData(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DataService_StopData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	Authenticate(context.Context, *AuthRequest) (*empty.Empty, error)
	StartServer(*DataRequest, DataService_StartServerServer) error
	StopData(context.Context, *StopRequest) (*empty.Empty, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) Authenticate(context.Context, *AuthRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedDataServiceServer) StartServer(*DataRequest, DataService_StartServerServer) error {
	return status.Errorf(codes.Unimplemented, "method StartServer not implemented")
}
func (UnimplementedDataServiceServer) StopData(context.Context, *StopRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopData not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_StartServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).StartServer(m, &dataServiceStartServerServer{stream})
}

type DataService_StartServerServer interface {
	Send(*DataResponse) error
	grpc.ServerStream
}

type dataServiceStartServerServer struct {
	grpc.ServerStream
}

func (x *dataServiceStartServerServer) Send(m *DataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DataService_StopData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).StopData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_StopData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).StopData(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gRPC.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _DataService_Authenticate_Handler,
		},
		{
			MethodName: "StopData",
			Handler:    _DataService_StopData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartServer",
			Handler:       _DataService_StartServer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/pattern.proto",
}
