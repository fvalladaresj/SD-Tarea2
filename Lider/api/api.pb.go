// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: Lider/api/api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PeticionParticipar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Participar string `protobuf:"bytes,1,opt,name=Participar,proto3" json:"Participar,omitempty"`
}

func (x *PeticionParticipar) Reset() {
	*x = PeticionParticipar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeticionParticipar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeticionParticipar) ProtoMessage() {}

func (x *PeticionParticipar) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeticionParticipar.ProtoReflect.Descriptor instead.
func (*PeticionParticipar) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *PeticionParticipar) GetParticipar() string {
	if x != nil {
		return x.Participar
	}
	return ""
}

type ConfirmacionParticipacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confirmacion string `protobuf:"bytes,1,opt,name=Confirmacion,proto3" json:"Confirmacion,omitempty"`
}

func (x *ConfirmacionParticipacion) Reset() {
	*x = ConfirmacionParticipacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmacionParticipacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmacionParticipacion) ProtoMessage() {}

func (x *ConfirmacionParticipacion) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmacionParticipacion.ProtoReflect.Descriptor instead.
func (*ConfirmacionParticipacion) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *ConfirmacionParticipacion) GetConfirmacion() string {
	if x != nil {
		return x.Confirmacion
	}
	return ""
}

type Jugadas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plays []int32 `protobuf:"varint,1,rep,packed,name=Plays,proto3" json:"Plays,omitempty"`
	Etapa int32   `protobuf:"varint,2,opt,name=Etapa,proto3" json:"Etapa,omitempty"`
}

func (x *Jugadas) Reset() {
	*x = Jugadas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jugadas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jugadas) ProtoMessage() {}

func (x *Jugadas) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jugadas.ProtoReflect.Descriptor instead.
func (*Jugadas) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *Jugadas) GetPlays() []int32 {
	if x != nil {
		return x.Plays
	}
	return nil
}

func (x *Jugadas) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

type EstadoJugador struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado      []int32 `protobuf:"varint,1,rep,packed,name=Estado,proto3" json:"Estado,omitempty"`
	Ronda       int32   `protobuf:"varint,2,opt,name=Ronda,proto3" json:"Ronda,omitempty"`
	JugadorGano int32   `protobuf:"varint,3,opt,name=JugadorGano,proto3" json:"JugadorGano,omitempty"`
}

func (x *EstadoJugador) Reset() {
	*x = EstadoJugador{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstadoJugador) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoJugador) ProtoMessage() {}

func (x *EstadoJugador) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoJugador.ProtoReflect.Descriptor instead.
func (*EstadoJugador) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *EstadoJugador) GetEstado() []int32 {
	if x != nil {
		return x.Estado
	}
	return nil
}

func (x *EstadoJugador) GetRonda() int32 {
	if x != nil {
		return x.Ronda
	}
	return 0
}

func (x *EstadoJugador) GetJugadorGano() int32 {
	if x != nil {
		return x.JugadorGano
	}
	return 0
}

type PedirMonto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Solicitud string `protobuf:"bytes,1,opt,name=Solicitud,proto3" json:"Solicitud,omitempty"`
}

func (x *PedirMonto) Reset() {
	*x = PedirMonto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PedirMonto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PedirMonto) ProtoMessage() {}

func (x *PedirMonto) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PedirMonto.ProtoReflect.Descriptor instead.
func (*PedirMonto) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *PedirMonto) GetSolicitud() string {
	if x != nil {
		return x.Solicitud
	}
	return ""
}

type MontoJugador struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Monto int32 `protobuf:"varint,1,opt,name=Monto,proto3" json:"Monto,omitempty"`
}

func (x *MontoJugador) Reset() {
	*x = MontoJugador{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MontoJugador) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MontoJugador) ProtoMessage() {}

func (x *MontoJugador) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MontoJugador.ProtoReflect.Descriptor instead.
func (*MontoJugador) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *MontoJugador) GetMonto() int32 {
	if x != nil {
		return x.Monto
	}
	return 0
}

type Check struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sign bool `protobuf:"varint,1,opt,name=Sign,proto3" json:"Sign,omitempty"`
}

func (x *Check) Reset() {
	*x = Check{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Check) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Check) ProtoMessage() {}

func (x *Check) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Check.ProtoReflect.Descriptor instead.
func (*Check) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *Check) GetSign() bool {
	if x != nil {
		return x.Sign
	}
	return false
}

type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etapa int32 `protobuf:"varint,1,opt,name=Etapa,proto3" json:"Etapa,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{7}
}

func (x *State) GetEtapa() int32 {
	if x != nil {
		return x.Etapa
	}
	return 0
}

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sign bool `protobuf:"varint,1,opt,name=Sign,proto3" json:"Sign,omitempty"`
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{8}
}

func (x *Signal) GetSign() bool {
	if x != nil {
		return x.Sign
	}
	return false
}

type CantidadJugadores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CJugadores int32 `protobuf:"varint,1,opt,name=CJugadores,proto3" json:"CJugadores,omitempty"`
}

func (x *CantidadJugadores) Reset() {
	*x = CantidadJugadores{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Lider_api_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CantidadJugadores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CantidadJugadores) ProtoMessage() {}

func (x *CantidadJugadores) ProtoReflect() protoreflect.Message {
	mi := &file_Lider_api_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CantidadJugadores.ProtoReflect.Descriptor instead.
func (*CantidadJugadores) Descriptor() ([]byte, []int) {
	return file_Lider_api_api_proto_rawDescGZIP(), []int{9}
}

func (x *CantidadJugadores) GetCJugadores() int32 {
	if x != nil {
		return x.CJugadores
	}
	return 0
}

var File_Lider_api_api_proto protoreflect.FileDescriptor

var file_Lider_api_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x4c, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x34, 0x0a, 0x12, 0x50, 0x65,
	0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x72,
	0x22, 0x3f, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f,
	0x6e, 0x22, 0x35, 0x0a, 0x07, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x6c, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x50, 0x6c, 0x61,
	0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x74, 0x61, 0x70, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x45, 0x74, 0x61, 0x70, 0x61, 0x22, 0x5f, 0x0a, 0x0d, 0x45, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x6f, 0x6e, 0x64, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x52, 0x6f, 0x6e, 0x64, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x75, 0x67, 0x61, 0x64,
	0x6f, 0x72, 0x47, 0x61, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4a, 0x75,
	0x67, 0x61, 0x64, 0x6f, 0x72, 0x47, 0x61, 0x6e, 0x6f, 0x22, 0x2a, 0x0a, 0x0a, 0x50, 0x65, 0x64,
	0x69, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x75, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x75, 0x64, 0x22, 0x24, 0x0a, 0x0c, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x4a, 0x75,
	0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x05, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x22, 0x1d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x74, 0x61, 0x70, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x45, 0x74, 0x61, 0x70, 0x61, 0x22, 0x1c, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x53, 0x69, 0x67, 0x6e, 0x22, 0x33, 0x0a, 0x11, 0x43, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61,
	0x64, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x4a,
	0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x43, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x32, 0xbd, 0x02, 0x0a, 0x05, 0x4c,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x72, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x65,
	0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x72,
	0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x2b, 0x0a, 0x05, 0x4a, 0x75, 0x67, 0x61, 0x72, 0x12, 0x0c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x73, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50,
	0x65, 0x64, 0x69, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x6f, 0x6e, 0x74, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x0c, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x45, 0x74, 0x61, 0x70, 0x61, 0x73, 0x12, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x63,
	0x69, 0x61, 0x72, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x10, 0x43, 0x75, 0x61, 0x6e, 0x74, 0x6f, 0x73, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x74, 0x69, 0x64, 0x61, 0x64, 0x4a,
	0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Lider_api_api_proto_rawDescOnce sync.Once
	file_Lider_api_api_proto_rawDescData = file_Lider_api_api_proto_rawDesc
)

func file_Lider_api_api_proto_rawDescGZIP() []byte {
	file_Lider_api_api_proto_rawDescOnce.Do(func() {
		file_Lider_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_Lider_api_api_proto_rawDescData)
	})
	return file_Lider_api_api_proto_rawDescData
}

var file_Lider_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_Lider_api_api_proto_goTypes = []interface{}{
	(*PeticionParticipar)(nil),        // 0: api.PeticionParticipar
	(*ConfirmacionParticipacion)(nil), // 1: api.ConfirmacionParticipacion
	(*Jugadas)(nil),                   // 2: api.Jugadas
	(*EstadoJugador)(nil),             // 3: api.EstadoJugador
	(*PedirMonto)(nil),                // 4: api.PedirMonto
	(*MontoJugador)(nil),              // 5: api.MontoJugador
	(*Check)(nil),                     // 6: api.Check
	(*State)(nil),                     // 7: api.State
	(*Signal)(nil),                    // 8: api.Signal
	(*CantidadJugadores)(nil),         // 9: api.CantidadJugadores
}
var file_Lider_api_api_proto_depIdxs = []int32{
	0, // 0: api.Lider.ParticiparJuego:input_type -> api.PeticionParticipar
	2, // 1: api.Lider.Jugar:input_type -> api.Jugadas
	4, // 2: api.Lider.Monto:input_type -> api.PedirMonto
	6, // 3: api.Lider.EstadoEtapas:input_type -> api.Check
	8, // 4: api.Lider.Iniciar:input_type -> api.Signal
	8, // 5: api.Lider.CuantosJugadores:input_type -> api.Signal
	1, // 6: api.Lider.ParticiparJuego:output_type -> api.ConfirmacionParticipacion
	3, // 7: api.Lider.Jugar:output_type -> api.EstadoJugador
	5, // 8: api.Lider.Monto:output_type -> api.MontoJugador
	7, // 9: api.Lider.EstadoEtapas:output_type -> api.State
	8, // 10: api.Lider.Iniciar:output_type -> api.Signal
	9, // 11: api.Lider.CuantosJugadores:output_type -> api.CantidadJugadores
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Lider_api_api_proto_init() }
func file_Lider_api_api_proto_init() {
	if File_Lider_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Lider_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeticionParticipar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmacionParticipacion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jugadas); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstadoJugador); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PedirMonto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MontoJugador); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Check); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Lider_api_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CantidadJugadores); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Lider_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Lider_api_api_proto_goTypes,
		DependencyIndexes: file_Lider_api_api_proto_depIdxs,
		MessageInfos:      file_Lider_api_api_proto_msgTypes,
	}.Build()
	File_Lider_api_api_proto = out.File
	file_Lider_api_api_proto_rawDesc = nil
	file_Lider_api_api_proto_goTypes = nil
	file_Lider_api_api_proto_depIdxs = nil
}
