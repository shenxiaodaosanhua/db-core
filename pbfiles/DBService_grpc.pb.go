// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbfiles

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

// DBServiceClient is the client API for DBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBServiceClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	Get(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	First(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*FirstResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
	Tx(ctx context.Context, opts ...grpc.CallOption) (DBService_TxClient, error)
}

type dBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDBServiceClient(cc grpc.ClientConnInterface) DBServiceClient {
	return &dBServiceClient{cc}
}

func (c *dBServiceClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/DBService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) Get(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/DBService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) First(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*FirstResponse, error) {
	out := new(FirstResponse)
	err := c.cc.Invoke(ctx, "/DBService/First", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := c.cc.Invoke(ctx, "/DBService/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) Tx(ctx context.Context, opts ...grpc.CallOption) (DBService_TxClient, error) {
	stream, err := c.cc.NewStream(ctx, &DBService_ServiceDesc.Streams[0], "/DBService/Tx", opts...)
	if err != nil {
		return nil, err
	}
	x := &dBServiceTxClient{stream}
	return x, nil
}

type DBService_TxClient interface {
	Send(*TxRequest) error
	Recv() (*TxResponse, error)
	grpc.ClientStream
}

type dBServiceTxClient struct {
	grpc.ClientStream
}

func (x *dBServiceTxClient) Send(m *TxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dBServiceTxClient) Recv() (*TxResponse, error) {
	m := new(TxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DBServiceServer is the server API for DBService service.
// All implementations must embed UnimplementedDBServiceServer
// for forward compatibility
type DBServiceServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Get(context.Context, *QueryRequest) (*QueryResponse, error)
	First(context.Context, *QueryRequest) (*FirstResponse, error)
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
	Tx(DBService_TxServer) error
	mustEmbedUnimplementedDBServiceServer()
}

// UnimplementedDBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDBServiceServer struct {
}

func (UnimplementedDBServiceServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedDBServiceServer) Get(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDBServiceServer) First(context.Context, *QueryRequest) (*FirstResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method First not implemented")
}
func (UnimplementedDBServiceServer) Exec(context.Context, *ExecRequest) (*ExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (UnimplementedDBServiceServer) Tx(DBService_TxServer) error {
	return status.Errorf(codes.Unimplemented, "method Tx not implemented")
}
func (UnimplementedDBServiceServer) mustEmbedUnimplementedDBServiceServer() {}

// UnsafeDBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBServiceServer will
// result in compilation errors.
type UnsafeDBServiceServer interface {
	mustEmbedUnimplementedDBServiceServer()
}

func RegisterDBServiceServer(s grpc.ServiceRegistrar, srv DBServiceServer) {
	s.RegisterService(&DBService_ServiceDesc, srv)
}

func _DBService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DBService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DBService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).Get(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_First_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).First(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DBService/First",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).First(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DBService/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_Tx_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DBServiceServer).Tx(&dBServiceTxServer{stream})
}

type DBService_TxServer interface {
	Send(*TxResponse) error
	Recv() (*TxRequest, error)
	grpc.ServerStream
}

type dBServiceTxServer struct {
	grpc.ServerStream
}

func (x *dBServiceTxServer) Send(m *TxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dBServiceTxServer) Recv() (*TxRequest, error) {
	m := new(TxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DBService_ServiceDesc is the grpc.ServiceDesc for DBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DBService",
	HandlerType: (*DBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _DBService_Query_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DBService_Get_Handler,
		},
		{
			MethodName: "First",
			Handler:    _DBService_First_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _DBService_Exec_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tx",
			Handler:       _DBService_Tx_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "DBService.proto",
}
