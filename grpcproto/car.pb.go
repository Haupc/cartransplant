// Code generated by protoc-gen-go. DO NOT EDIT.
// source: car.proto

package grpcproto

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

type RegisterTripRequest struct {
	CarID                int64    `protobuf:"varint,1,opt,name=carID,proto3" json:"carID,omitempty"`
	BeginLeaveTime       int64    `protobuf:"varint,2,opt,name=beginLeaveTime,proto3" json:"beginLeaveTime,omitempty"`
	EndLeaveTime         int64    `protobuf:"varint,3,opt,name=endLeaveTime,proto3" json:"endLeaveTime,omitempty"`
	From                 *Point   `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Point   `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	MaxDistance          int32    `protobuf:"varint,6,opt,name=maxDistance,proto3" json:"maxDistance,omitempty"`
	FeeEachKm            int64    `protobuf:"varint,7,opt,name=feeEachKm,proto3" json:"feeEachKm,omitempty"`
	Seat                 int32    `protobuf:"varint,8,opt,name=seat,proto3" json:"seat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterTripRequest) Reset()         { *m = RegisterTripRequest{} }
func (m *RegisterTripRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterTripRequest) ProtoMessage()    {}
func (*RegisterTripRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_35f65f7f665cc550, []int{0}
}
func (m *RegisterTripRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterTripRequest.Unmarshal(m, b)
}
func (m *RegisterTripRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterTripRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterTripRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterTripRequest.Merge(dst, src)
}
func (m *RegisterTripRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterTripRequest.Size(m)
}
func (m *RegisterTripRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterTripRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterTripRequest proto.InternalMessageInfo

func (m *RegisterTripRequest) GetCarID() int64 {
	if m != nil {
		return m.CarID
	}
	return 0
}

func (m *RegisterTripRequest) GetBeginLeaveTime() int64 {
	if m != nil {
		return m.BeginLeaveTime
	}
	return 0
}

func (m *RegisterTripRequest) GetEndLeaveTime() int64 {
	if m != nil {
		return m.EndLeaveTime
	}
	return 0
}

func (m *RegisterTripRequest) GetFrom() *Point {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *RegisterTripRequest) GetTo() *Point {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *RegisterTripRequest) GetMaxDistance() int32 {
	if m != nil {
		return m.MaxDistance
	}
	return 0
}

func (m *RegisterTripRequest) GetFeeEachKm() int64 {
	if m != nil {
		return m.FeeEachKm
	}
	return 0
}

func (m *RegisterTripRequest) GetSeat() int32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

type FindTripRequest struct {
	BeginLeaveTime       int64    `protobuf:"varint,1,opt,name=beginLeaveTime,proto3" json:"beginLeaveTime,omitempty"`
	EndLeaveTime         int64    `protobuf:"varint,2,opt,name=endLeaveTime,proto3" json:"endLeaveTime,omitempty"`
	From                 *Point   `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Point   `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Option               int32    `protobuf:"varint,5,opt,name=option,proto3" json:"option,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindTripRequest) Reset()         { *m = FindTripRequest{} }
func (m *FindTripRequest) String() string { return proto.CompactTextString(m) }
func (*FindTripRequest) ProtoMessage()    {}
func (*FindTripRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_35f65f7f665cc550, []int{1}
}
func (m *FindTripRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTripRequest.Unmarshal(m, b)
}
func (m *FindTripRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTripRequest.Marshal(b, m, deterministic)
}
func (dst *FindTripRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTripRequest.Merge(dst, src)
}
func (m *FindTripRequest) XXX_Size() int {
	return xxx_messageInfo_FindTripRequest.Size(m)
}
func (m *FindTripRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTripRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindTripRequest proto.InternalMessageInfo

func (m *FindTripRequest) GetBeginLeaveTime() int64 {
	if m != nil {
		return m.BeginLeaveTime
	}
	return 0
}

func (m *FindTripRequest) GetEndLeaveTime() int64 {
	if m != nil {
		return m.EndLeaveTime
	}
	return 0
}

func (m *FindTripRequest) GetFrom() *Point {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *FindTripRequest) GetTo() *Point {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *FindTripRequest) GetOption() int32 {
	if m != nil {
		return m.Option
	}
	return 0
}

type CarObject struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LicensePlate         string   `protobuf:"bytes,2,opt,name=licensePlate,proto3" json:"licensePlate,omitempty"`
	Color                string   `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Model                string   `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarObject) Reset()         { *m = CarObject{} }
func (m *CarObject) String() string { return proto.CompactTextString(m) }
func (*CarObject) ProtoMessage()    {}
func (*CarObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_35f65f7f665cc550, []int{2}
}
func (m *CarObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarObject.Unmarshal(m, b)
}
func (m *CarObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarObject.Marshal(b, m, deterministic)
}
func (dst *CarObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarObject.Merge(dst, src)
}
func (m *CarObject) XXX_Size() int {
	return xxx_messageInfo_CarObject.Size(m)
}
func (m *CarObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CarObject.DiscardUnknown(m)
}

var xxx_messageInfo_CarObject proto.InternalMessageInfo

func (m *CarObject) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CarObject) GetLicensePlate() string {
	if m != nil {
		return m.LicensePlate
	}
	return ""
}

func (m *CarObject) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *CarObject) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

type ListCarResponse struct {
	Cars                 []*CarObject `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListCarResponse) Reset()         { *m = ListCarResponse{} }
func (m *ListCarResponse) String() string { return proto.CompactTextString(m) }
func (*ListCarResponse) ProtoMessage()    {}
func (*ListCarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_35f65f7f665cc550, []int{3}
}
func (m *ListCarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCarResponse.Unmarshal(m, b)
}
func (m *ListCarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCarResponse.Marshal(b, m, deterministic)
}
func (dst *ListCarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCarResponse.Merge(dst, src)
}
func (m *ListCarResponse) XXX_Size() int {
	return xxx_messageInfo_ListCarResponse.Size(m)
}
func (m *ListCarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCarResponse proto.InternalMessageInfo

func (m *ListCarResponse) GetCars() []*CarObject {
	if m != nil {
		return m.Cars
	}
	return nil
}

type TakeTripRequest struct {
	DriverTripID         int32    `protobuf:"varint,1,opt,name=driverTripID,proto3" json:"driverTripID,omitempty"`
	BeginLeaveTime       int64    `protobuf:"varint,2,opt,name=beginLeaveTime,proto3" json:"beginLeaveTime,omitempty"`
	EndLeaveTime         int64    `protobuf:"varint,3,opt,name=endLeaveTime,proto3" json:"endLeaveTime,omitempty"`
	From                 *Point   `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Point   `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	Seat                 int32    `protobuf:"varint,6,opt,name=seat,proto3" json:"seat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TakeTripRequest) Reset()         { *m = TakeTripRequest{} }
func (m *TakeTripRequest) String() string { return proto.CompactTextString(m) }
func (*TakeTripRequest) ProtoMessage()    {}
func (*TakeTripRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_35f65f7f665cc550, []int{4}
}
func (m *TakeTripRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TakeTripRequest.Unmarshal(m, b)
}
func (m *TakeTripRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TakeTripRequest.Marshal(b, m, deterministic)
}
func (dst *TakeTripRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TakeTripRequest.Merge(dst, src)
}
func (m *TakeTripRequest) XXX_Size() int {
	return xxx_messageInfo_TakeTripRequest.Size(m)
}
func (m *TakeTripRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TakeTripRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TakeTripRequest proto.InternalMessageInfo

func (m *TakeTripRequest) GetDriverTripID() int32 {
	if m != nil {
		return m.DriverTripID
	}
	return 0
}

func (m *TakeTripRequest) GetBeginLeaveTime() int64 {
	if m != nil {
		return m.BeginLeaveTime
	}
	return 0
}

func (m *TakeTripRequest) GetEndLeaveTime() int64 {
	if m != nil {
		return m.EndLeaveTime
	}
	return 0
}

func (m *TakeTripRequest) GetFrom() *Point {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *TakeTripRequest) GetTo() *Point {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *TakeTripRequest) GetSeat() int32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

type UserTrip struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BeginLeaveTime       int64    `protobuf:"varint,2,opt,name=beginLeaveTime,proto3" json:"beginLeaveTime,omitempty"`
	EndLeaveTime         int64    `protobuf:"varint,3,opt,name=endLeaveTime,proto3" json:"endLeaveTime,omitempty"`
	From                 *Point   `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Point   `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	State                int32    `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTrip) Reset()         { *m = UserTrip{} }
func (m *UserTrip) String() string { return proto.CompactTextString(m) }
func (*UserTrip) ProtoMessage()    {}
func (*UserTrip) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_35f65f7f665cc550, []int{5}
}
func (m *UserTrip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTrip.Unmarshal(m, b)
}
func (m *UserTrip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTrip.Marshal(b, m, deterministic)
}
func (dst *UserTrip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTrip.Merge(dst, src)
}
func (m *UserTrip) XXX_Size() int {
	return xxx_messageInfo_UserTrip.Size(m)
}
func (m *UserTrip) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTrip.DiscardUnknown(m)
}

var xxx_messageInfo_UserTrip proto.InternalMessageInfo

func (m *UserTrip) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserTrip) GetBeginLeaveTime() int64 {
	if m != nil {
		return m.BeginLeaveTime
	}
	return 0
}

func (m *UserTrip) GetEndLeaveTime() int64 {
	if m != nil {
		return m.EndLeaveTime
	}
	return 0
}

func (m *UserTrip) GetFrom() *Point {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *UserTrip) GetTo() *Point {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *UserTrip) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

type ListUserTripResponse struct {
	UserTrip             []*UserTrip `protobuf:"bytes,1,rep,name=userTrip,proto3" json:"userTrip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListUserTripResponse) Reset()         { *m = ListUserTripResponse{} }
func (m *ListUserTripResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserTripResponse) ProtoMessage()    {}
func (*ListUserTripResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_35f65f7f665cc550, []int{6}
}
func (m *ListUserTripResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserTripResponse.Unmarshal(m, b)
}
func (m *ListUserTripResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserTripResponse.Marshal(b, m, deterministic)
}
func (dst *ListUserTripResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserTripResponse.Merge(dst, src)
}
func (m *ListUserTripResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserTripResponse.Size(m)
}
func (m *ListUserTripResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserTripResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserTripResponse proto.InternalMessageInfo

func (m *ListUserTripResponse) GetUserTrip() []*UserTrip {
	if m != nil {
		return m.UserTrip
	}
	return nil
}

type DeleteCarRequest struct {
	Ids                  []int32  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCarRequest) Reset()         { *m = DeleteCarRequest{} }
func (m *DeleteCarRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCarRequest) ProtoMessage()    {}
func (*DeleteCarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_35f65f7f665cc550, []int{7}
}
func (m *DeleteCarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCarRequest.Unmarshal(m, b)
}
func (m *DeleteCarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCarRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCarRequest.Merge(dst, src)
}
func (m *DeleteCarRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCarRequest.Size(m)
}
func (m *DeleteCarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCarRequest proto.InternalMessageInfo

func (m *DeleteCarRequest) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterTripRequest)(nil), "grpcproto.registerTripRequest")
	proto.RegisterType((*FindTripRequest)(nil), "grpcproto.findTripRequest")
	proto.RegisterType((*CarObject)(nil), "grpcproto.carObject")
	proto.RegisterType((*ListCarResponse)(nil), "grpcproto.listCarResponse")
	proto.RegisterType((*TakeTripRequest)(nil), "grpcproto.takeTripRequest")
	proto.RegisterType((*UserTrip)(nil), "grpcproto.UserTrip")
	proto.RegisterType((*ListUserTripResponse)(nil), "grpcproto.listUserTripResponse")
	proto.RegisterType((*DeleteCarRequest)(nil), "grpcproto.deleteCarRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CarClient is the client API for Car service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarClient interface {
	// driver
	RegisterTrip(ctx context.Context, in *RegisterTripRequest, opts ...grpc.CallOption) (*Bool, error)
	RegisterCar(ctx context.Context, in *CarObject, opts ...grpc.CallOption) (*Bool, error)
	ListMyCar(ctx context.Context, in *Int, opts ...grpc.CallOption) (*ListCarResponse, error)
	UpdateCar(ctx context.Context, in *CarObject, opts ...grpc.CallOption) (*Bool, error)
	DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*Bool, error)
	// user
	FindTrip(ctx context.Context, in *FindTripRequest, opts ...grpc.CallOption) (*JsonResponse, error)
	TakeTrip(ctx context.Context, in *TakeTripRequest, opts ...grpc.CallOption) (*Bool, error)
	ListUserTrip(ctx context.Context, in *Int, opts ...grpc.CallOption) (*ListUserTripResponse, error)
}

type carClient struct {
	cc *grpc.ClientConn
}

func NewCarClient(cc *grpc.ClientConn) CarClient {
	return &carClient{cc}
}

func (c *carClient) RegisterTrip(ctx context.Context, in *RegisterTripRequest, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/grpcproto.car/registerTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) RegisterCar(ctx context.Context, in *CarObject, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/grpcproto.car/registerCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) ListMyCar(ctx context.Context, in *Int, opts ...grpc.CallOption) (*ListCarResponse, error) {
	out := new(ListCarResponse)
	err := c.cc.Invoke(ctx, "/grpcproto.car/listMyCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) UpdateCar(ctx context.Context, in *CarObject, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/grpcproto.car/updateCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/grpcproto.car/deleteCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) FindTrip(ctx context.Context, in *FindTripRequest, opts ...grpc.CallOption) (*JsonResponse, error) {
	out := new(JsonResponse)
	err := c.cc.Invoke(ctx, "/grpcproto.car/findTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) TakeTrip(ctx context.Context, in *TakeTripRequest, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, "/grpcproto.car/takeTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carClient) ListUserTrip(ctx context.Context, in *Int, opts ...grpc.CallOption) (*ListUserTripResponse, error) {
	out := new(ListUserTripResponse)
	err := c.cc.Invoke(ctx, "/grpcproto.car/listUserTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarServer is the server API for Car service.
type CarServer interface {
	// driver
	RegisterTrip(context.Context, *RegisterTripRequest) (*Bool, error)
	RegisterCar(context.Context, *CarObject) (*Bool, error)
	ListMyCar(context.Context, *Int) (*ListCarResponse, error)
	UpdateCar(context.Context, *CarObject) (*Bool, error)
	DeleteCar(context.Context, *DeleteCarRequest) (*Bool, error)
	// user
	FindTrip(context.Context, *FindTripRequest) (*JsonResponse, error)
	TakeTrip(context.Context, *TakeTripRequest) (*Bool, error)
	ListUserTrip(context.Context, *Int) (*ListUserTripResponse, error)
}

func RegisterCarServer(s *grpc.Server, srv CarServer) {
	s.RegisterService(&_Car_serviceDesc, srv)
}

func _Car_RegisterTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).RegisterTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/RegisterTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).RegisterTrip(ctx, req.(*RegisterTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_RegisterCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).RegisterCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/RegisterCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).RegisterCar(ctx, req.(*CarObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_ListMyCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).ListMyCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/ListMyCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).ListMyCar(ctx, req.(*Int))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_UpdateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).UpdateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/UpdateCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).UpdateCar(ctx, req.(*CarObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_DeleteCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).DeleteCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/DeleteCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).DeleteCar(ctx, req.(*DeleteCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_FindTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).FindTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/FindTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).FindTrip(ctx, req.(*FindTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_TakeTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).TakeTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/TakeTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).TakeTrip(ctx, req.(*TakeTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Car_ListUserTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).ListUserTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/ListUserTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).ListUserTrip(ctx, req.(*Int))
	}
	return interceptor(ctx, in, info, handler)
}

var _Car_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcproto.car",
	HandlerType: (*CarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerTrip",
			Handler:    _Car_RegisterTrip_Handler,
		},
		{
			MethodName: "registerCar",
			Handler:    _Car_RegisterCar_Handler,
		},
		{
			MethodName: "listMyCar",
			Handler:    _Car_ListMyCar_Handler,
		},
		{
			MethodName: "updateCar",
			Handler:    _Car_UpdateCar_Handler,
		},
		{
			MethodName: "deleteCar",
			Handler:    _Car_DeleteCar_Handler,
		},
		{
			MethodName: "findTrip",
			Handler:    _Car_FindTrip_Handler,
		},
		{
			MethodName: "takeTrip",
			Handler:    _Car_TakeTrip_Handler,
		},
		{
			MethodName: "listUserTrip",
			Handler:    _Car_ListUserTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "car.proto",
}

func init() { proto.RegisterFile("car.proto", fileDescriptor_car_35f65f7f665cc550) }

var fileDescriptor_car_35f65f7f665cc550 = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0x92, 0xa6, 0x34, 0xb7, 0x55, 0x3b, 0x79, 0x15, 0x44, 0x05, 0x41, 0x15, 0x4d, 0xa8,
	0x4f, 0x43, 0x2a, 0x48, 0x13, 0xe2, 0x61, 0x0f, 0x1b, 0x42, 0x08, 0x10, 0xc8, 0x1a, 0x1f, 0xe0,
	0x26, 0x77, 0xc5, 0x5b, 0x12, 0x07, 0xdb, 0x9b, 0xd8, 0x77, 0xf0, 0x3d, 0xbc, 0xf1, 0x1b, 0x88,
	0x5f, 0x41, 0xb1, 0xdb, 0xcc, 0x4d, 0xc7, 0xb4, 0xc7, 0xbd, 0xf9, 0xde, 0x73, 0xed, 0x9c, 0x73,
	0xef, 0xc9, 0x85, 0x28, 0x65, 0x72, 0xbf, 0x92, 0x42, 0x0b, 0x12, 0x2d, 0x65, 0x95, 0x9a, 0xe3,
	0x64, 0xb8, 0x44, 0x51, 0xa0, 0x96, 0x57, 0x16, 0x9a, 0x0c, 0x52, 0x51, 0x14, 0xa2, 0xb4, 0x51,
	0xf2, 0xd3, 0x87, 0x5d, 0x89, 0x4b, 0xae, 0x34, 0xca, 0x13, 0xc9, 0x2b, 0x8a, 0xdf, 0x2f, 0x50,
	0x69, 0x32, 0x86, 0x30, 0x65, 0xf2, 0xfd, 0x71, 0xec, 0x4d, 0xbd, 0x59, 0x40, 0x6d, 0x40, 0x9e,
	0xc3, 0x70, 0x81, 0x4b, 0x5e, 0x7e, 0x44, 0x76, 0x89, 0x27, 0xbc, 0xc0, 0xd8, 0x37, 0x70, 0x2b,
	0x4b, 0x12, 0x18, 0x60, 0x99, 0x5d, 0x57, 0x05, 0xa6, 0x6a, 0x23, 0x47, 0xf6, 0xa0, 0x73, 0x2a,
	0x45, 0x11, 0x77, 0xa6, 0xde, 0xac, 0x3f, 0xdf, 0xd9, 0x6f, 0x18, 0xef, 0x57, 0x82, 0x97, 0x9a,
	0x1a, 0x94, 0x4c, 0xc1, 0xd7, 0x22, 0x0e, 0xff, 0x53, 0xe3, 0x6b, 0x41, 0xa6, 0xd0, 0x2f, 0xd8,
	0x8f, 0x63, 0xae, 0x34, 0x2b, 0x53, 0x8c, 0xbb, 0x53, 0x6f, 0x16, 0x52, 0x37, 0x45, 0x9e, 0x40,
	0x74, 0x8a, 0xf8, 0x96, 0xa5, 0xdf, 0x3e, 0x14, 0xf1, 0x03, 0x43, 0xe5, 0x3a, 0x41, 0x08, 0x74,
	0x14, 0x32, 0x1d, 0xf7, 0xcc, 0x45, 0x73, 0x4e, 0x7e, 0x79, 0x30, 0x3a, 0xe5, 0x65, 0xe6, 0x76,
	0x64, 0x5b, 0xbb, 0x77, 0x27, 0xed, 0xfe, 0x2d, 0xda, 0x83, 0x3b, 0x68, 0xef, 0xdc, 0xa2, 0xfd,
	0x21, 0x74, 0x45, 0xa5, 0xb9, 0x28, 0x4d, 0x87, 0x42, 0xba, 0x8a, 0x92, 0x73, 0xe3, 0x85, 0xcf,
	0x8b, 0x33, 0x4c, 0x35, 0x19, 0x82, 0xcf, 0x33, 0x43, 0x36, 0xa4, 0x3e, 0xcf, 0x6a, 0x82, 0x39,
	0x4f, 0xb1, 0x54, 0xf8, 0x25, 0x67, 0xda, 0x12, 0x8c, 0xe8, 0x46, 0xce, 0x8c, 0x5f, 0xe4, 0x42,
	0x1a, 0x86, 0x11, 0xb5, 0x41, 0x9d, 0x2d, 0x44, 0x86, 0xb9, 0xe1, 0x14, 0x51, 0x1b, 0x24, 0x6f,
	0x60, 0x94, 0x73, 0xa5, 0x8f, 0x98, 0xa4, 0xa8, 0x2a, 0x51, 0x2a, 0x24, 0x33, 0xe8, 0xa4, 0x4c,
	0xaa, 0xd8, 0x9b, 0x06, 0xb3, 0xfe, 0x7c, 0xec, 0x70, 0x6f, 0x68, 0x51, 0x53, 0x91, 0xfc, 0xf1,
	0x60, 0xa4, 0xd9, 0x39, 0xba, 0x9d, 0x4e, 0x60, 0x90, 0x49, 0x7e, 0x69, 0x0d, 0xb9, 0xb2, 0x60,
	0x48, 0x37, 0x72, 0xf7, 0xd2, 0x89, 0x6b, 0x27, 0x75, 0x1d, 0x27, 0xfd, 0xf6, 0xa0, 0xf7, 0x55,
	0x59, 0xda, 0x5b, 0x93, 0xb8, 0x8f, 0x22, 0xc6, 0x10, 0x2a, 0x5d, 0xdb, 0xc2, 0xaa, 0xb0, 0x41,
	0xf2, 0x0e, 0xc6, 0xf5, 0x8c, 0xd7, 0x4a, 0x9a, 0x41, 0xbf, 0x80, 0xde, 0xc5, 0x2a, 0xb7, 0x1a,
	0xf6, 0xae, 0xf3, 0x6a, 0x53, 0xde, 0x14, 0x25, 0x7b, 0xb0, 0x93, 0x61, 0x8e, 0x1a, 0x8d, 0x5d,
	0xec, 0xbc, 0x77, 0x20, 0xe0, 0x99, 0x35, 0x4b, 0x48, 0xeb, 0xe3, 0xfc, 0x6f, 0x00, 0x41, 0xca,
	0x24, 0x39, 0x84, 0x81, 0xbb, 0x9c, 0xc8, 0x53, 0xe7, 0xf1, 0x1b, 0xb6, 0xd6, 0x64, 0xe4, 0xe0,
	0x0b, 0x21, 0x72, 0xf2, 0x0a, 0xfa, 0xeb, 0xba, 0x23, 0x26, 0xc9, 0x8d, 0x4e, 0xdc, 0xbe, 0x75,
	0x00, 0x51, 0xad, 0xf6, 0xd3, 0x55, 0x7d, 0x67, 0xe8, 0xa0, 0xbc, 0xd4, 0x93, 0x89, 0x13, 0xb7,
	0x7d, 0x3f, 0x87, 0xe8, 0xa2, 0xca, 0x98, 0x51, 0x77, 0xd7, 0x8f, 0xbd, 0x86, 0xa8, 0xe9, 0x08,
	0x79, 0xec, 0xa0, 0xed, 0x3e, 0x6d, 0x5f, 0x3d, 0x84, 0xde, 0x7a, 0x4b, 0x11, 0x97, 0x56, 0x6b,
	0x75, 0x4d, 0x1e, 0x39, 0xd8, 0x99, 0x12, 0x65, 0xc3, 0xf7, 0x00, 0x7a, 0xeb, 0x9f, 0x6f, 0xe3,
	0x81, 0xd6, 0x1f, 0x79, 0xd3, 0x97, 0x07, 0xae, 0x1f, 0xb6, 0x9a, 0xf4, 0xac, 0xd5, 0xa4, 0xb6,
	0x71, 0x16, 0x5d, 0x83, 0xbd, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0xec, 0x01, 0x6c, 0x95, 0xb4,
	0x06, 0x00, 0x00,
}
