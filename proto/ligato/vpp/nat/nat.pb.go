// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/nat/nat.proto

package vpp_nat

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

// Available protocols.
type DNat44_Protocol int32

const (
	DNat44_TCP DNat44_Protocol = 0
	DNat44_UDP DNat44_Protocol = 1
	// ICMP is not permitted for load balanced entries.
	DNat44_ICMP DNat44_Protocol = 2
)

var DNat44_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
	2: "ICMP",
}

var DNat44_Protocol_value = map[string]int32{
	"TCP":  0,
	"UDP":  1,
	"ICMP": 2,
}

func (x DNat44_Protocol) String() string {
	return proto.EnumName(DNat44_Protocol_name, int32(x))
}

func (DNat44_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{1, 0}
}

// Available twice-NAT modes.
type DNat44_StaticMapping_TwiceNatMode int32

const (
	DNat44_StaticMapping_DISABLED DNat44_StaticMapping_TwiceNatMode = 0
	DNat44_StaticMapping_ENABLED  DNat44_StaticMapping_TwiceNatMode = 1
	DNat44_StaticMapping_SELF     DNat44_StaticMapping_TwiceNatMode = 2
)

var DNat44_StaticMapping_TwiceNatMode_name = map[int32]string{
	0: "DISABLED",
	1: "ENABLED",
	2: "SELF",
}

var DNat44_StaticMapping_TwiceNatMode_value = map[string]int32{
	"DISABLED": 0,
	"ENABLED":  1,
	"SELF":     2,
}

func (x DNat44_StaticMapping_TwiceNatMode) String() string {
	return proto.EnumName(DNat44_StaticMapping_TwiceNatMode_name, int32(x))
}

func (DNat44_StaticMapping_TwiceNatMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{1, 0, 0}
}

// Nat44Global defines global NAT44 configuration.
type Nat44Global struct {
	// Enable/disable forwarding.
	Forwarding bool `protobuf:"varint,1,opt,name=forwarding,proto3" json:"forwarding,omitempty"`
	// List of NAT-enabled interfaces. Deprecated - use separate Nat44Interface entries instead.
	NatInterfaces []*Nat44Global_Interface `protobuf:"bytes,2,rep,name=nat_interfaces,json=natInterfaces,proto3" json:"nat_interfaces,omitempty"` // Deprecated: Do not use.
	// Address pool used for source IP NAT. Deprecated - use separate Nat44AddressPool entries instead.
	AddressPool []*Nat44Global_Address `protobuf:"bytes,3,rep,name=address_pool,json=addressPool,proto3" json:"address_pool,omitempty"` // Deprecated: Do not use.
	// Virtual reassembly for IPv4.
	VirtualReassembly    *VirtualReassembly `protobuf:"bytes,4,opt,name=virtual_reassembly,json=virtualReassembly,proto3" json:"virtual_reassembly,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Nat44Global) Reset()         { *m = Nat44Global{} }
func (m *Nat44Global) String() string { return proto.CompactTextString(m) }
func (*Nat44Global) ProtoMessage()    {}
func (*Nat44Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{0}
}

func (m *Nat44Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nat44Global.Unmarshal(m, b)
}
func (m *Nat44Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nat44Global.Marshal(b, m, deterministic)
}
func (m *Nat44Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nat44Global.Merge(m, src)
}
func (m *Nat44Global) XXX_Size() int {
	return xxx_messageInfo_Nat44Global.Size(m)
}
func (m *Nat44Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Nat44Global.DiscardUnknown(m)
}

var xxx_messageInfo_Nat44Global proto.InternalMessageInfo

func (m *Nat44Global) GetForwarding() bool {
	if m != nil {
		return m.Forwarding
	}
	return false
}

// Deprecated: Do not use.
func (m *Nat44Global) GetNatInterfaces() []*Nat44Global_Interface {
	if m != nil {
		return m.NatInterfaces
	}
	return nil
}

// Deprecated: Do not use.
func (m *Nat44Global) GetAddressPool() []*Nat44Global_Address {
	if m != nil {
		return m.AddressPool
	}
	return nil
}

func (m *Nat44Global) GetVirtualReassembly() *VirtualReassembly {
	if m != nil {
		return m.VirtualReassembly
	}
	return nil
}

// Interface defines a network interface enabled for NAT.
type Nat44Global_Interface struct {
	// Interface name (logical).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Distinguish between inside/outside interface.
	IsInside bool `protobuf:"varint,2,opt,name=is_inside,json=isInside,proto3" json:"is_inside,omitempty"`
	// Enable/disable output feature.
	OutputFeature        bool     `protobuf:"varint,3,opt,name=output_feature,json=outputFeature,proto3" json:"output_feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nat44Global_Interface) Reset()         { *m = Nat44Global_Interface{} }
func (m *Nat44Global_Interface) String() string { return proto.CompactTextString(m) }
func (*Nat44Global_Interface) ProtoMessage()    {}
func (*Nat44Global_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{0, 0}
}

func (m *Nat44Global_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nat44Global_Interface.Unmarshal(m, b)
}
func (m *Nat44Global_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nat44Global_Interface.Marshal(b, m, deterministic)
}
func (m *Nat44Global_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nat44Global_Interface.Merge(m, src)
}
func (m *Nat44Global_Interface) XXX_Size() int {
	return xxx_messageInfo_Nat44Global_Interface.Size(m)
}
func (m *Nat44Global_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_Nat44Global_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_Nat44Global_Interface proto.InternalMessageInfo

func (m *Nat44Global_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nat44Global_Interface) GetIsInside() bool {
	if m != nil {
		return m.IsInside
	}
	return false
}

func (m *Nat44Global_Interface) GetOutputFeature() bool {
	if m != nil {
		return m.OutputFeature
	}
	return false
}

// Address defines an address to be used for source IP NAT.
type Nat44Global_Address struct {
	// IPv4 address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// VRF id of tenant, 0xFFFFFFFF means independent of VRF.
	// Non-zero (and not all-ones) VRF has to be explicitly created (see api/models/vpp/l3/vrf.proto).
	VrfId uint32 `protobuf:"varint,2,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	// Enable/disable twice NAT.
	TwiceNat             bool     `protobuf:"varint,3,opt,name=twice_nat,json=twiceNat,proto3" json:"twice_nat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nat44Global_Address) Reset()         { *m = Nat44Global_Address{} }
func (m *Nat44Global_Address) String() string { return proto.CompactTextString(m) }
func (*Nat44Global_Address) ProtoMessage()    {}
func (*Nat44Global_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{0, 1}
}

func (m *Nat44Global_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nat44Global_Address.Unmarshal(m, b)
}
func (m *Nat44Global_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nat44Global_Address.Marshal(b, m, deterministic)
}
func (m *Nat44Global_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nat44Global_Address.Merge(m, src)
}
func (m *Nat44Global_Address) XXX_Size() int {
	return xxx_messageInfo_Nat44Global_Address.Size(m)
}
func (m *Nat44Global_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Nat44Global_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Nat44Global_Address proto.InternalMessageInfo

func (m *Nat44Global_Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Nat44Global_Address) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Nat44Global_Address) GetTwiceNat() bool {
	if m != nil {
		return m.TwiceNat
	}
	return false
}

// DNat44 defines destination NAT44 configuration.
type DNat44 struct {
	// Unique identifier for the DNAT configuration.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// A list of static mappings in DNAT.
	StMappings []*DNat44_StaticMapping `protobuf:"bytes,2,rep,name=st_mappings,json=stMappings,proto3" json:"st_mappings,omitempty"`
	// A list of identity mappings in DNAT.
	IdMappings           []*DNat44_IdentityMapping `protobuf:"bytes,3,rep,name=id_mappings,json=idMappings,proto3" json:"id_mappings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DNat44) Reset()         { *m = DNat44{} }
func (m *DNat44) String() string { return proto.CompactTextString(m) }
func (*DNat44) ProtoMessage()    {}
func (*DNat44) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{1}
}

func (m *DNat44) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNat44.Unmarshal(m, b)
}
func (m *DNat44) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNat44.Marshal(b, m, deterministic)
}
func (m *DNat44) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNat44.Merge(m, src)
}
func (m *DNat44) XXX_Size() int {
	return xxx_messageInfo_DNat44.Size(m)
}
func (m *DNat44) XXX_DiscardUnknown() {
	xxx_messageInfo_DNat44.DiscardUnknown(m)
}

var xxx_messageInfo_DNat44 proto.InternalMessageInfo

func (m *DNat44) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *DNat44) GetStMappings() []*DNat44_StaticMapping {
	if m != nil {
		return m.StMappings
	}
	return nil
}

func (m *DNat44) GetIdMappings() []*DNat44_IdentityMapping {
	if m != nil {
		return m.IdMappings
	}
	return nil
}

// StaticMapping defines a list of static mappings in DNAT.
type DNat44_StaticMapping struct {
	// Interface to use external IP from; preferred over external_ip.
	ExternalInterface string `protobuf:"bytes,1,opt,name=external_interface,json=externalInterface,proto3" json:"external_interface,omitempty"`
	// External address.
	ExternalIp string `protobuf:"bytes,2,opt,name=external_ip,json=externalIp,proto3" json:"external_ip,omitempty"`
	// Port (do not set for address mapping).
	ExternalPort uint32 `protobuf:"varint,3,opt,name=external_port,json=externalPort,proto3" json:"external_port,omitempty"`
	// List of local IP addresses. If there is more than one entry, load-balancing is enabled.
	LocalIps []*DNat44_StaticMapping_LocalIP `protobuf:"bytes,4,rep,name=local_ips,json=localIps,proto3" json:"local_ips,omitempty"`
	// Protocol used for static mapping.
	Protocol DNat44_Protocol `protobuf:"varint,5,opt,name=protocol,proto3,enum=ligato.vpp.nat.DNat44_Protocol" json:"protocol,omitempty"`
	// Enable/disable (self-)twice NAT.
	TwiceNat DNat44_StaticMapping_TwiceNatMode `protobuf:"varint,6,opt,name=twice_nat,json=twiceNat,proto3,enum=ligato.vpp.nat.DNat44_StaticMapping_TwiceNatMode" json:"twice_nat,omitempty"`
	// Session affinity. 0 means disabled, otherwise client IP affinity sticky time in seconds.
	SessionAffinity      uint32   `protobuf:"varint,7,opt,name=session_affinity,json=sessionAffinity,proto3" json:"session_affinity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNat44_StaticMapping) Reset()         { *m = DNat44_StaticMapping{} }
func (m *DNat44_StaticMapping) String() string { return proto.CompactTextString(m) }
func (*DNat44_StaticMapping) ProtoMessage()    {}
func (*DNat44_StaticMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{1, 0}
}

func (m *DNat44_StaticMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNat44_StaticMapping.Unmarshal(m, b)
}
func (m *DNat44_StaticMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNat44_StaticMapping.Marshal(b, m, deterministic)
}
func (m *DNat44_StaticMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNat44_StaticMapping.Merge(m, src)
}
func (m *DNat44_StaticMapping) XXX_Size() int {
	return xxx_messageInfo_DNat44_StaticMapping.Size(m)
}
func (m *DNat44_StaticMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_DNat44_StaticMapping.DiscardUnknown(m)
}

var xxx_messageInfo_DNat44_StaticMapping proto.InternalMessageInfo

func (m *DNat44_StaticMapping) GetExternalInterface() string {
	if m != nil {
		return m.ExternalInterface
	}
	return ""
}

func (m *DNat44_StaticMapping) GetExternalIp() string {
	if m != nil {
		return m.ExternalIp
	}
	return ""
}

func (m *DNat44_StaticMapping) GetExternalPort() uint32 {
	if m != nil {
		return m.ExternalPort
	}
	return 0
}

func (m *DNat44_StaticMapping) GetLocalIps() []*DNat44_StaticMapping_LocalIP {
	if m != nil {
		return m.LocalIps
	}
	return nil
}

func (m *DNat44_StaticMapping) GetProtocol() DNat44_Protocol {
	if m != nil {
		return m.Protocol
	}
	return DNat44_TCP
}

func (m *DNat44_StaticMapping) GetTwiceNat() DNat44_StaticMapping_TwiceNatMode {
	if m != nil {
		return m.TwiceNat
	}
	return DNat44_StaticMapping_DISABLED
}

func (m *DNat44_StaticMapping) GetSessionAffinity() uint32 {
	if m != nil {
		return m.SessionAffinity
	}
	return 0
}

// LocalIP defines a local IP addresses.
type DNat44_StaticMapping_LocalIP struct {
	// VRF (table) ID. Non-zero VRF has to be explicitly created (see api/models/vpp/l3/vrf.proto).
	VrfId uint32 `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	// Local IP address).
	LocalIp string `protobuf:"bytes,2,opt,name=local_ip,json=localIp,proto3" json:"local_ip,omitempty"`
	// Port (do not set for address mapping).
	LocalPort uint32 `protobuf:"varint,3,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	// Probability level for load-balancing mode.
	Probability          uint32   `protobuf:"varint,4,opt,name=probability,proto3" json:"probability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNat44_StaticMapping_LocalIP) Reset()         { *m = DNat44_StaticMapping_LocalIP{} }
func (m *DNat44_StaticMapping_LocalIP) String() string { return proto.CompactTextString(m) }
func (*DNat44_StaticMapping_LocalIP) ProtoMessage()    {}
func (*DNat44_StaticMapping_LocalIP) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{1, 0, 0}
}

func (m *DNat44_StaticMapping_LocalIP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNat44_StaticMapping_LocalIP.Unmarshal(m, b)
}
func (m *DNat44_StaticMapping_LocalIP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNat44_StaticMapping_LocalIP.Marshal(b, m, deterministic)
}
func (m *DNat44_StaticMapping_LocalIP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNat44_StaticMapping_LocalIP.Merge(m, src)
}
func (m *DNat44_StaticMapping_LocalIP) XXX_Size() int {
	return xxx_messageInfo_DNat44_StaticMapping_LocalIP.Size(m)
}
func (m *DNat44_StaticMapping_LocalIP) XXX_DiscardUnknown() {
	xxx_messageInfo_DNat44_StaticMapping_LocalIP.DiscardUnknown(m)
}

var xxx_messageInfo_DNat44_StaticMapping_LocalIP proto.InternalMessageInfo

func (m *DNat44_StaticMapping_LocalIP) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *DNat44_StaticMapping_LocalIP) GetLocalIp() string {
	if m != nil {
		return m.LocalIp
	}
	return ""
}

func (m *DNat44_StaticMapping_LocalIP) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *DNat44_StaticMapping_LocalIP) GetProbability() uint32 {
	if m != nil {
		return m.Probability
	}
	return 0
}

// IdentityMapping defines an identity mapping in DNAT.
type DNat44_IdentityMapping struct {
	// VRF (table) ID. Non-zero VRF has to be explicitly created (see api/models/vpp/l3/vrf.proto).
	VrfId uint32 `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	// Name of the interface to use address from; preferred over ip_address.
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	// IP address.
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// Port (do not set for address mapping).
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// Protocol used for identity mapping.
	Protocol             DNat44_Protocol `protobuf:"varint,5,opt,name=protocol,proto3,enum=ligato.vpp.nat.DNat44_Protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DNat44_IdentityMapping) Reset()         { *m = DNat44_IdentityMapping{} }
func (m *DNat44_IdentityMapping) String() string { return proto.CompactTextString(m) }
func (*DNat44_IdentityMapping) ProtoMessage()    {}
func (*DNat44_IdentityMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{1, 1}
}

func (m *DNat44_IdentityMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNat44_IdentityMapping.Unmarshal(m, b)
}
func (m *DNat44_IdentityMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNat44_IdentityMapping.Marshal(b, m, deterministic)
}
func (m *DNat44_IdentityMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNat44_IdentityMapping.Merge(m, src)
}
func (m *DNat44_IdentityMapping) XXX_Size() int {
	return xxx_messageInfo_DNat44_IdentityMapping.Size(m)
}
func (m *DNat44_IdentityMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_DNat44_IdentityMapping.DiscardUnknown(m)
}

var xxx_messageInfo_DNat44_IdentityMapping proto.InternalMessageInfo

func (m *DNat44_IdentityMapping) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *DNat44_IdentityMapping) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *DNat44_IdentityMapping) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *DNat44_IdentityMapping) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *DNat44_IdentityMapping) GetProtocol() DNat44_Protocol {
	if m != nil {
		return m.Protocol
	}
	return DNat44_TCP
}

// Nat44Interface defines a local network interfaces enabled for NAT44.
type Nat44Interface struct {
	// Interface name (logical).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Enable/disable NAT on inside.
	NatInside bool `protobuf:"varint,2,opt,name=nat_inside,json=natInside,proto3" json:"nat_inside,omitempty"`
	// Enable/disable NAT on outside.
	NatOutside bool `protobuf:"varint,3,opt,name=nat_outside,json=natOutside,proto3" json:"nat_outside,omitempty"`
	//  Enable/disable output feature.
	OutputFeature        bool     `protobuf:"varint,4,opt,name=output_feature,json=outputFeature,proto3" json:"output_feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nat44Interface) Reset()         { *m = Nat44Interface{} }
func (m *Nat44Interface) String() string { return proto.CompactTextString(m) }
func (*Nat44Interface) ProtoMessage()    {}
func (*Nat44Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{2}
}

func (m *Nat44Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nat44Interface.Unmarshal(m, b)
}
func (m *Nat44Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nat44Interface.Marshal(b, m, deterministic)
}
func (m *Nat44Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nat44Interface.Merge(m, src)
}
func (m *Nat44Interface) XXX_Size() int {
	return xxx_messageInfo_Nat44Interface.Size(m)
}
func (m *Nat44Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_Nat44Interface.DiscardUnknown(m)
}

var xxx_messageInfo_Nat44Interface proto.InternalMessageInfo

func (m *Nat44Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nat44Interface) GetNatInside() bool {
	if m != nil {
		return m.NatInside
	}
	return false
}

func (m *Nat44Interface) GetNatOutside() bool {
	if m != nil {
		return m.NatOutside
	}
	return false
}

func (m *Nat44Interface) GetOutputFeature() bool {
	if m != nil {
		return m.OutputFeature
	}
	return false
}

// Nat44AddressPool defines an address pool used for NAT44.
type Nat44AddressPool struct {
	// VRF id of tenant, 0xFFFFFFFF means independent of VRF.
	// Non-zero (and not all-ones) VRF has to be explicitly created (see api/models/vpp/l3/vrf.proto).
	VrfId uint32 `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	// First IP address of the pool.
	FirstIp string `protobuf:"bytes,2,opt,name=first_ip,json=firstIp,proto3" json:"first_ip,omitempty"`
	// Last IP address of the pool. Should be higher than first_ip or empty.
	LastIp string `protobuf:"bytes,3,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	// Enable/disable twice NAT.
	TwiceNat             bool     `protobuf:"varint,4,opt,name=twice_nat,json=twiceNat,proto3" json:"twice_nat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nat44AddressPool) Reset()         { *m = Nat44AddressPool{} }
func (m *Nat44AddressPool) String() string { return proto.CompactTextString(m) }
func (*Nat44AddressPool) ProtoMessage()    {}
func (*Nat44AddressPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{3}
}

func (m *Nat44AddressPool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nat44AddressPool.Unmarshal(m, b)
}
func (m *Nat44AddressPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nat44AddressPool.Marshal(b, m, deterministic)
}
func (m *Nat44AddressPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nat44AddressPool.Merge(m, src)
}
func (m *Nat44AddressPool) XXX_Size() int {
	return xxx_messageInfo_Nat44AddressPool.Size(m)
}
func (m *Nat44AddressPool) XXX_DiscardUnknown() {
	xxx_messageInfo_Nat44AddressPool.DiscardUnknown(m)
}

var xxx_messageInfo_Nat44AddressPool proto.InternalMessageInfo

func (m *Nat44AddressPool) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Nat44AddressPool) GetFirstIp() string {
	if m != nil {
		return m.FirstIp
	}
	return ""
}

func (m *Nat44AddressPool) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

func (m *Nat44AddressPool) GetTwiceNat() bool {
	if m != nil {
		return m.TwiceNat
	}
	return false
}

// VirtualReassembly defines NAT virtual reassembly settings.
type VirtualReassembly struct {
	// Reassembly timeout.
	Timeout uint32 `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Maximum number of concurrent reassemblies.
	MaxReassemblies uint32 `protobuf:"varint,2,opt,name=max_reassemblies,json=maxReassemblies,proto3" json:"max_reassemblies,omitempty"`
	// Maximum number of fragments per reassembly.
	MaxFragments uint32 `protobuf:"varint,3,opt,name=max_fragments,json=maxFragments,proto3" json:"max_fragments,omitempty"`
	// If set to true fragments are dropped, translated otherwise.
	DropFragments        bool     `protobuf:"varint,4,opt,name=drop_fragments,json=dropFragments,proto3" json:"drop_fragments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualReassembly) Reset()         { *m = VirtualReassembly{} }
func (m *VirtualReassembly) String() string { return proto.CompactTextString(m) }
func (*VirtualReassembly) ProtoMessage()    {}
func (*VirtualReassembly) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c5496f531b4b7d3, []int{4}
}

func (m *VirtualReassembly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualReassembly.Unmarshal(m, b)
}
func (m *VirtualReassembly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualReassembly.Marshal(b, m, deterministic)
}
func (m *VirtualReassembly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualReassembly.Merge(m, src)
}
func (m *VirtualReassembly) XXX_Size() int {
	return xxx_messageInfo_VirtualReassembly.Size(m)
}
func (m *VirtualReassembly) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualReassembly.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualReassembly proto.InternalMessageInfo

func (m *VirtualReassembly) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *VirtualReassembly) GetMaxReassemblies() uint32 {
	if m != nil {
		return m.MaxReassemblies
	}
	return 0
}

func (m *VirtualReassembly) GetMaxFragments() uint32 {
	if m != nil {
		return m.MaxFragments
	}
	return 0
}

func (m *VirtualReassembly) GetDropFragments() bool {
	if m != nil {
		return m.DropFragments
	}
	return false
}

func init() {
	proto.RegisterEnum("ligato.vpp.nat.DNat44_Protocol", DNat44_Protocol_name, DNat44_Protocol_value)
	proto.RegisterEnum("ligato.vpp.nat.DNat44_StaticMapping_TwiceNatMode", DNat44_StaticMapping_TwiceNatMode_name, DNat44_StaticMapping_TwiceNatMode_value)
	proto.RegisterType((*Nat44Global)(nil), "ligato.vpp.nat.Nat44Global")
	proto.RegisterType((*Nat44Global_Interface)(nil), "ligato.vpp.nat.Nat44Global.Interface")
	proto.RegisterType((*Nat44Global_Address)(nil), "ligato.vpp.nat.Nat44Global.Address")
	proto.RegisterType((*DNat44)(nil), "ligato.vpp.nat.DNat44")
	proto.RegisterType((*DNat44_StaticMapping)(nil), "ligato.vpp.nat.DNat44.StaticMapping")
	proto.RegisterType((*DNat44_StaticMapping_LocalIP)(nil), "ligato.vpp.nat.DNat44.StaticMapping.LocalIP")
	proto.RegisterType((*DNat44_IdentityMapping)(nil), "ligato.vpp.nat.DNat44.IdentityMapping")
	proto.RegisterType((*Nat44Interface)(nil), "ligato.vpp.nat.Nat44Interface")
	proto.RegisterType((*Nat44AddressPool)(nil), "ligato.vpp.nat.Nat44AddressPool")
	proto.RegisterType((*VirtualReassembly)(nil), "ligato.vpp.nat.VirtualReassembly")
}

func init() { proto.RegisterFile("ligato/vpp/nat/nat.proto", fileDescriptor_6c5496f531b4b7d3) }

var fileDescriptor_6c5496f531b4b7d3 = []byte{
	// 892 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x6e, 0xdb, 0x36,
	0x14, 0xad, 0x6c, 0xc7, 0x92, 0xaf, 0x62, 0xd7, 0x21, 0x36, 0x4c, 0xf3, 0xd6, 0xc5, 0x73, 0xd7,
	0x22, 0x03, 0x56, 0x1b, 0x4b, 0x8a, 0x61, 0x40, 0x9f, 0x92, 0x26, 0x29, 0x34, 0x24, 0x99, 0xa0,
	0x74, 0x1b, 0xb0, 0x17, 0x81, 0xb6, 0x28, 0x83, 0x80, 0x2c, 0x12, 0x24, 0xed, 0x26, 0xc3, 0xfe,
	0x60, 0xbf, 0xb1, 0x6f, 0xd8, 0xc3, 0xbe, 0x61, 0x7f, 0xb0, 0x9f, 0x19, 0x44, 0x52, 0xb6, 0xe2,
	0x35, 0x45, 0x81, 0x3e, 0x18, 0x10, 0xcf, 0xbd, 0x3c, 0xbc, 0xf7, 0xf0, 0xfa, 0x10, 0x82, 0x9c,
	0xce, 0xb1, 0x62, 0x93, 0x15, 0xe7, 0x93, 0x02, 0xab, 0xf2, 0x37, 0xe6, 0x82, 0x29, 0x86, 0x7a,
	0x26, 0x32, 0x5e, 0x71, 0x3e, 0x2e, 0xb0, 0x1a, 0xfd, 0xdb, 0x04, 0xff, 0x0a, 0xab, 0xe7, 0xcf,
	0x5f, 0xe5, 0x6c, 0x8a, 0x73, 0xf4, 0x05, 0x40, 0xc6, 0xc4, 0x1b, 0x2c, 0x52, 0x5a, 0xcc, 0x03,
	0x67, 0xe8, 0x1c, 0x78, 0x71, 0x0d, 0x41, 0x11, 0xf4, 0x0a, 0xac, 0x12, 0x5a, 0x28, 0x22, 0x32,
	0x3c, 0x23, 0x32, 0x68, 0x0c, 0x9b, 0x07, 0xfe, 0xe1, 0x93, 0xf1, 0x5d, 0xe2, 0x71, 0x8d, 0x74,
	0x1c, 0x56, 0xd9, 0x27, 0x8d, 0xc0, 0x89, 0xbb, 0x05, 0x56, 0x6b, 0x44, 0xa2, 0x1f, 0x60, 0x17,
	0xa7, 0xa9, 0x20, 0x52, 0x26, 0x9c, 0xb1, 0x3c, 0x68, 0x6a, 0xbe, 0xc7, 0xef, 0xe2, 0x3b, 0x36,
	0xf9, 0x9a, 0xcd, 0xb7, 0x9b, 0x23, 0xc6, 0x72, 0x14, 0x01, 0x5a, 0x51, 0xa1, 0x96, 0x38, 0x4f,
	0x04, 0xc1, 0x52, 0x92, 0xc5, 0x34, 0xbf, 0x0d, 0x5a, 0x43, 0xe7, 0xc0, 0x3f, 0xfc, 0x72, 0x9b,
	0xf1, 0x67, 0x93, 0x19, 0xaf, 0x13, 0xe3, 0xbd, 0xd5, 0x36, 0x34, 0x98, 0x41, 0x67, 0x5d, 0x2b,
	0x42, 0xd0, 0x2a, 0xf0, 0x82, 0x68, 0x59, 0x3a, 0xb1, 0xfe, 0x46, 0x9f, 0x41, 0x87, 0xca, 0x84,
	0x16, 0x92, 0xa6, 0x24, 0x68, 0x68, 0xbd, 0x3c, 0x2a, 0x43, 0xbd, 0x46, 0x4f, 0xa0, 0xc7, 0x96,
	0x8a, 0x2f, 0x55, 0x92, 0x11, 0xac, 0x96, 0x82, 0x04, 0x4d, 0x9d, 0xd1, 0x35, 0xe8, 0xb9, 0x01,
	0x07, 0xbf, 0x80, 0x6b, 0x5b, 0x42, 0x01, 0xb8, 0xb6, 0x21, 0x7b, 0x4a, 0xb5, 0x44, 0x1f, 0x43,
	0x7b, 0x25, 0xb2, 0x84, 0xa6, 0xfa, 0x94, 0x6e, 0xbc, 0xb3, 0x12, 0x59, 0x98, 0x96, 0xe7, 0xab,
	0x37, 0x74, 0x46, 0x92, 0x02, 0x2b, 0xcb, 0xee, 0x69, 0xe0, 0x0a, 0xab, 0xd1, 0x3f, 0x2e, 0xb4,
	0x4f, 0xb5, 0x72, 0xe8, 0x23, 0xd8, 0xc9, 0xf1, 0x94, 0xe4, 0x96, 0xd6, 0x2c, 0xd0, 0x19, 0xf8,
	0x52, 0x25, 0x0b, 0xcc, 0x39, 0x2d, 0xe6, 0xd5, 0x5d, 0x7e, 0xb5, 0xad, 0x94, 0xa1, 0x18, 0x5f,
	0x2b, 0xac, 0xe8, 0xec, 0xd2, 0x24, 0xc7, 0x20, 0x95, 0xfd, 0x94, 0xe8, 0x15, 0xf8, 0x34, 0xdd,
	0xd0, 0x98, 0x2b, 0x7c, 0x7a, 0x0f, 0x4d, 0x98, 0x92, 0x42, 0x51, 0x75, 0xbb, 0x26, 0xa2, 0x69,
	0x45, 0x34, 0xf8, 0xbb, 0x05, 0xdd, 0x3b, 0xc7, 0xa0, 0x67, 0x80, 0xc8, 0x8d, 0x22, 0xa2, 0xc0,
	0xf9, 0x66, 0xea, 0x6c, 0x13, 0x7b, 0x55, 0x64, 0x73, 0x45, 0xfb, 0xe0, 0x6f, 0xd2, 0xb9, 0x96,
	0xaa, 0x13, 0xc3, 0x3a, 0x8f, 0xa3, 0xc7, 0xd0, 0x5d, 0x27, 0x70, 0x26, 0x8c, 0x66, 0xdd, 0x78,
	0xb7, 0x02, 0x23, 0x26, 0x14, 0x0a, 0xa1, 0x93, 0xb3, 0x99, 0xa6, 0x90, 0x41, 0x4b, 0x77, 0xf3,
	0xcd, 0xfb, 0x88, 0x32, 0xbe, 0x28, 0x77, 0x85, 0x51, 0xec, 0xe9, 0xed, 0x21, 0x97, 0xe8, 0x05,
	0x78, 0xfa, 0x9f, 0x37, 0x63, 0x79, 0xb0, 0x33, 0x74, 0x0e, 0x7a, 0x87, 0xfb, 0xf7, 0x30, 0x45,
	0x36, 0x2d, 0x5e, 0x6f, 0x40, 0x57, 0xf5, 0xcb, 0x6d, 0xeb, 0xdd, 0xdf, 0xbe, 0x57, 0x1d, 0xaf,
	0xed, 0x04, 0x5c, 0xb2, 0x94, 0x6c, 0xe6, 0x01, 0x7d, 0x0d, 0x7d, 0x49, 0xa4, 0xa4, 0xac, 0x48,
	0x70, 0x96, 0xd1, 0x82, 0xaa, 0xdb, 0xc0, 0xd5, 0xfd, 0x3f, 0xb4, 0xf8, 0xb1, 0x85, 0x07, 0xbf,
	0x83, 0x6b, 0x9b, 0xa9, 0x4d, 0x9e, 0x53, 0x9f, 0xbc, 0x4f, 0xc1, 0xab, 0x44, 0xb2, 0x3a, 0xbb,
	0xb6, 0x6b, 0xf4, 0x08, 0xc0, 0x84, 0x6a, 0x0a, 0x1b, 0x45, 0xb5, 0xbc, 0x43, 0xf0, 0xb9, 0x60,
	0x53, 0x3c, 0xa5, 0x79, 0x59, 0x41, 0x4b, 0xc7, 0xeb, 0xd0, 0xe8, 0x08, 0x76, 0xeb, 0x2d, 0xa0,
	0x5d, 0xf0, 0x4e, 0xc3, 0xeb, 0xe3, 0x93, 0x8b, 0xb3, 0xd3, 0xfe, 0x03, 0xe4, 0x83, 0x7b, 0x76,
	0x65, 0x16, 0x0e, 0xf2, 0xa0, 0x75, 0x7d, 0x76, 0x71, 0xde, 0x6f, 0x0c, 0xfe, 0x72, 0xe0, 0xe1,
	0xd6, 0x70, 0xdd, 0x57, 0xfb, 0xe7, 0xd0, 0xd9, 0x0c, 0x93, 0x29, 0x7e, 0x03, 0x94, 0xe5, 0x53,
	0x9e, 0x54, 0xff, 0xc3, 0xa6, 0x0d, 0xf3, 0xea, 0x3f, 0x8a, 0xa0, 0xa5, 0xfb, 0x32, 0x75, 0xeb,
	0xef, 0x0f, 0xba, 0xe6, 0xd1, 0x53, 0xf0, 0x2a, 0x14, 0xb9, 0xd0, 0x7c, 0xfd, 0x32, 0xea, 0x3f,
	0x28, 0x3f, 0x7e, 0x3a, 0x8d, 0x4c, 0x83, 0xe1, 0xcb, 0xcb, 0xa8, 0xdf, 0x18, 0xfd, 0xe1, 0x40,
	0x4f, 0x93, 0xbc, 0xdb, 0x92, 0x1e, 0x01, 0x18, 0x8f, 0xae, 0x79, 0x52, 0x47, 0x9b, 0xae, 0x36,
	0xa5, 0x7d, 0xf0, 0xcb, 0x30, 0x5b, 0x2a, 0x1d, 0x37, 0x9e, 0x51, 0xee, 0xf8, 0xd1, 0x20, 0x6f,
	0x71, 0xad, 0xd6, 0x5b, 0x5c, 0x6b, 0xf4, 0x1b, 0xf4, 0x75, 0x31, 0xc7, 0x35, 0x03, 0xbe, 0x7f,
	0x54, 0x32, 0x2a, 0xa4, 0xaa, 0x8d, 0x8a, 0x5e, 0x87, 0x1c, 0x7d, 0x02, 0x6e, 0x8e, 0x4d, 0xc4,
	0x08, 0xdd, 0x2e, 0x97, 0x21, 0xbf, 0x6b, 0x6c, 0xad, 0x2d, 0x63, 0xfb, 0xd3, 0x81, 0xbd, 0xff,
	0xf9, 0x77, 0x69, 0x9e, 0x8a, 0x2e, 0x08, 0x5b, 0x2a, 0x7b, 0x7c, 0xb5, 0x2c, 0x07, 0x7f, 0x81,
	0x6f, 0x36, 0x8f, 0x02, 0xd5, 0x0f, 0x97, 0x1e, 0xfc, 0x05, 0xbe, 0x89, 0x6b, 0x70, 0x69, 0x10,
	0x65, 0x6a, 0x26, 0xf0, 0x7c, 0x41, 0x0a, 0x25, 0x2b, 0x83, 0x58, 0xe0, 0x9b, 0xf3, 0x0a, 0x2b,
	0x25, 0x4a, 0x05, 0xe3, 0xb5, 0x2c, 0x2b, 0x51, 0x89, 0xae, 0xd3, 0x4e, 0xbe, 0xff, 0xf5, 0xbb,
	0x39, 0xab, 0xe6, 0x80, 0xea, 0xf7, 0xf8, 0x19, 0x9e, 0x93, 0x42, 0x4d, 0x56, 0x47, 0x13, 0x7d,
	0xfd, 0x93, 0xbb, 0x2f, 0xf5, 0x8b, 0x15, 0xe7, 0x65, 0xc3, 0xd3, 0xb6, 0x8e, 0x1e, 0xfd, 0x17,
	0x00, 0x00, 0xff, 0xff, 0xd6, 0xf8, 0x55, 0x38, 0xca, 0x07, 0x00, 0x00,
}