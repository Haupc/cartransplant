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
	return fileDescriptor_car_2d6e3600a8732949, []int{0}
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
	return fileDescriptor_car_2d6e3600a8732949, []int{1}
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
	return fileDescriptor_car_2d6e3600a8732949, []int{2}
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
	return fileDescriptor_car_2d6e3600a8732949, []int{3}
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
	return fileDescriptor_car_2d6e3600a8732949, []int{4}
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
	Id                   int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BeginLeaveTime       int64        `protobuf:"varint,2,opt,name=beginLeaveTime,proto3" json:"beginLeaveTime,omitempty"`
	EndLeaveTime         int64        `protobuf:"varint,3,opt,name=endLeaveTime,proto3" json:"endLeaveTime,omitempty"`
	From                 string       `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   string       `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	State                int32        `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	Driver               *UserProfile `protobuf:"bytes,7,opt,name=driver,proto3" json:"driver,omitempty"`
	Car                  *CarObject   `protobuf:"bytes,8,opt,name=car,proto3" json:"car,omitempty"`
	Price                int64        `protobuf:"varint,9,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserTrip) Reset()         { *m = UserTrip{} }
func (m *UserTrip) String() string { return proto.CompactTextString(m) }
func (*UserTrip) ProtoMessage()    {}
func (*UserTrip) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_2d6e3600a8732949, []int{5}
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

func (m *UserTrip) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *UserTrip) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *UserTrip) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *UserTrip) GetDriver() *UserProfile {
	if m != nil {
		return m.Driver
	}
	return nil
}

func (m *UserTrip) GetCar() *CarObject {
	if m != nil {
		return m.Car
	}
	return nil
}

func (m *UserTrip) GetPrice() int64 {
	if m != nil {
		return m.Price
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
	return fileDescriptor_car_2d6e3600a8732949, []int{6}
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
	return fileDescriptor_car_2d6e3600a8732949, []int{7}
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

type DriverTrip struct {
	UserTrips            []*UserTrip `protobuf:"bytes,1,rep,name=userTrips,proto3" json:"userTrips,omitempty"`
	BeginLeaveTime       int64       `protobuf:"varint,2,opt,name=beginLeaveTime,proto3" json:"beginLeaveTime,omitempty"`
	EndLeaveTime         int64       `protobuf:"varint,3,opt,name=endLeaveTime,proto3" json:"endLeaveTime,omitempty"`
	From                 string      `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To                   string      `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	State                int32       `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	TotalSeat            int32       `protobuf:"varint,7,opt,name=totalSeat,proto3" json:"totalSeat,omitempty"`
	ReamaingSeat         int32       `protobuf:"varint,8,opt,name=reamaingSeat,proto3" json:"reamaingSeat,omitempty"`
	Car                  *CarObject  `protobuf:"bytes,9,opt,name=car,proto3" json:"car,omitempty"`
	PriceEachKm          int32       `protobuf:"varint,10,opt,name=priceEachKm,proto3" json:"priceEachKm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DriverTrip) Reset()         { *m = DriverTrip{} }
func (m *DriverTrip) String() string { return proto.CompactTextString(m) }
func (*DriverTrip) ProtoMessage()    {}
func (*DriverTrip) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_2d6e3600a8732949, []int{8}
}
func (m *DriverTrip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DriverTrip.Unmarshal(m, b)
}
func (m *DriverTrip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DriverTrip.Marshal(b, m, deterministic)
}
func (dst *DriverTrip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DriverTrip.Merge(dst, src)
}
func (m *DriverTrip) XXX_Size() int {
	return xxx_messageInfo_DriverTrip.Size(m)
}
func (m *DriverTrip) XXX_DiscardUnknown() {
	xxx_messageInfo_DriverTrip.DiscardUnknown(m)
}

var xxx_messageInfo_DriverTrip proto.InternalMessageInfo

func (m *DriverTrip) GetUserTrips() []*UserTrip {
	if m != nil {
		return m.UserTrips
	}
	return nil
}

func (m *DriverTrip) GetBeginLeaveTime() int64 {
	if m != nil {
		return m.BeginLeaveTime
	}
	return 0
}

func (m *DriverTrip) GetEndLeaveTime() int64 {
	if m != nil {
		return m.EndLeaveTime
	}
	return 0
}

func (m *DriverTrip) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *DriverTrip) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *DriverTrip) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *DriverTrip) GetTotalSeat() int32 {
	if m != nil {
		return m.TotalSeat
	}
	return 0
}

func (m *DriverTrip) GetReamaingSeat() int32 {
	if m != nil {
		return m.ReamaingSeat
	}
	return 0
}

func (m *DriverTrip) GetCar() *CarObject {
	if m != nil {
		return m.Car
	}
	return nil
}

func (m *DriverTrip) GetPriceEachKm() int32 {
	if m != nil {
		return m.PriceEachKm
	}
	return 0
}

type ListDriverTripResponse struct {
	Trips                []*DriverTrip `protobuf:"bytes,1,rep,name=trips,proto3" json:"trips,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListDriverTripResponse) Reset()         { *m = ListDriverTripResponse{} }
func (m *ListDriverTripResponse) String() string { return proto.CompactTextString(m) }
func (*ListDriverTripResponse) ProtoMessage()    {}
func (*ListDriverTripResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_car_2d6e3600a8732949, []int{9}
}
func (m *ListDriverTripResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDriverTripResponse.Unmarshal(m, b)
}
func (m *ListDriverTripResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDriverTripResponse.Marshal(b, m, deterministic)
}
func (dst *ListDriverTripResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDriverTripResponse.Merge(dst, src)
}
func (m *ListDriverTripResponse) XXX_Size() int {
	return xxx_messageInfo_ListDriverTripResponse.Size(m)
}
func (m *ListDriverTripResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDriverTripResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDriverTripResponse proto.InternalMessageInfo

func (m *ListDriverTripResponse) GetTrips() []*DriverTrip {
	if m != nil {
		return m.Trips
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterTripRequest)(nil), "grpcproto.registerTripRequest")
	proto.RegisterType((*FindTripRequest)(nil), "grpcproto.findTripRequest")
	proto.RegisterType((*CarObject)(nil), "grpcproto.carObject")
	proto.RegisterType((*ListCarResponse)(nil), "grpcproto.listCarResponse")
	proto.RegisterType((*TakeTripRequest)(nil), "grpcproto.takeTripRequest")
	proto.RegisterType((*UserTrip)(nil), "grpcproto.userTrip")
	proto.RegisterType((*ListUserTripResponse)(nil), "grpcproto.listUserTripResponse")
	proto.RegisterType((*DeleteCarRequest)(nil), "grpcproto.deleteCarRequest")
	proto.RegisterType((*DriverTrip)(nil), "grpcproto.driverTrip")
	proto.RegisterType((*ListDriverTripResponse)(nil), "grpcproto.listDriverTripResponse")
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
	ListDriverTrip(ctx context.Context, in *Int, opts ...grpc.CallOption) (*ListDriverTripResponse, error)
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

func (c *carClient) ListDriverTrip(ctx context.Context, in *Int, opts ...grpc.CallOption) (*ListDriverTripResponse, error) {
	out := new(ListDriverTripResponse)
	err := c.cc.Invoke(ctx, "/grpcproto.car/listDriverTrip", in, out, opts...)
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
	ListDriverTrip(context.Context, *Int) (*ListDriverTripResponse, error)
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

func _Car_ListDriverTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServer).ListDriverTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.car/ListDriverTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServer).ListDriverTrip(ctx, req.(*Int))
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
			MethodName: "listDriverTrip",
			Handler:    _Car_ListDriverTrip_Handler,
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

func init() { proto.RegisterFile("car.proto", fileDescriptor_car_2d6e3600a8732949) }

var fileDescriptor_car_2d6e3600a8732949 = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xd2, 0xa6, 0x6b, 0x4e, 0xa7, 0x76, 0xf2, 0xca, 0x88, 0xca, 0x04, 0x25, 0x9a, 0xa6,
	0x4a, 0x48, 0x45, 0x14, 0xa4, 0x09, 0x71, 0xb1, 0x8b, 0x6d, 0x42, 0x08, 0x10, 0x53, 0x36, 0x1e,
	0xc0, 0x4d, 0xdc, 0xe2, 0x2d, 0x89, 0x83, 0xe3, 0x4d, 0xec, 0x25, 0xb8, 0xe1, 0x15, 0x78, 0x0d,
	0x1e, 0x82, 0x07, 0xe0, 0x5d, 0x90, 0xed, 0x34, 0x71, 0xd3, 0x6d, 0xec, 0x82, 0x8b, 0xdd, 0xf9,
	0xfc, 0xd8, 0xfe, 0xce, 0xf9, 0x3e, 0x1f, 0x83, 0x1b, 0x62, 0x3e, 0xce, 0x38, 0x13, 0x0c, 0xb9,
	0x73, 0x9e, 0x85, 0x6a, 0x39, 0xe8, 0xce, 0x09, 0x4b, 0x88, 0xe0, 0x57, 0x3a, 0x34, 0x58, 0x0f,
	0x59, 0x92, 0xb0, 0xb4, 0xb0, 0xba, 0x09, 0x11, 0x38, 0xc2, 0x02, 0x6b, 0xdb, 0xff, 0x61, 0xc3,
	0x26, 0x27, 0x73, 0x9a, 0x0b, 0xc2, 0x4f, 0x39, 0xcd, 0x02, 0xf2, 0xf5, 0x82, 0xe4, 0x02, 0xf5,
	0xc1, 0x09, 0x31, 0x7f, 0x77, 0xe8, 0x59, 0x43, 0x6b, 0xd4, 0x08, 0xb4, 0x81, 0x76, 0xa1, 0x3b,
	0x25, 0x73, 0x9a, 0x7e, 0x20, 0xf8, 0x92, 0x9c, 0xd2, 0x84, 0x78, 0xb6, 0x0a, 0xd7, 0xbc, 0xc8,
	0x87, 0x75, 0x92, 0x46, 0x55, 0x56, 0x43, 0x65, 0x2d, 0xf9, 0xd0, 0x0e, 0x34, 0x67, 0x9c, 0x25,
	0x5e, 0x73, 0x68, 0x8d, 0x3a, 0x93, 0x8d, 0x71, 0x59, 0xc1, 0x38, 0x63, 0x34, 0x15, 0x81, 0x8a,
	0xa2, 0x21, 0xd8, 0x82, 0x79, 0xce, 0x0d, 0x39, 0xb6, 0x60, 0x68, 0x08, 0x9d, 0x04, 0x7f, 0x3b,
	0xa4, 0xb9, 0xc0, 0x69, 0x48, 0xbc, 0xd6, 0xd0, 0x1a, 0x39, 0x81, 0xe9, 0x42, 0xdb, 0xe0, 0xce,
	0x08, 0x39, 0xc2, 0xe1, 0x97, 0xf7, 0x89, 0xb7, 0xa6, 0xa0, 0x54, 0x0e, 0x84, 0xa0, 0x99, 0x13,
	0x2c, 0xbc, 0xb6, 0xda, 0xa8, 0xd6, 0xfe, 0x2f, 0x0b, 0x7a, 0x33, 0x9a, 0x46, 0x66, 0x47, 0x56,
	0x6b, 0xb7, 0xee, 0x54, 0xbb, 0x7d, 0x4b, 0xed, 0x8d, 0x3b, 0xd4, 0xde, 0xbc, 0xa5, 0xf6, 0x2d,
	0x68, 0xb1, 0x4c, 0x50, 0x96, 0xaa, 0x0e, 0x39, 0x41, 0x61, 0xf9, 0xe7, 0x4a, 0x1b, 0x9f, 0xa6,
	0x67, 0x24, 0x14, 0xa8, 0x0b, 0x36, 0x8d, 0x14, 0x58, 0x27, 0xb0, 0x69, 0x24, 0x01, 0xc6, 0x34,
	0x24, 0x69, 0x4e, 0x8e, 0x63, 0x2c, 0x34, 0x40, 0x37, 0x58, 0xf2, 0x29, 0xfa, 0x59, 0xcc, 0xb8,
	0x42, 0xe8, 0x06, 0xda, 0x90, 0xde, 0x84, 0x45, 0x24, 0x56, 0x98, 0xdc, 0x40, 0x1b, 0xfe, 0x1b,
	0xe8, 0xc5, 0x34, 0x17, 0x07, 0x98, 0x07, 0x24, 0xcf, 0x58, 0x9a, 0x13, 0x34, 0x82, 0x66, 0x88,
	0x79, 0xee, 0x59, 0xc3, 0xc6, 0xa8, 0x33, 0xe9, 0x1b, 0xd8, 0x4b, 0x58, 0x81, 0xca, 0xf0, 0xff,
	0x58, 0xd0, 0x13, 0xf8, 0x9c, 0x98, 0x9d, 0xf6, 0x61, 0x3d, 0xe2, 0xf4, 0x52, 0x0b, 0xb2, 0x90,
	0xa0, 0x13, 0x2c, 0xf9, 0xee, 0xa5, 0x12, 0x17, 0x4a, 0x6a, 0x19, 0x4a, 0xfa, 0x6e, 0x43, 0xfb,
	0x22, 0xd7, 0xb0, 0x57, 0x98, 0xf8, 0x9f, 0x45, 0x20, 0xa3, 0x08, 0xb7, 0x80, 0xdc, 0x2d, 0x21,
	0xbb, 0x0a, 0x60, 0x1f, 0x9c, 0x5c, 0x48, 0xca, 0x35, 0x42, 0x6d, 0xa0, 0x31, 0xb4, 0x74, 0x6b,
	0xd5, 0xdb, 0xe8, 0x4c, 0xb6, 0x8c, 0xe2, 0x24, 0xf4, 0x63, 0xce, 0x66, 0x34, 0x26, 0x41, 0x91,
	0x85, 0x76, 0xa1, 0x11, 0x62, 0xae, 0xde, 0xcb, 0x4d, 0xdc, 0xca, 0x04, 0x79, 0x5b, 0xc6, 0x69,
	0x48, 0x3c, 0x57, 0x8f, 0x10, 0x65, 0xf8, 0x6f, 0xa1, 0x2f, 0xd5, 0xf2, 0x39, 0x5f, 0xcc, 0x9b,
	0x42, 0x32, 0xcf, 0xab, 0x3e, 0x15, 0xb2, 0xd9, 0xac, 0xe1, 0x50, 0xe9, 0x65, 0x92, 0xbf, 0x03,
	0x1b, 0x11, 0x89, 0x89, 0x20, 0x4a, 0x78, 0x5a, 0x39, 0x1b, 0xd0, 0xa0, 0x91, 0x96, 0x9d, 0x13,
	0xc8, 0xa5, 0xff, 0xdb, 0x06, 0xa8, 0x84, 0x83, 0x5e, 0x80, 0xbb, 0x38, 0x20, 0xbf, 0xed, 0x9a,
	0x2a, 0xeb, 0x9e, 0x90, 0xb4, 0x0d, 0xae, 0x60, 0x02, 0xc7, 0x27, 0x52, 0x60, 0x6b, 0x2a, 0x52,
	0x39, 0xe4, 0xdd, 0x9c, 0xe0, 0x04, 0xd3, 0x74, 0x7e, 0x52, 0xcd, 0xb2, 0x25, 0xdf, 0x82, 0x36,
	0xf7, 0x5f, 0xb4, 0x0d, 0xa1, 0xa3, 0x98, 0x2a, 0xe6, 0x25, 0xe8, 0x79, 0x6a, 0xb8, 0xfc, 0x23,
	0xd8, 0x92, 0x14, 0x1e, 0x96, 0x6d, 0x2d, 0x49, 0x7c, 0x06, 0x8e, 0x30, 0x5a, 0xfb, 0xc0, 0xb8,
	0xa5, 0x22, 0x21, 0xd0, 0x39, 0x93, 0x9f, 0x4d, 0x85, 0x08, 0xed, 0x4b, 0xf0, 0xd5, 0x0f, 0x84,
	0x1e, 0x1b, 0xbb, 0xae, 0xf9, 0x9a, 0x06, 0x3d, 0x23, 0x3e, 0x65, 0x2c, 0x46, 0xaf, 0xa0, 0xb3,
	0xc8, 0x3b, 0x90, 0xba, 0xbb, 0xae, 0xb6, 0xd5, 0x5d, 0x7b, 0xe0, 0xca, 0x2a, 0x3e, 0x5e, 0xc9,
	0x3d, 0x5d, 0x23, 0x4a, 0x53, 0x31, 0x18, 0x18, 0x76, 0x7d, 0xb8, 0x4d, 0xc0, 0xbd, 0xc8, 0x22,
	0xac, 0x84, 0x77, 0xd7, 0xcb, 0x5e, 0x83, 0x5b, 0x8a, 0x15, 0x3d, 0x32, 0xdb, 0x52, 0x93, 0xf0,
	0xea, 0xd6, 0x03, 0xe8, 0x2e, 0x77, 0x7b, 0x05, 0xec, 0xd3, 0x1a, 0xd8, 0x6b, 0x88, 0xd9, 0x87,
	0xf6, 0xe2, 0x3f, 0x43, 0x66, 0x6d, 0xb5, 0x4f, 0x6e, 0xf0, 0xd0, 0x88, 0x9d, 0xe5, 0x2c, 0x2d,
	0x0f, 0xd8, 0x83, 0xf6, 0x62, 0x4c, 0x2f, 0x1d, 0x50, 0x9b, 0xdd, 0xab, 0xf0, 0xf7, 0xe5, 0x6f,
	0x53, 0xbd, 0xf7, 0x15, 0xf0, 0x4f, 0x6a, 0xe0, 0xeb, 0x83, 0x61, 0xda, 0x52, 0xb1, 0x97, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xec, 0x1f, 0x92, 0xee, 0x08, 0x00, 0x00,
}
