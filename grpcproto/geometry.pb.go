// Code generated by protoc-gen-go. DO NOT EDIT.
// source: geometry.proto

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

// import "schema.tl.crc32.proto";
// import "schema.tl.sync.proto";
// option go_package = ".;proto";
type Point struct {
	Latitude             string   `protobuf:"bytes,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            string   `protobuf:"bytes,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_geometry_25cbdbf5c23ff554, []int{0}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (dst *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(dst, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

func (m *Point) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

type PolyLine struct {
	Points               []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolyLine) Reset()         { *m = PolyLine{} }
func (m *PolyLine) String() string { return proto.CompactTextString(m) }
func (*PolyLine) ProtoMessage()    {}
func (*PolyLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_geometry_25cbdbf5c23ff554, []int{1}
}
func (m *PolyLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolyLine.Unmarshal(m, b)
}
func (m *PolyLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolyLine.Marshal(b, m, deterministic)
}
func (dst *PolyLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolyLine.Merge(dst, src)
}
func (m *PolyLine) XXX_Size() int {
	return xxx_messageInfo_PolyLine.Size(m)
}
func (m *PolyLine) XXX_DiscardUnknown() {
	xxx_messageInfo_PolyLine.DiscardUnknown(m)
}

var xxx_messageInfo_PolyLine proto.InternalMessageInfo

func (m *PolyLine) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type Polygon struct {
	Points               []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Polygon) Reset()         { *m = Polygon{} }
func (m *Polygon) String() string { return proto.CompactTextString(m) }
func (*Polygon) ProtoMessage()    {}
func (*Polygon) Descriptor() ([]byte, []int) {
	return fileDescriptor_geometry_25cbdbf5c23ff554, []int{2}
}
func (m *Polygon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Polygon.Unmarshal(m, b)
}
func (m *Polygon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Polygon.Marshal(b, m, deterministic)
}
func (dst *Polygon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Polygon.Merge(dst, src)
}
func (m *Polygon) XXX_Size() int {
	return xxx_messageInfo_Polygon.Size(m)
}
func (m *Polygon) XXX_DiscardUnknown() {
	xxx_messageInfo_Polygon.DiscardUnknown(m)
}

var xxx_messageInfo_Polygon proto.InternalMessageInfo

func (m *Polygon) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type JsonResponse struct {
	JsonResponse         []byte   `protobuf:"bytes,1,opt,name=jsonResponse,proto3" json:"jsonResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JsonResponse) Reset()         { *m = JsonResponse{} }
func (m *JsonResponse) String() string { return proto.CompactTextString(m) }
func (*JsonResponse) ProtoMessage()    {}
func (*JsonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_geometry_25cbdbf5c23ff554, []int{3}
}
func (m *JsonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JsonResponse.Unmarshal(m, b)
}
func (m *JsonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JsonResponse.Marshal(b, m, deterministic)
}
func (dst *JsonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JsonResponse.Merge(dst, src)
}
func (m *JsonResponse) XXX_Size() int {
	return xxx_messageInfo_JsonResponse.Size(m)
}
func (m *JsonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JsonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JsonResponse proto.InternalMessageInfo

func (m *JsonResponse) GetJsonResponse() []byte {
	if m != nil {
		return m.JsonResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*Point)(nil), "grpcproto.point")
	proto.RegisterType((*PolyLine)(nil), "grpcproto.polyLine")
	proto.RegisterType((*Polygon)(nil), "grpcproto.polygon")
	proto.RegisterType((*JsonResponse)(nil), "grpcproto.jsonResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeometryClient is the client API for Geometry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeometryClient interface {
	GetCurrentAddress(ctx context.Context, in *Point, opts ...grpc.CallOption) (*JsonResponse, error)
}

type geometryClient struct {
	cc *grpc.ClientConn
}

func NewGeometryClient(cc *grpc.ClientConn) GeometryClient {
	return &geometryClient{cc}
}

func (c *geometryClient) GetCurrentAddress(ctx context.Context, in *Point, opts ...grpc.CallOption) (*JsonResponse, error) {
	out := new(JsonResponse)
	err := c.cc.Invoke(ctx, "/grpcproto.geometry/GetCurrentAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeometryServer is the server API for Geometry service.
type GeometryServer interface {
	GetCurrentAddress(context.Context, *Point) (*JsonResponse, error)
}

func RegisterGeometryServer(s *grpc.Server, srv GeometryServer) {
	s.RegisterService(&_Geometry_serviceDesc, srv)
}

func _Geometry_GetCurrentAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Point)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeometryServer).GetCurrentAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcproto.geometry/GetCurrentAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeometryServer).GetCurrentAddress(ctx, req.(*Point))
	}
	return interceptor(ctx, in, info, handler)
}

var _Geometry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcproto.geometry",
	HandlerType: (*GeometryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentAddress",
			Handler:    _Geometry_GetCurrentAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geometry.proto",
}

func init() { proto.RegisterFile("geometry.proto", fileDescriptor_geometry_25cbdbf5c23ff554) }

var fileDescriptor_geometry_25cbdbf5c23ff554 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4f, 0xcd, 0xcf,
	0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x2f, 0x2a, 0x48,
	0x06, 0x33, 0x95, 0x1c, 0xb9, 0x58, 0x0b, 0xf2, 0x33, 0xf3, 0x4a, 0x84, 0xa4, 0xb8, 0x38, 0x72,
	0x12, 0x4b, 0x32, 0x4b, 0x4a, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c,
	0x21, 0x19, 0x2e, 0xce, 0x9c, 0xfc, 0xbc, 0x74, 0x88, 0x24, 0x13, 0x58, 0x12, 0x21, 0xa0, 0x64,
	0xc2, 0xc5, 0x51, 0x90, 0x9f, 0x53, 0xe9, 0x93, 0x99, 0x97, 0x2a, 0xa4, 0xc1, 0xc5, 0x06, 0x36,
	0xae, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x40, 0x0f, 0x6e, 0x95, 0x1e, 0x58, 0x22,
	0x08, 0x2a, 0xaf, 0x64, 0xcc, 0xc5, 0x0e, 0xd2, 0x95, 0x9e, 0x9f, 0x47, 0x82, 0x26, 0x23, 0x2e,
	0x9e, 0xac, 0xe2, 0xfc, 0xbc, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x25, 0x54,
	0x3e, 0xd8, 0xe1, 0x3c, 0x41, 0x28, 0x62, 0x46, 0x5e, 0x5c, 0x1c, 0x30, 0xef, 0x0b, 0xd9, 0x71,
	0x09, 0xba, 0xa7, 0x96, 0x38, 0x97, 0x16, 0x15, 0xa5, 0xe6, 0x95, 0x38, 0xa6, 0xa4, 0x14, 0xa5,
	0x16, 0x17, 0x0b, 0x61, 0x58, 0x27, 0x25, 0x8e, 0x24, 0x82, 0x6c, 0x56, 0x12, 0x1b, 0x58, 0xcc,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x13, 0x8f, 0x76, 0x21, 0x51, 0x01, 0x00, 0x00,
}
