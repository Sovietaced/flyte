// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/waitable.proto

package plugins

import (
	fmt "fmt"
	math "math"

	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
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

// Represents an Execution that was launched and could be waited on.
type Waitable struct {
	WfExecId             *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=wf_exec_id,json=wfExecId,proto3" json:"wf_exec_id,omitempty"`
	Phase                core.WorkflowExecution_Phase      `protobuf:"varint,2,opt,name=phase,proto3,enum=flyteidl.core.WorkflowExecution_Phase" json:"phase,omitempty"`
	WorkflowId           string                            `protobuf:"bytes,3,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Waitable) Reset()         { *m = Waitable{} }
func (m *Waitable) String() string { return proto.CompactTextString(m) }
func (*Waitable) ProtoMessage()    {}
func (*Waitable) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff9364c45e52921, []int{0}
}

func (m *Waitable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Waitable.Unmarshal(m, b)
}
func (m *Waitable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Waitable.Marshal(b, m, deterministic)
}
func (m *Waitable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Waitable.Merge(m, src)
}
func (m *Waitable) XXX_Size() int {
	return xxx_messageInfo_Waitable.Size(m)
}
func (m *Waitable) XXX_DiscardUnknown() {
	xxx_messageInfo_Waitable.DiscardUnknown(m)
}

var xxx_messageInfo_Waitable proto.InternalMessageInfo

func (m *Waitable) GetWfExecId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.WfExecId
	}
	return nil
}

func (m *Waitable) GetPhase() core.WorkflowExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.WorkflowExecution_UNDEFINED
}

func (m *Waitable) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func init() {
	proto.RegisterType((*Waitable)(nil), "flyteidl.plugins.Waitable")
}

func init() { proto.RegisterFile("flyteidl/plugins/waitable.proto", fileDescriptor_6ff9364c45e52921) }

var fileDescriptor_6ff9364c45e52921 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcb, 0xa9, 0x2c,
	0x49, 0xcd, 0x4c, 0xc9, 0xd1, 0x2f, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x2b, 0xd6, 0x2f, 0x4f, 0xcc,
	0x2c, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x29, 0xd0,
	0x83, 0x2a, 0x90, 0x92, 0x85, 0x6b, 0x49, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xad, 0x48, 0x4d, 0x2e,
	0x2d, 0xc9, 0xcc, 0xcf, 0x83, 0x68, 0x90, 0x92, 0x43, 0x95, 0xce, 0x4c, 0x49, 0xcd, 0x2b, 0xc9,
	0x4c, 0xcb, 0x4c, 0x2d, 0x82, 0xc8, 0x2b, 0x6d, 0x66, 0xe4, 0xe2, 0x08, 0x87, 0xda, 0x21, 0xe4,
	0xc1, 0xc5, 0x55, 0x9e, 0x16, 0x0f, 0x32, 0x22, 0x3e, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x4b, 0x0f, 0x6e, 0x25, 0xc8, 0x04, 0xbd, 0xf0, 0xfc, 0xa2, 0xec, 0xb4, 0x9c, 0xfc,
	0x72, 0x57, 0x98, 0x45, 0x9e, 0x70, 0x23, 0x83, 0x38, 0xca, 0xd3, 0x40, 0xc2, 0x9e, 0x29, 0x42,
	0x36, 0x5c, 0xac, 0x05, 0x19, 0x89, 0xc5, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x6a,
	0x84, 0x0c, 0xd1, 0x0b, 0x00, 0xa9, 0x0e, 0x82, 0x68, 0x12, 0x92, 0xe7, 0xe2, 0x2e, 0x87, 0xaa,
	0x00, 0x39, 0x84, 0x59, 0x81, 0x51, 0x83, 0x33, 0x88, 0x0b, 0x26, 0xe4, 0x99, 0xe2, 0x64, 0x1f,
	0x65, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x36, 0x3b, 0xbf,
	0x28, 0x1d, 0xc2, 0xd0, 0x87, 0xfb, 0x38, 0x3d, 0x35, 0x4f, 0xbf, 0x20, 0x49, 0x37, 0x3d, 0x5f,
	0x1f, 0x3d, 0x58, 0x93, 0xd8, 0xc0, 0xbe, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x78,
	0xb5, 0x11, 0x71, 0x01, 0x00, 0x00,
}
