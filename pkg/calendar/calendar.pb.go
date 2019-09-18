// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package calendarpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Event struct {
	UUID                 uint64               `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UserId               uint64               `protobuf:"varint,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetUUID() uint64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Event) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateEventReq struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventReq) Reset()         { *m = CreateEventReq{} }
func (m *CreateEventReq) String() string { return proto.CompactTextString(m) }
func (*CreateEventReq) ProtoMessage()    {}
func (*CreateEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{1}
}

func (m *CreateEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventReq.Unmarshal(m, b)
}
func (m *CreateEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventReq.Marshal(b, m, deterministic)
}
func (m *CreateEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventReq.Merge(m, src)
}
func (m *CreateEventReq) XXX_Size() int {
	return xxx_messageInfo_CreateEventReq.Size(m)
}
func (m *CreateEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventReq proto.InternalMessageInfo

func (m *CreateEventReq) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type CreateEventRes struct {
	UUID                 int64    `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Created              bool     `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventRes) Reset()         { *m = CreateEventRes{} }
func (m *CreateEventRes) String() string { return proto.CompactTextString(m) }
func (*CreateEventRes) ProtoMessage()    {}
func (*CreateEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{2}
}

func (m *CreateEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRes.Unmarshal(m, b)
}
func (m *CreateEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRes.Marshal(b, m, deterministic)
}
func (m *CreateEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRes.Merge(m, src)
}
func (m *CreateEventRes) XXX_Size() int {
	return xxx_messageInfo_CreateEventRes.Size(m)
}
func (m *CreateEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRes proto.InternalMessageInfo

func (m *CreateEventRes) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

func (m *CreateEventRes) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

type GetEventReq struct {
	UUID                 int64    `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventReq) Reset()         { *m = GetEventReq{} }
func (m *GetEventReq) String() string { return proto.CompactTextString(m) }
func (*GetEventReq) ProtoMessage()    {}
func (*GetEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{3}
}

func (m *GetEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventReq.Unmarshal(m, b)
}
func (m *GetEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventReq.Marshal(b, m, deterministic)
}
func (m *GetEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventReq.Merge(m, src)
}
func (m *GetEventReq) XXX_Size() int {
	return xxx_messageInfo_GetEventReq.Size(m)
}
func (m *GetEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventReq proto.InternalMessageInfo

func (m *GetEventReq) GetUUID() int64 {
	if m != nil {
		return m.UUID
	}
	return 0
}

type GetEventRes struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Getted               bool     `protobuf:"varint,2,opt,name=getted,proto3" json:"getted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventRes) Reset()         { *m = GetEventRes{} }
func (m *GetEventRes) String() string { return proto.CompactTextString(m) }
func (*GetEventRes) ProtoMessage()    {}
func (*GetEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{4}
}

func (m *GetEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRes.Unmarshal(m, b)
}
func (m *GetEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRes.Marshal(b, m, deterministic)
}
func (m *GetEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRes.Merge(m, src)
}
func (m *GetEventRes) XXX_Size() int {
	return xxx_messageInfo_GetEventRes.Size(m)
}
func (m *GetEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRes proto.InternalMessageInfo

func (m *GetEventRes) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *GetEventRes) GetGetted() bool {
	if m != nil {
		return m.Getted
	}
	return false
}

type UpdateEventReq struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventReq) Reset()         { *m = UpdateEventReq{} }
func (m *UpdateEventReq) String() string { return proto.CompactTextString(m) }
func (*UpdateEventReq) ProtoMessage()    {}
func (*UpdateEventReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{5}
}

func (m *UpdateEventReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventReq.Unmarshal(m, b)
}
func (m *UpdateEventReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventReq.Marshal(b, m, deterministic)
}
func (m *UpdateEventReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventReq.Merge(m, src)
}
func (m *UpdateEventReq) XXX_Size() int {
	return xxx_messageInfo_UpdateEventReq.Size(m)
}
func (m *UpdateEventReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventReq proto.InternalMessageInfo

func (m *UpdateEventReq) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateEventRes struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Updated              bool     `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventRes) Reset()         { *m = UpdateEventRes{} }
func (m *UpdateEventRes) String() string { return proto.CompactTextString(m) }
func (*UpdateEventRes) ProtoMessage()    {}
func (*UpdateEventRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d25d49f056cdb2, []int{6}
}

func (m *UpdateEventRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventRes.Unmarshal(m, b)
}
func (m *UpdateEventRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventRes.Marshal(b, m, deterministic)
}
func (m *UpdateEventRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventRes.Merge(m, src)
}
func (m *UpdateEventRes) XXX_Size() int {
	return xxx_messageInfo_UpdateEventRes.Size(m)
}
func (m *UpdateEventRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventRes proto.InternalMessageInfo

func (m *UpdateEventRes) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *UpdateEventRes) GetUpdated() bool {
	if m != nil {
		return m.Updated
	}
	return false
}

func init() {
	proto.RegisterType((*Event)(nil), "calendar.Event")
	proto.RegisterType((*CreateEventReq)(nil), "calendar.CreateEventReq")
	proto.RegisterType((*CreateEventRes)(nil), "calendar.CreateEventRes")
	proto.RegisterType((*GetEventReq)(nil), "calendar.GetEventReq")
	proto.RegisterType((*GetEventRes)(nil), "calendar.GetEventRes")
	proto.RegisterType((*UpdateEventReq)(nil), "calendar.UpdateEventReq")
	proto.RegisterType((*UpdateEventRes)(nil), "calendar.UpdateEventRes")
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor_e3d25d49f056cdb2) }

var fileDescriptor_e3d25d49f056cdb2 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x6b, 0xf2, 0x40,
	0x10, 0x25, 0x9f, 0x46, 0xe3, 0xe4, 0xc3, 0xc2, 0xd0, 0x96, 0x25, 0x97, 0xa6, 0x81, 0x82, 0xa7,
	0x08, 0xf6, 0x50, 0x4f, 0x05, 0xab, 0xa5, 0x08, 0xbd, 0x34, 0x98, 0x4b, 0x6f, 0xd1, 0x4c, 0x25,
	0xa0, 0x26, 0xcd, 0xae, 0xfd, 0x47, 0xfd, 0x33, 0xfd, 0x55, 0x25, 0xbb, 0x89, 0x59, 0x45, 0x41,
	0x7a, 0xdb, 0x37, 0xf3, 0x9e, 0xf3, 0xde, 0x33, 0xd0, 0x5d, 0x44, 0x2b, 0xda, 0xc4, 0x51, 0xee,
	0x67, 0x79, 0x2a, 0x52, 0xb4, 0x2a, 0xec, 0xdc, 0x2c, 0xd3, 0x74, 0xb9, 0xa2, 0xbe, 0x9c, 0xcf,
	0xb7, 0x1f, 0x7d, 0x91, 0xac, 0x89, 0x8b, 0x68, 0x9d, 0x29, 0xaa, 0xf7, 0x6d, 0x80, 0xf9, 0xfc,
	0x45, 0x1b, 0x81, 0x08, 0xcd, 0x30, 0x9c, 0x4e, 0x98, 0xe1, 0x1a, 0xbd, 0x66, 0x20, 0xdf, 0x78,
	0x09, 0xe6, 0x2c, 0x11, 0x2b, 0x62, 0xff, 0x5c, 0xa3, 0xd7, 0x09, 0x14, 0x40, 0x17, 0xec, 0x09,
	0xf1, 0x45, 0x9e, 0x64, 0x22, 0x49, 0x37, 0xac, 0x21, 0x77, 0xfa, 0x08, 0x87, 0xd0, 0x19, 0xe7,
	0x14, 0x09, 0x8a, 0x47, 0x82, 0x35, 0x5d, 0xa3, 0x67, 0x0f, 0x1c, 0x5f, 0x59, 0xf1, 0x2b, 0x2b,
	0xfe, 0xac, 0xb2, 0x12, 0xd4, 0x64, 0xbc, 0x86, 0x56, 0xc8, 0x29, 0x9f, 0xc6, 0xcc, 0x94, 0x3e,
	0x4a, 0xe4, 0x3d, 0x40, 0x57, 0x91, 0xa4, 0xd9, 0x80, 0x3e, 0xf1, 0x0e, 0x4c, 0x2a, 0xde, 0xd2,
	0xb0, 0x3d, 0xb8, 0xf0, 0x77, 0x25, 0x28, 0x8a, 0xda, 0x7a, 0x8f, 0x07, 0x42, 0xbe, 0x17, 0xb4,
	0x51, 0x06, 0x65, 0xd0, 0x5e, 0x28, 0x0f, 0x32, 0xaa, 0x15, 0x54, 0xd0, 0xbb, 0x05, 0xfb, 0x85,
	0xc4, 0xee, 0xea, 0x11, 0xb1, 0xf7, 0xaa, 0x53, 0xf8, 0x99, 0xc6, 0x8a, 0xa4, 0x4b, 0x12, 0xf5,
	0xc5, 0x12, 0x15, 0x49, 0xc3, 0x2c, 0xfe, 0x43, 0xd2, 0xb7, 0x03, 0xe1, 0xd9, 0x4e, 0x18, 0xb4,
	0xb7, 0x52, 0xb8, 0x0b, 0x5f, 0xc2, 0xc1, 0x8f, 0x01, 0xd6, 0xb8, 0xd4, 0xe0, 0x08, 0x6c, 0xad,
	0x49, 0x64, 0xf5, 0xaf, 0xed, 0xff, 0x33, 0xce, 0xa9, 0x0d, 0xc7, 0x21, 0x58, 0x55, 0x53, 0x78,
	0x55, 0xb3, 0xb4, 0x82, 0x9d, 0xa3, 0x63, 0x5e, 0x1c, 0xd7, 0xc2, 0xe9, 0xc7, 0xf7, 0xcb, 0x72,
	0x4e, 0x6d, 0xf8, 0xd3, 0xff, 0x77, 0xa8, 0x56, 0xd9, 0x7c, 0xde, 0x92, 0xdf, 0xe1, 0xfd, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x16, 0xb8, 0x3d, 0x3c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	CreateEvent(ctx context.Context, in *CreateEventReq, opts ...grpc.CallOption) (*CreateEventRes, error)
	GetEvent(ctx context.Context, in *GetEventReq, opts ...grpc.CallOption) (*GetEventRes, error)
	UpdateEvent(ctx context.Context, in *UpdateEventReq, opts ...grpc.CallOption) (*UpdateEventRes, error)
}

type calendarClient struct {
	cc *grpc.ClientConn
}

func NewCalendarClient(cc *grpc.ClientConn) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) CreateEvent(ctx context.Context, in *CreateEventReq, opts ...grpc.CallOption) (*CreateEventRes, error) {
	out := new(CreateEventRes)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEvent(ctx context.Context, in *GetEventReq, opts ...grpc.CallOption) (*GetEventRes, error) {
	out := new(GetEventRes)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) UpdateEvent(ctx context.Context, in *UpdateEventReq, opts ...grpc.CallOption) (*UpdateEventRes, error) {
	out := new(UpdateEventRes)
	err := c.cc.Invoke(ctx, "/calendar.Calendar/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	CreateEvent(context.Context, *CreateEventReq) (*CreateEventRes, error)
	GetEvent(context.Context, *GetEventReq) (*GetEventRes, error)
	UpdateEvent(context.Context, *UpdateEventReq) (*UpdateEventRes, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) CreateEvent(ctx context.Context, req *CreateEventReq) (*CreateEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedCalendarServer) GetEvent(ctx context.Context, req *GetEventReq) (*GetEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServer) UpdateEvent(ctx context.Context, req *UpdateEventReq) (*UpdateEventRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CreateEvent(ctx, req.(*CreateEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEvent(ctx, req.(*GetEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calendar.Calendar/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).UpdateEvent(ctx, req.(*UpdateEventReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calendar.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _Calendar_CreateEvent_Handler,
		},
		{
			MethodName: "GetEvent",
			Handler:    _Calendar_GetEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Calendar_UpdateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}