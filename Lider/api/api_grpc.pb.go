// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// LiderClient is the client API for Lider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiderClient interface {
	ParticiparJuego(ctx context.Context, in *PeticionParticipar, opts ...grpc.CallOption) (*ConfirmacionParticipacion, error)
	Jugar(ctx context.Context, in *Jugadas, opts ...grpc.CallOption) (*EstadoJugador, error)
	Monto(ctx context.Context, in *PedirMonto, opts ...grpc.CallOption) (*MontoJugador, error)
	EstadoEtapas(ctx context.Context, in *Check, opts ...grpc.CallOption) (*State, error)
	Iniciar(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*Signal, error)
	CuantosJugadores(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*CantidadJugadores, error)
}

type liderClient struct {
	cc grpc.ClientConnInterface
}

func NewLiderClient(cc grpc.ClientConnInterface) LiderClient {
	return &liderClient{cc}
}

func (c *liderClient) ParticiparJuego(ctx context.Context, in *PeticionParticipar, opts ...grpc.CallOption) (*ConfirmacionParticipacion, error) {
	out := new(ConfirmacionParticipacion)
	err := c.cc.Invoke(ctx, "/api.Lider/ParticiparJuego", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderClient) Jugar(ctx context.Context, in *Jugadas, opts ...grpc.CallOption) (*EstadoJugador, error) {
	out := new(EstadoJugador)
	err := c.cc.Invoke(ctx, "/api.Lider/Jugar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderClient) Monto(ctx context.Context, in *PedirMonto, opts ...grpc.CallOption) (*MontoJugador, error) {
	out := new(MontoJugador)
	err := c.cc.Invoke(ctx, "/api.Lider/Monto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderClient) EstadoEtapas(ctx context.Context, in *Check, opts ...grpc.CallOption) (*State, error) {
	out := new(State)
	err := c.cc.Invoke(ctx, "/api.Lider/EstadoEtapas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderClient) Iniciar(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*Signal, error) {
	out := new(Signal)
	err := c.cc.Invoke(ctx, "/api.Lider/Iniciar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderClient) CuantosJugadores(ctx context.Context, in *Signal, opts ...grpc.CallOption) (*CantidadJugadores, error) {
	out := new(CantidadJugadores)
	err := c.cc.Invoke(ctx, "/api.Lider/CuantosJugadores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiderServer is the server API for Lider service.
// All implementations must embed UnimplementedLiderServer
// for forward compatibility
type LiderServer interface {
	ParticiparJuego(context.Context, *PeticionParticipar) (*ConfirmacionParticipacion, error)
	Jugar(context.Context, *Jugadas) (*EstadoJugador, error)
	Monto(context.Context, *PedirMonto) (*MontoJugador, error)
	EstadoEtapas(context.Context, *Check) (*State, error)
	Iniciar(context.Context, *Signal) (*Signal, error)
	CuantosJugadores(context.Context, *Signal) (*CantidadJugadores, error)
	mustEmbedUnimplementedLiderServer()
}

// UnimplementedLiderServer must be embedded to have forward compatible implementations.
type UnimplementedLiderServer struct {
}

func (UnimplementedLiderServer) ParticiparJuego(context.Context, *PeticionParticipar) (*ConfirmacionParticipacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParticiparJuego not implemented")
}
func (UnimplementedLiderServer) Jugar(context.Context, *Jugadas) (*EstadoJugador, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Jugar not implemented")
}
func (UnimplementedLiderServer) Monto(context.Context, *PedirMonto) (*MontoJugador, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Monto not implemented")
}
func (UnimplementedLiderServer) EstadoEtapas(context.Context, *Check) (*State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstadoEtapas not implemented")
}
func (UnimplementedLiderServer) Iniciar(context.Context, *Signal) (*Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Iniciar not implemented")
}
func (UnimplementedLiderServer) CuantosJugadores(context.Context, *Signal) (*CantidadJugadores, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CuantosJugadores not implemented")
}
func (UnimplementedLiderServer) mustEmbedUnimplementedLiderServer() {}

// UnsafeLiderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiderServer will
// result in compilation errors.
type UnsafeLiderServer interface {
	mustEmbedUnimplementedLiderServer()
}

func RegisterLiderServer(s grpc.ServiceRegistrar, srv LiderServer) {
	s.RegisterService(&Lider_ServiceDesc, srv)
}

func _Lider_ParticiparJuego_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeticionParticipar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).ParticiparJuego(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Lider/ParticiparJuego",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).ParticiparJuego(ctx, req.(*PeticionParticipar))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lider_Jugar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugadas)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).Jugar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Lider/Jugar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).Jugar(ctx, req.(*Jugadas))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lider_Monto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PedirMonto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).Monto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Lider/Monto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).Monto(ctx, req.(*PedirMonto))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lider_EstadoEtapas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Check)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).EstadoEtapas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Lider/EstadoEtapas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).EstadoEtapas(ctx, req.(*Check))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lider_Iniciar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).Iniciar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Lider/Iniciar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).Iniciar(ctx, req.(*Signal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lider_CuantosJugadores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Signal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServer).CuantosJugadores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Lider/CuantosJugadores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServer).CuantosJugadores(ctx, req.(*Signal))
	}
	return interceptor(ctx, in, info, handler)
}

// Lider_ServiceDesc is the grpc.ServiceDesc for Lider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Lider",
	HandlerType: (*LiderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParticiparJuego",
			Handler:    _Lider_ParticiparJuego_Handler,
		},
		{
			MethodName: "Jugar",
			Handler:    _Lider_Jugar_Handler,
		},
		{
			MethodName: "Monto",
			Handler:    _Lider_Monto_Handler,
		},
		{
			MethodName: "EstadoEtapas",
			Handler:    _Lider_EstadoEtapas_Handler,
		},
		{
			MethodName: "Iniciar",
			Handler:    _Lider_Iniciar_Handler,
		},
		{
			MethodName: "CuantosJugadores",
			Handler:    _Lider_CuantosJugadores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Lider/api/api.proto",
}
