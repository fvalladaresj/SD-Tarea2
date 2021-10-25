// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiNameNode

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

// NameNodeClient is the client API for NameNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NameNodeClient interface {
	EscribirJugada(ctx context.Context, in *JugadaJugador, opts ...grpc.CallOption) (*Signal, error)
	PedirJugadasJugador(ctx context.Context, in *Jugador, opts ...grpc.CallOption) (*TodasLasJugadas, error)
}

type nameNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNameNodeClient(cc grpc.ClientConnInterface) NameNodeClient {
	return &nameNodeClient{cc}
}

func (c *nameNodeClient) EscribirJugada(ctx context.Context, in *JugadaJugador, opts ...grpc.CallOption) (*Signal, error) {
	out := new(Signal)
	err := c.cc.Invoke(ctx, "/apiNameNode.NameNode/EscribirJugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameNodeClient) PedirJugadasJugador(ctx context.Context, in *Jugador, opts ...grpc.CallOption) (*TodasLasJugadas, error) {
	out := new(TodasLasJugadas)
	err := c.cc.Invoke(ctx, "/apiNameNode.NameNode/PedirJugadasJugador", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameNodeServer is the server API for NameNode service.
// All implementations must embed UnimplementedNameNodeServer
// for forward compatibility
type NameNodeServer interface {
	EscribirJugada(context.Context, *JugadaJugador) (*Signal, error)
	PedirJugadasJugador(context.Context, *Jugador) (*TodasLasJugadas, error)
	mustEmbedUnimplementedNameNodeServer()
}

// UnimplementedNameNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNameNodeServer struct {
}

func (UnimplementedNameNodeServer) EscribirJugada(context.Context, *JugadaJugador) (*Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EscribirJugada not implemented")
}
func (UnimplementedNameNodeServer) PedirJugadasJugador(context.Context, *Jugador) (*TodasLasJugadas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirJugadasJugador not implemented")
}
func (UnimplementedNameNodeServer) mustEmbedUnimplementedNameNodeServer() {}

// UnsafeNameNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NameNodeServer will
// result in compilation errors.
type UnsafeNameNodeServer interface {
	mustEmbedUnimplementedNameNodeServer()
}

func RegisterNameNodeServer(s grpc.ServiceRegistrar, srv NameNodeServer) {
	s.RegisterService(&NameNode_ServiceDesc, srv)
}

func _NameNode_EscribirJugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JugadaJugador)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServer).EscribirJugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiNameNode.NameNode/EscribirJugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServer).EscribirJugada(ctx, req.(*JugadaJugador))
	}
	return interceptor(ctx, in, info, handler)
}

func _NameNode_PedirJugadasJugador_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugador)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameNodeServer).PedirJugadasJugador(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiNameNode.NameNode/PedirJugadasJugador",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameNodeServer).PedirJugadasJugador(ctx, req.(*Jugador))
	}
	return interceptor(ctx, in, info, handler)
}

// NameNode_ServiceDesc is the grpc.ServiceDesc for NameNode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NameNode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiNameNode.NameNode",
	HandlerType: (*NameNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EscribirJugada",
			Handler:    _NameNode_EscribirJugada_Handler,
		},
		{
			MethodName: "PedirJugadasJugador",
			Handler:    _NameNode_PedirJugadasJugador_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "NameNode/apiNameNode/apiNamenode.proto",
}
