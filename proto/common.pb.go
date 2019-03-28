//@generated
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/protobuf/common.proto

package proto // import "github.com/yuwenhui/gin-starter/proto"

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

type TaskItem struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title"`
	Summary              string   `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskItem) Reset()         { *m = TaskItem{} }
func (m *TaskItem) String() string { return proto.CompactTextString(m) }
func (*TaskItem) ProtoMessage()    {}
func (*TaskItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_common_82dac53e88236dbf, []int{0}
}
func (m *TaskItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskItem.Unmarshal(m, b)
}
func (m *TaskItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskItem.Marshal(b, m, deterministic)
}
func (dst *TaskItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskItem.Merge(dst, src)
}
func (m *TaskItem) XXX_Size() int {
	return xxx_messageInfo_TaskItem.Size(m)
}
func (m *TaskItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskItem.DiscardUnknown(m)
}

var xxx_messageInfo_TaskItem proto.InternalMessageInfo

func (m *TaskItem) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TaskItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TaskItem) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func init() {
	proto.RegisterType((*TaskItem)(nil), "common.TaskItem")
}

func init() { proto.RegisterFile("proto/protobuf/common.proto", fileDescriptor_common_82dac53e88236dbf) }

var fileDescriptor_common_82dac53e88236dbf = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x07, 0x93, 0x49, 0xa5, 0x69, 0xfa, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x7a, 0x60,
	0xbe, 0x10, 0x1b, 0x84, 0xa7, 0x14, 0xcc, 0xc5, 0x11, 0x92, 0x58, 0x9c, 0xed, 0x59, 0x92, 0x9a,
	0x2b, 0x24, 0xce, 0xc5, 0x5e, 0x5a, 0x9c, 0x5a, 0x14, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x1a, 0xc4, 0x06, 0xe2, 0x7a, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4,
	0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa5, 0xb9,
	0xb9, 0x89, 0x45, 0x95, 0x12, 0xcc, 0x60, 0x71, 0x18, 0xd7, 0x49, 0x3d, 0x4a, 0x35, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0xb2, 0xb4, 0x3c, 0x35, 0x2f, 0xa3, 0x34,
	0x53, 0x3f, 0x3d, 0x33, 0x4f, 0xb7, 0xb8, 0x24, 0xb1, 0xa8, 0x24, 0xb5, 0x08, 0xea, 0x2a, 0x36,
	0x30, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x44, 0xff, 0x9f, 0xe4, 0xab, 0x00, 0x00, 0x00,
}