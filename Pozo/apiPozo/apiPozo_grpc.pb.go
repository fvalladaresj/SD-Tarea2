// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiPozo

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

// DataNodePozoClient is the client API for DataNodePozo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNodePozoClient interface {
	EscribirJugada(ctx context.Context, in *JugadaJugador, opts ...grpc.CallOption) (*Signal, error)
	RetornarJugadas(ctx context.Context, in *JugadorYEtapa, opts ...grpc.CallOption) (*JugadasArchivo, error)
}

type dataNodePozoClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNodePozoClient(cc grpc.ClientConnInterface) DataNodePozoClient {
	return &dataNodePozoClient{cc}
}

func (c *dataNodePozoClient) EscribirJugada(ctx context.Context, in *JugadaJugador, opts ...grpc.CallOption) (*Signal, error) {
	out := new(Signal)
	err := c.cc.Invoke(ctx, "/apiPozo.DataNodePozo/EscribirJugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataNodePozoClient) RetornarJugadas(ctx context.Context, in *JugadorYEtapa, opts ...grpc.CallOption) (*JugadasArchivo, error) {
	out := new(JugadasArchivo)
	err := c.cc.Invoke(ctx, "/apiPozo.DataNodePozo/RetornarJugadas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNodePozoServer is the server API for DataNodePozo service.
// All implementations must embed UnimplementedDataNodePozoServer
// for forward compatibility
type DataNodePozoServer interface {
	EscribirJugada(context.Context, *JugadaJugador) (*Signal, error)
	RetornarJugadas(context.Context, *JugadorYEtapa) (*JugadasArchivo, error)
	mustEmbedUnimplementedDataNodePozoServer()
}

// UnimplementedDataNodePozoServer must be embedded to have forward compatible implementations.
type UnimplementedDataNodePozoServer struct {
}

func (UnimplementedDataNodePozoServer) EscribirJugada(context.Context, *JugadaJugador) (*Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscribirJugada not implemented")
}
func (UnimplementedDataNodePozoServer) RetornarJugadas(context.Context, *JugadorYEtapa) (*JugadasArchivo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetornarJugadas not implemented")
}
func (UnimplementedDataNodePozoServer) mustEmbedUnimplementedDataNodePozoServer() {}

// UnsafeDataNodePozoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNodePozoServer will
// result in compilation errors.
type UnsafeDataNodePozoServer interface {
	mustEmbedUnimplementedDataNodePozoServer()
}

func RegisterDataNodePozoServer(s grpc.ServiceRegistrar, srv DataNodePozoServer) {
	s.RegisterService(&DataNodePozo_ServiceDesc, srv)
}

func _DataNodePozo_EscribirJugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JugadaJugador)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodePozoServer).EscribirJugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiPozo.DataNodePozo/EscribirJugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodePozoServer).EscribirJugada(ctx, req.(*JugadaJugador))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataNodePozo_RetornarJugadas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JugadorYEtapa)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNodePozoServer).RetornarJugadas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiPozo.DataNodePozo/RetornarJugadas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNodePozoServer).RetornarJugadas(ctx, req.(*JugadorYEtapa))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNodePozo_ServiceDesc is the grpc.ServiceDesc for DataNodePozo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNodePozo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiPozo.DataNodePozo",
	HandlerType: (*DataNodePozoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EscribirJugada",
			Handler:    _DataNodePozo_EscribirJugada_Handler,
		},
		{
			MethodName: "RetornarJugadas",
			Handler:    _DataNodePozo_RetornarJugadas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Pozo/apiPozo/apiPozo.proto",
}