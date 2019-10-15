// Code generated by protoc-gen-go. DO NOT EDIT.
// source: process.proto

// Package process provides a data model for process manager plugin template

package process

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Template struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cmd                  string            `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	POptions             *TemplatePOptions `protobuf:"bytes,3,opt,name=p_options,json=pOptions,proto3" json:"p_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Template) Reset()         { *m = Template{} }
func (m *Template) String() string { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()    {}
func (*Template) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0}
}

func (m *Template) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Template.Unmarshal(m, b)
}
func (m *Template) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Template.Marshal(b, m, deterministic)
}
func (m *Template) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Template.Merge(m, src)
}
func (m *Template) XXX_Size() int {
	return xxx_messageInfo_Template.Size(m)
}
func (m *Template) XXX_DiscardUnknown() {
	xxx_messageInfo_Template.DiscardUnknown(m)
}

var xxx_messageInfo_Template proto.InternalMessageInfo

func (m *Template) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Template) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *Template) GetPOptions() *TemplatePOptions {
	if m != nil {
		return m.POptions
	}
	return nil
}

type TemplatePOptions struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	OutWriter            bool     `protobuf:"varint,2,opt,name=out_writer,json=outWriter,proto3" json:"out_writer,omitempty"`
	ErrWriter            bool     `protobuf:"varint,3,opt,name=err_writer,json=errWriter,proto3" json:"err_writer,omitempty"`
	Restart              int32    `protobuf:"varint,4,opt,name=restart,proto3" json:"restart,omitempty"`
	Detach               bool     `protobuf:"varint,5,opt,name=detach,proto3" json:"detach,omitempty"`
	RunOnStartup         bool     `protobuf:"varint,6,opt,name=run_on_startup,json=runOnStartup,proto3" json:"run_on_startup,omitempty"`
	Notify               bool     `protobuf:"varint,7,opt,name=notify,proto3" json:"notify,omitempty"`
	AutoTerminate        bool     `protobuf:"varint,8,opt,name=auto_terminate,json=autoTerminate,proto3" json:"auto_terminate,omitempty"`
	CpuAffinity          string   `protobuf:"bytes,9,opt,name=cpu_affinity,json=cpuAffinity,proto3" json:"cpu_affinity,omitempty"`
	CpuAffinityDelay     string   `protobuf:"bytes,10,opt,name=cpu_affinity_delay,json=cpuAffinityDelay,proto3" json:"cpu_affinity_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplatePOptions) Reset()         { *m = TemplatePOptions{} }
func (m *TemplatePOptions) String() string { return proto.CompactTextString(m) }
func (*TemplatePOptions) ProtoMessage()    {}
func (*TemplatePOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0}
}

func (m *TemplatePOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplatePOptions.Unmarshal(m, b)
}
func (m *TemplatePOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplatePOptions.Marshal(b, m, deterministic)
}
func (m *TemplatePOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplatePOptions.Merge(m, src)
}
func (m *TemplatePOptions) XXX_Size() int {
	return xxx_messageInfo_TemplatePOptions.Size(m)
}
func (m *TemplatePOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplatePOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TemplatePOptions proto.InternalMessageInfo

func (m *TemplatePOptions) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *TemplatePOptions) GetOutWriter() bool {
	if m != nil {
		return m.OutWriter
	}
	return false
}

func (m *TemplatePOptions) GetErrWriter() bool {
	if m != nil {
		return m.ErrWriter
	}
	return false
}

func (m *TemplatePOptions) GetRestart() int32 {
	if m != nil {
		return m.Restart
	}
	return 0
}

func (m *TemplatePOptions) GetDetach() bool {
	if m != nil {
		return m.Detach
	}
	return false
}

func (m *TemplatePOptions) GetRunOnStartup() bool {
	if m != nil {
		return m.RunOnStartup
	}
	return false
}

func (m *TemplatePOptions) GetNotify() bool {
	if m != nil {
		return m.Notify
	}
	return false
}

func (m *TemplatePOptions) GetAutoTerminate() bool {
	if m != nil {
		return m.AutoTerminate
	}
	return false
}

func (m *TemplatePOptions) GetCpuAffinity() string {
	if m != nil {
		return m.CpuAffinity
	}
	return ""
}

func (m *TemplatePOptions) GetCpuAffinityDelay() string {
	if m != nil {
		return m.CpuAffinityDelay
	}
	return ""
}

func init() {
	proto.RegisterType((*Template)(nil), "process.Template")
	proto.RegisterType((*TemplatePOptions)(nil), "process.Template.pOptions")
}

func init() { proto.RegisterFile("process.proto", fileDescriptor_54c4d0e8c0aaf5c3) }

var fileDescriptor_54c4d0e8c0aaf5c3 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcd, 0x6a, 0xc3, 0x30,
	0x10, 0x84, 0x71, 0x9c, 0x1f, 0x7b, 0xf3, 0x43, 0xd0, 0xa1, 0x88, 0x40, 0xc1, 0x2d, 0x2d, 0xe4,
	0x50, 0x72, 0x68, 0x0f, 0x3d, 0x17, 0x7a, 0x0f, 0xb8, 0x81, 0x1e, 0x85, 0xea, 0x28, 0xad, 0x21,
	0x96, 0xc4, 0x7a, 0x45, 0xf1, 0x5b, 0xf5, 0x21, 0xfa, 0x60, 0x45, 0xb2, 0x55, 0x72, 0x9b, 0xf9,
	0x66, 0x76, 0x17, 0x24, 0x58, 0x5a, 0x34, 0x95, 0x6a, 0xdb, 0x9d, 0x45, 0x43, 0x86, 0xcd, 0x06,
	0x7b, 0xfb, 0x93, 0x42, 0x76, 0x50, 0x8d, 0x3d, 0x4b, 0x52, 0x8c, 0xc1, 0x58, 0xcb, 0x46, 0xf1,
	0xa4, 0x48, 0xb6, 0x79, 0x19, 0x34, 0x5b, 0x43, 0x5a, 0x35, 0x47, 0x3e, 0x0a, 0xc8, 0x4b, 0xf6,
	0x0c, 0xb9, 0x15, 0xc6, 0x52, 0x6d, 0x74, 0xcb, 0xd3, 0x22, 0xd9, 0xce, 0x1f, 0x37, 0xbb, 0xb8,
	0x3e, 0xee, 0xda, 0xd9, 0x7d, 0xdf, 0x28, 0xb3, 0xa8, 0x36, 0xbf, 0x23, 0xf8, 0x37, 0xfe, 0x96,
	0xc4, 0xcf, 0x96, 0x27, 0x45, 0xea, 0x6f, 0x79, 0xcd, 0xae, 0x01, 0x8c, 0x23, 0xf1, 0x8d, 0x35,
	0x29, 0x0c, 0x27, 0xb3, 0x32, 0x37, 0x8e, 0xde, 0x03, 0xf0, 0xb1, 0x42, 0x8c, 0x71, 0xda, 0xc7,
	0x0a, 0x71, 0x88, 0x39, 0xcc, 0x50, 0xb5, 0x24, 0x91, 0xf8, 0xb8, 0x48, 0xb6, 0x93, 0x32, 0x5a,
	0x76, 0x05, 0xd3, 0xa3, 0x22, 0x59, 0x7d, 0xf1, 0x49, 0x18, 0x1a, 0x1c, 0xbb, 0x83, 0x15, 0x3a,
	0x2d, 0x8c, 0x16, 0xa1, 0xe7, 0x2c, 0x9f, 0x86, 0x7c, 0x81, 0x4e, 0xef, 0xf5, 0x5b, 0xcf, 0xfc,
	0xb4, 0x36, 0x54, 0x9f, 0x3a, 0x3e, 0xeb, 0xa7, 0x7b, 0xc7, 0xee, 0x61, 0x25, 0x1d, 0x19, 0x41,
	0x0a, 0x9b, 0x5a, 0x4b, 0x52, 0x3c, 0x0b, 0xf9, 0xd2, 0xd3, 0x43, 0x84, 0xec, 0x06, 0x16, 0x95,
	0x75, 0x42, 0x9e, 0x4e, 0xb5, 0xae, 0xa9, 0xe3, 0x79, 0x78, 0xc9, 0x79, 0x65, 0xdd, 0xcb, 0x80,
	0xd8, 0x03, 0xb0, 0xcb, 0x8a, 0x38, 0xaa, 0xb3, 0xec, 0x38, 0x84, 0xe2, 0xfa, 0xa2, 0xf8, 0xea,
	0xf9, 0xc7, 0x34, 0x7c, 0xe1, 0xd3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xbb, 0x3e, 0xd6,
	0xd3, 0x01, 0x00, 0x00,
}
