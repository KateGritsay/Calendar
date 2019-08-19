package calendar

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion3

type Event struct {
	UUID                 uint64               `protobuf:"varint,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UserID               uint64               `protobuf:"varint,5,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d17a9d3f0ddf27e, []int{0}
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

func (m *Event) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "event")
}

func init() { proto.RegisterFile("calendar.proto", fileDescriptor_2d17a9d3f0ddf27e) }

var fileDescriptor_2d17a9d3f0ddf27e = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0xa1, 0x92, 0xf2, 0xe8, 0x92,
	0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x10, 0x05, 0x4a, 0x61, 0x5c, 0xac, 0xae,
	0x20, 0xc3, 0x84, 0x2c, 0xb9, 0xb8, 0x92, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x53, 0xe2, 0x13, 0x4b,
	0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xa4, 0xf4, 0x20, 0xda, 0xf5, 0x60, 0xda, 0xf5, 0x42,
	0x60, 0xda, 0x83, 0x38, 0xa1, 0xaa, 0x1d, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xd2, 0x8a, 0xf2, 0x73,
	0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x23, 0x03, 0x2e, 0x36, 0xb0, 0xb9, 0xc5,
	0x42, 0x6a, 0x5c, 0x2c, 0x01, 0xa5, 0xc5, 0x19, 0x42, 0x6c, 0x7a, 0x60, 0x01, 0x29, 0x31, 0x0c,
	0x43, 0x5d, 0x41, 0x0e, 0x4e, 0x62, 0x03, 0xf3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x35,
	0x77, 0x1d, 0x30, 0xdd, 0x00, 0x00, 0x00,
}
