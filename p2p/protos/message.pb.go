// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package protos

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

type Node struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	IP                   string   `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=Port,proto3" json:"Port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_4c999e2d4f999698, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Node) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Node) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Ping struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_4c999e2d4f999698, []int{1}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (dst *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(dst, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

type Pong struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Ping                 []byte   `protobuf:"bytes,2,opt,name=Ping,proto3" json:"Ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_4c999e2d4f999698, []int{2}
}
func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (dst *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(dst, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Pong) GetPing() []byte {
	if m != nil {
		return m.Ping
	}
	return nil
}

type FindNode struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Target               []byte   `protobuf:"bytes,2,opt,name=Target,proto3" json:"Target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindNode) Reset()         { *m = FindNode{} }
func (m *FindNode) String() string { return proto.CompactTextString(m) }
func (*FindNode) ProtoMessage()    {}
func (*FindNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_4c999e2d4f999698, []int{3}
}
func (m *FindNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNode.Unmarshal(m, b)
}
func (m *FindNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNode.Marshal(b, m, deterministic)
}
func (dst *FindNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNode.Merge(dst, src)
}
func (m *FindNode) XXX_Size() int {
	return xxx_messageInfo_FindNode.Size(m)
}
func (m *FindNode) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNode.DiscardUnknown(m)
}

var xxx_messageInfo_FindNode proto.InternalMessageInfo

func (m *FindNode) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *FindNode) GetTarget() []byte {
	if m != nil {
		return m.Target
	}
	return nil
}

type Neighbors struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Nodes                []*Node  `protobuf:"bytes,2,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Neighbors) Reset()         { *m = Neighbors{} }
func (m *Neighbors) String() string { return proto.CompactTextString(m) }
func (*Neighbors) ProtoMessage()    {}
func (*Neighbors) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_4c999e2d4f999698, []int{4}
}
func (m *Neighbors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbors.Unmarshal(m, b)
}
func (m *Neighbors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbors.Marshal(b, m, deterministic)
}
func (dst *Neighbors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbors.Merge(dst, src)
}
func (m *Neighbors) XXX_Size() int {
	return xxx_messageInfo_Neighbors.Size(m)
}
func (m *Neighbors) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbors.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbors proto.InternalMessageInfo

func (m *Neighbors) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Neighbors) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "protos.Node")
	proto.RegisterType((*Ping)(nil), "protos.Ping")
	proto.RegisterType((*Pong)(nil), "protos.Pong")
	proto.RegisterType((*FindNode)(nil), "protos.FindNode")
	proto.RegisterType((*Neighbors)(nil), "protos.Neighbors")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_4c999e2d4f999698) }

var fileDescriptor_message_4c999e2d4f999698 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x56,
	0x5c, 0x2c, 0x7e, 0xf9, 0x29, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x3c, 0x41, 0x4c, 0x9e, 0x2e, 0x60, 0x7e, 0x80, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x93, 0x67, 0x80, 0x90, 0x10, 0x17, 0x4b, 0x40, 0x7e, 0x51, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0x6f, 0x10, 0x98, 0xad, 0x24, 0xc6, 0xc5, 0x12, 0x90, 0x99, 0x97, 0x8e, 0xae, 0x57, 0x49, 0x0b,
	0xa4, 0x16, 0x53, 0x1c, 0x6c, 0x46, 0x66, 0x5e, 0x3a, 0xd8, 0x54, 0x9e, 0x20, 0x30, 0x5b, 0xc9,
	0x88, 0x8b, 0xc3, 0x2d, 0x33, 0x2f, 0x05, 0xab, 0x1b, 0xc4, 0xb8, 0xd8, 0x42, 0x12, 0x8b, 0xd2,
	0x53, 0x4b, 0xa0, 0x3a, 0xa0, 0x3c, 0x25, 0x7b, 0x2e, 0x4e, 0xbf, 0xd4, 0xcc, 0xf4, 0x8c, 0xa4,
	0xfc, 0xa2, 0x62, 0x0c, 0x4d, 0x4a, 0x5c, 0xac, 0x20, 0xc3, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35,
	0xb8, 0x8d, 0x78, 0x20, 0xfe, 0x2d, 0xd6, 0x03, 0x09, 0x06, 0x41, 0xa4, 0x92, 0x20, 0x9e, 0x37,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x44, 0x78, 0xd9, 0x14, 0x01, 0x00, 0x00,
}