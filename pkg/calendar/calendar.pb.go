// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calendar.proto

package calendar

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
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6b, 0xc2, 0x40,
	0x14, 0x24, 0xd5, 0x68, 0x7c, 0x01, 0x0b, 0x8f, 0xb6, 0x2c, 0xb9, 0x34, 0x0d, 0x14, 0x3c, 0x45,
	0xb0, 0x87, 0x7a, 0x2a, 0x58, 0x2d, 0x45, 0xe8, 0xa5, 0xc1, 0x5c, 0x7a, 0x8b, 0xe6, 0x55, 0x02,
	0x6a, 0xd2, 0xec, 0xda, 0x7f, 0xd4, 0x3f, 0xd3, 0x5f, 0x55, 0xb2, 0x9b, 0x4f, 0x51, 0x90, 0xde,
	0x76, 0xf6, 0xcd, 0xe4, 0xcd, 0x4c, 0x12, 0xe8, 0xaf, 0x82, 0x0d, 0xed, 0xc2, 0x20, 0x75, 0x93,
	0x34, 0x16, 0x31, 0x1a, 0x05, 0xb6, 0x6e, 0xd7, 0x71, 0xbc, 0xde, 0xd0, 0x50, 0xde, 0x2f, 0xf7,
	0x9f, 0x43, 0x11, 0x6d, 0x89, 0x8b, 0x60, 0x9b, 0x28, 0xaa, 0xf3, 0xa3, 0x81, 0xfe, 0xf2, 0x4d,
	0x3b, 0x81, 0x08, 0x6d, 0xdf, 0x9f, 0xcf, 0x98, 0x66, 0x6b, 0x83, 0xb6, 0x27, 0xcf, 0x78, 0x05,
	0xfa, 0x22, 0x12, 0x1b, 0x62, 0x17, 0xb6, 0x36, 0xe8, 0x79, 0x0a, 0xa0, 0x0d, 0xe6, 0x8c, 0xf8,
	0x2a, 0x8d, 0x12, 0x11, 0xc5, 0x3b, 0xd6, 0x92, 0xb3, 0xfa, 0x15, 0x8e, 0xa1, 0x37, 0x4d, 0x29,
	0x10, 0x14, 0x4e, 0x04, 0x6b, 0xdb, 0xda, 0xc0, 0x1c, 0x59, 0xae, 0xb2, 0xe2, 0x16, 0x56, 0xdc,
	0x45, 0x61, 0xc5, 0xab, 0xc8, 0x78, 0x03, 0x1d, 0x9f, 0x53, 0x3a, 0x0f, 0x99, 0x2e, 0x7d, 0xe4,
	0xc8, 0x79, 0x84, 0xbe, 0x22, 0x49, 0xb3, 0x1e, 0x7d, 0xe1, 0x3d, 0xe8, 0x94, 0x9d, 0xa5, 0x61,
	0x73, 0x74, 0xe9, 0x96, 0x25, 0x28, 0x8a, 0x9a, 0x3a, 0x4f, 0x07, 0x42, 0xde, 0x08, 0xda, 0xca,
	0x83, 0x32, 0xe8, 0xae, 0x94, 0x07, 0x19, 0xd5, 0xf0, 0x0a, 0xe8, 0xdc, 0x81, 0xf9, 0x4a, 0xa2,
	0xdc, 0x7a, 0x44, 0xec, 0xbc, 0xd5, 0x29, 0xfc, 0x4c, 0x63, 0x59, 0xd2, 0x35, 0x89, 0x6a, 0x63,
	0x8e, 0xb2, 0xa4, 0x7e, 0x12, 0xfe, 0x23, 0xe9, 0xfb, 0x81, 0xf0, 0x6c, 0x27, 0x0c, 0xba, 0x7b,
	0x29, 0x2c, 0xc3, 0xe7, 0x70, 0xf4, 0xab, 0x81, 0x31, 0xcd, 0x35, 0x38, 0x01, 0xb3, 0xd6, 0x24,
	0xb2, 0xea, 0x69, 0xcd, 0x37, 0x63, 0x9d, 0x9a, 0x70, 0x1c, 0x83, 0x51, 0x34, 0x85, 0xd7, 0x15,
	0xab, 0x56, 0xb0, 0x75, 0xf4, 0x9a, 0x67, 0xcb, 0x6b, 0xe1, 0xea, 0xcb, 0x9b, 0x65, 0x59, 0xa7,
	0x26, 0xfc, 0x19, 0x3e, 0xca, 0xff, 0x62, 0xd9, 0x91, 0x5f, 0xe1, 0xc3, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3c, 0x2e, 0x27, 0xf8, 0x3a, 0x03, 0x00, 0x00,
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
