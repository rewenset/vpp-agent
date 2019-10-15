// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/linux/iptables/iptables.proto

package linux_iptables

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	namespace "github.com/ligato/vpp-agent/api/models/linux/namespace"
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

type RuleChain_Protocol int32

const (
	RuleChain_IPv4 RuleChain_Protocol = 0
	RuleChain_IPv6 RuleChain_Protocol = 1
)

var RuleChain_Protocol_name = map[int32]string{
	0: "IPv4",
	1: "IPv6",
}

var RuleChain_Protocol_value = map[string]int32{
	"IPv4": 0,
	"IPv6": 1,
}

func (x RuleChain_Protocol) String() string {
	return proto.EnumName(RuleChain_Protocol_name, int32(x))
}

func (RuleChain_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f96037bbc28a0531, []int{0, 0}
}

type RuleChain_Table int32

const (
	RuleChain_FILTER   RuleChain_Table = 0
	RuleChain_NAT      RuleChain_Table = 1
	RuleChain_MANGLE   RuleChain_Table = 2
	RuleChain_RAW      RuleChain_Table = 3
	RuleChain_SECURITY RuleChain_Table = 4
)

var RuleChain_Table_name = map[int32]string{
	0: "FILTER",
	1: "NAT",
	2: "MANGLE",
	3: "RAW",
	4: "SECURITY",
}

var RuleChain_Table_value = map[string]int32{
	"FILTER":   0,
	"NAT":      1,
	"MANGLE":   2,
	"RAW":      3,
	"SECURITY": 4,
}

func (x RuleChain_Table) String() string {
	return proto.EnumName(RuleChain_Table_name, int32(x))
}

func (RuleChain_Table) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f96037bbc28a0531, []int{0, 1}
}

type RuleChain_ChainType int32

const (
	RuleChain_CUSTOM      RuleChain_ChainType = 0
	RuleChain_INPUT       RuleChain_ChainType = 1
	RuleChain_OUTPUT      RuleChain_ChainType = 2
	RuleChain_FORWARD     RuleChain_ChainType = 3
	RuleChain_PREROUTING  RuleChain_ChainType = 4
	RuleChain_POSTROUTING RuleChain_ChainType = 5
)

var RuleChain_ChainType_name = map[int32]string{
	0: "CUSTOM",
	1: "INPUT",
	2: "OUTPUT",
	3: "FORWARD",
	4: "PREROUTING",
	5: "POSTROUTING",
}

var RuleChain_ChainType_value = map[string]int32{
	"CUSTOM":      0,
	"INPUT":       1,
	"OUTPUT":      2,
	"FORWARD":     3,
	"PREROUTING":  4,
	"POSTROUTING": 5,
}

func (x RuleChain_ChainType) String() string {
	return proto.EnumName(RuleChain_ChainType_name, int32(x))
}

func (RuleChain_ChainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f96037bbc28a0531, []int{0, 2}
}

type RuleChain_Policy int32

const (
	RuleChain_NONE   RuleChain_Policy = 0
	RuleChain_ACCEPT RuleChain_Policy = 1
	RuleChain_DROP   RuleChain_Policy = 2
	RuleChain_QUEUE  RuleChain_Policy = 3
	RuleChain_RETURN RuleChain_Policy = 4
)

var RuleChain_Policy_name = map[int32]string{
	0: "NONE",
	1: "ACCEPT",
	2: "DROP",
	3: "QUEUE",
	4: "RETURN",
}

var RuleChain_Policy_value = map[string]int32{
	"NONE":   0,
	"ACCEPT": 1,
	"DROP":   2,
	"QUEUE":  3,
	"RETURN": 4,
}

func (x RuleChain_Policy) String() string {
	return proto.EnumName(RuleChain_Policy_name, int32(x))
}

func (RuleChain_Policy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f96037bbc28a0531, []int{0, 3}
}

type RuleChain struct {
	Name                 string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            *namespace.NetNamespace `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Interfaces           []string                `protobuf:"bytes,3,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	Protocol             RuleChain_Protocol      `protobuf:"varint,4,opt,name=protocol,proto3,enum=linux.iptables.RuleChain_Protocol" json:"protocol,omitempty"`
	Table                RuleChain_Table         `protobuf:"varint,5,opt,name=table,proto3,enum=linux.iptables.RuleChain_Table" json:"table,omitempty"`
	ChainType            RuleChain_ChainType     `protobuf:"varint,6,opt,name=chain_type,json=chainType,proto3,enum=linux.iptables.RuleChain_ChainType" json:"chain_type,omitempty"`
	ChainName            string                  `protobuf:"bytes,7,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	DefaultPolicy        RuleChain_Policy        `protobuf:"varint,8,opt,name=default_policy,json=defaultPolicy,proto3,enum=linux.iptables.RuleChain_Policy" json:"default_policy,omitempty"`
	Rules                []string                `protobuf:"bytes,10,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RuleChain) Reset()         { *m = RuleChain{} }
func (m *RuleChain) String() string { return proto.CompactTextString(m) }
func (*RuleChain) ProtoMessage()    {}
func (*RuleChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96037bbc28a0531, []int{0}
}

func (m *RuleChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleChain.Unmarshal(m, b)
}
func (m *RuleChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleChain.Marshal(b, m, deterministic)
}
func (m *RuleChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleChain.Merge(m, src)
}
func (m *RuleChain) XXX_Size() int {
	return xxx_messageInfo_RuleChain.Size(m)
}
func (m *RuleChain) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleChain.DiscardUnknown(m)
}

var xxx_messageInfo_RuleChain proto.InternalMessageInfo

func (m *RuleChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RuleChain) GetNamespace() *namespace.NetNamespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *RuleChain) GetInterfaces() []string {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *RuleChain) GetProtocol() RuleChain_Protocol {
	if m != nil {
		return m.Protocol
	}
	return RuleChain_IPv4
}

func (m *RuleChain) GetTable() RuleChain_Table {
	if m != nil {
		return m.Table
	}
	return RuleChain_FILTER
}

func (m *RuleChain) GetChainType() RuleChain_ChainType {
	if m != nil {
		return m.ChainType
	}
	return RuleChain_CUSTOM
}

func (m *RuleChain) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *RuleChain) GetDefaultPolicy() RuleChain_Policy {
	if m != nil {
		return m.DefaultPolicy
	}
	return RuleChain_NONE
}

func (m *RuleChain) GetRules() []string {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterEnum("linux.iptables.RuleChain_Protocol", RuleChain_Protocol_name, RuleChain_Protocol_value)
	proto.RegisterEnum("linux.iptables.RuleChain_Table", RuleChain_Table_name, RuleChain_Table_value)
	proto.RegisterEnum("linux.iptables.RuleChain_ChainType", RuleChain_ChainType_name, RuleChain_ChainType_value)
	proto.RegisterEnum("linux.iptables.RuleChain_Policy", RuleChain_Policy_name, RuleChain_Policy_value)
	proto.RegisterType((*RuleChain)(nil), "linux.iptables.RuleChain")
}

func init() {
	proto.RegisterFile("models/linux/iptables/iptables.proto", fileDescriptor_f96037bbc28a0531)
}

var fileDescriptor_f96037bbc28a0531 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x8b, 0x9b, 0x40,
	0x14, 0x8d, 0x51, 0x13, 0xbd, 0x69, 0xd3, 0x61, 0xe8, 0x83, 0x14, 0x76, 0x1b, 0x6c, 0x29, 0x79,
	0xa9, 0xc2, 0xf6, 0xe3, 0x65, 0xa1, 0x4b, 0x36, 0x31, 0x21, 0xb0, 0xab, 0x76, 0x32, 0xb2, 0xb4,
	0x14, 0x82, 0x71, 0x67, 0xb3, 0x82, 0x89, 0x92, 0xe8, 0xd2, 0xfc, 0xb0, 0xfe, 0xbf, 0x32, 0xe3,
	0x47, 0xba, 0xd0, 0xdd, 0x17, 0x39, 0xf7, 0xcc, 0x39, 0xf7, 0xce, 0x3d, 0x38, 0xf0, 0x7e, 0x93,
	0xde, 0xb2, 0x64, 0x6f, 0x27, 0xf1, 0xb6, 0xf8, 0x6d, 0xc7, 0x59, 0x1e, 0xae, 0x12, 0xb6, 0x6f,
	0x80, 0x95, 0xed, 0xd2, 0x3c, 0xc5, 0x7d, 0x71, 0x6c, 0xd5, 0xec, 0x9b, 0x0f, 0x8f, 0x5c, 0xdb,
	0x70, 0xc3, 0xf6, 0x59, 0x18, 0xb1, 0x23, 0x2a, 0x7d, 0xe6, 0x1f, 0x15, 0x74, 0x52, 0x24, 0x6c,
	0x7c, 0x1f, 0xc6, 0x5b, 0x8c, 0x41, 0xe1, 0x02, 0x43, 0x1a, 0x48, 0x43, 0x9d, 0x08, 0x8c, 0xcf,
	0x41, 0x6f, 0x4c, 0x46, 0x7b, 0x20, 0x0d, 0x7b, 0x67, 0x27, 0x56, 0x39, 0xed, 0xd8, 0xcc, 0x65,
	0xb9, 0x5b, 0x17, 0xe4, 0xa8, 0xc7, 0xa7, 0x00, 0xf1, 0x36, 0x67, 0xbb, 0xbb, 0x30, 0x62, 0x7b,
	0x43, 0x1e, 0xc8, 0x43, 0x9d, 0xfc, 0xc3, 0xe0, 0x6f, 0xa0, 0x89, 0x7b, 0x44, 0x69, 0x62, 0x28,
	0x03, 0x69, 0xd8, 0x3f, 0x33, 0xad, 0xc7, 0x9b, 0x58, 0xcd, 0xed, 0x2c, 0xbf, 0x52, 0x92, 0xc6,
	0x83, 0xbf, 0x80, 0x2a, 0x64, 0x86, 0x2a, 0xcc, 0x6f, 0x9f, 0x36, 0x53, 0x4e, 0x90, 0x52, 0x8d,
	0x2f, 0x01, 0x22, 0xce, 0x2e, 0xf3, 0x43, 0xc6, 0x8c, 0x8e, 0xf0, 0xbe, 0x7b, 0xda, 0x2b, 0xbe,
	0xf4, 0x90, 0x31, 0xa2, 0x47, 0x35, 0xc4, 0x27, 0x75, 0x0f, 0x91, 0x58, 0x57, 0x24, 0x56, 0x1e,
	0xf3, 0x24, 0xf0, 0x0c, 0xfa, 0xb7, 0xec, 0x2e, 0x2c, 0x92, 0x7c, 0x99, 0xa5, 0x49, 0x1c, 0x1d,
	0x0c, 0x4d, 0x8c, 0x19, 0x3c, 0xb3, 0x9f, 0xd0, 0x91, 0x97, 0x95, 0xaf, 0x2c, 0xf1, 0x6b, 0x50,
	0x77, 0x45, 0xc2, 0xf6, 0x06, 0x88, 0xf4, 0xca, 0xc2, 0x3c, 0x05, 0xad, 0x8e, 0x03, 0x6b, 0xa0,
	0xcc, 0xfd, 0x87, 0xcf, 0xa8, 0x55, 0xa1, 0xaf, 0x48, 0x32, 0x2f, 0x40, 0x15, 0x1b, 0x63, 0x80,
	0xce, 0x74, 0x7e, 0x45, 0x1d, 0x82, 0x5a, 0xb8, 0x0b, 0xb2, 0x3b, 0xa2, 0x48, 0xe2, 0xe4, 0xf5,
	0xc8, 0x9d, 0x5d, 0x39, 0xa8, 0xcd, 0x49, 0x32, 0xba, 0x41, 0x32, 0x7e, 0x01, 0xda, 0xc2, 0x19,
	0x07, 0x64, 0x4e, 0x7f, 0x20, 0xc5, 0xfc, 0x05, 0x7a, 0xb3, 0x36, 0xd7, 0x8f, 0x83, 0x05, 0xf5,
	0xae, 0x51, 0x0b, 0xeb, 0xa0, 0xce, 0x5d, 0x3f, 0xa8, 0xda, 0x78, 0x01, 0xe5, 0xb8, 0x8d, 0x7b,
	0xd0, 0x9d, 0x7a, 0xe4, 0x66, 0x44, 0x26, 0x48, 0xc6, 0x7d, 0x00, 0x9f, 0x38, 0xc4, 0x0b, 0xe8,
	0xdc, 0x9d, 0x21, 0x05, 0xbf, 0x82, 0x9e, 0xef, 0x2d, 0x68, 0x4d, 0xa8, 0xe6, 0x05, 0x74, 0xaa,
	0xf5, 0x34, 0x50, 0x5c, 0xcf, 0x75, 0x50, 0x8b, 0x77, 0x1b, 0x8d, 0xc7, 0x8e, 0xcf, 0x3b, 0x6b,
	0xa0, 0x4c, 0x88, 0xe7, 0xa3, 0x36, 0x1f, 0xf7, 0x3d, 0x70, 0x02, 0x07, 0xc9, 0x5c, 0x40, 0x1c,
	0x1a, 0x10, 0x17, 0x29, 0x97, 0xd3, 0x9f, 0x93, 0x75, 0x9c, 0xdf, 0x17, 0x2b, 0x2b, 0x4a, 0x37,
	0x76, 0x12, 0xaf, 0xc3, 0x3c, 0xb5, 0x1f, 0xb2, 0xec, 0x63, 0xb8, 0x66, 0xdb, 0xdc, 0x0e, 0xb3,
	0xd8, 0xfe, 0xef, 0xbb, 0x39, 0x17, 0xe5, 0xb2, 0x2e, 0x57, 0x1d, 0xf1, 0x2b, 0x7d, 0xfa, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x9d, 0xf4, 0x62, 0x1c, 0x66, 0x03, 0x00, 0x00,
}
