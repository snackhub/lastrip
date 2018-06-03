// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Version              int32    `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Command              []byte   `protobuf:"bytes,2,opt,name=Command,proto3" json:"Command,omitempty"`
	Parameters           [][]byte `protobuf:"bytes,3,rep,name=Parameters,proto3" json:"Parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_5cf33893d400ab53, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Request) GetCommand() []byte {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Request) GetParameters() [][]byte {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type Response struct {
	Version              int32    `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_5cf33893d400ab53, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Response) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_protocol_5cf33893d400ab53) }

var fileDescriptor_protocol_5cf33893d400ab53 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x03, 0x33, 0x94, 0x62, 0xb9, 0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0xc3, 0x52, 0x8b, 0x8a, 0x33, 0xf3, 0xf3, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x83, 0x60, 0x5c, 0x90, 0x8c, 0x73, 0x7e, 0x6e, 0x6e, 0x62, 0x5e, 0x8a, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x8c, 0x2b, 0x24, 0xc7, 0xc5, 0x15, 0x90, 0x58, 0x94, 0x98,
	0x9b, 0x5a, 0x92, 0x5a, 0x54, 0x2c, 0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x13, 0x84, 0x24, 0xa2, 0x64,
	0xc7, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x4a, 0xc8, 0xfc, 0xbc, 0x92, 0xd4,
	0xbc, 0x12, 0x84, 0xf9, 0x60, 0x6e, 0x12, 0x1b, 0xd8, 0x95, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4a, 0x17, 0xdb, 0xf2, 0xb7, 0x00, 0x00, 0x00,
}
