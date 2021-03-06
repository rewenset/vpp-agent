# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: models/vpp/nat/nat.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from github.com.gogo.protobuf.gogoproto import gogo_pb2 as github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='models/vpp/nat/nat.proto',
  package='vpp.nat',
  syntax='proto3',
  serialized_options=_b('Z6github.com/ligato/vpp-agent/api/models/vpp/nat;vpp_nat\310\343\036\001'),
  serialized_pb=_b('\n\x18models/vpp/nat/nat.proto\x12\x07vpp.nat\x1a-github.com/gogo/protobuf/gogoproto/gogo.proto\"\xca\x02\n\x0bNat44Global\x12\x12\n\nforwarding\x18\x01 \x01(\x08\x12\x36\n\x0enat_interfaces\x18\x02 \x03(\x0b\x32\x1e.vpp.nat.Nat44Global.Interface\x12\x32\n\x0c\x61\x64\x64ress_pool\x18\x03 \x03(\x0b\x32\x1c.vpp.nat.Nat44Global.Address\x12\x36\n\x12virtual_reassembly\x18\x04 \x01(\x0b\x32\x1a.vpp.nat.VirtualReassembly\x1a\x44\n\tInterface\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\tis_inside\x18\x02 \x01(\x08\x12\x16\n\x0eoutput_feature\x18\x03 \x01(\x08\x1a=\n\x07\x41\x64\x64ress\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x12\x0e\n\x06vrf_id\x18\x02 \x01(\r\x12\x11\n\ttwice_nat\x18\x03 \x01(\x08\"\xd2\x05\n\x06\x44Nat44\x12\r\n\x05label\x18\x01 \x01(\t\x12\x32\n\x0bst_mappings\x18\x02 \x03(\x0b\x32\x1d.vpp.nat.DNat44.StaticMapping\x12\x34\n\x0bid_mappings\x18\x03 \x03(\x0b\x32\x1f.vpp.nat.DNat44.IdentityMapping\x1a\xa1\x03\n\rStaticMapping\x12\x1a\n\x12\x65xternal_interface\x18\x01 \x01(\t\x12\x13\n\x0b\x65xternal_ip\x18\x02 \x01(\t\x12\x15\n\rexternal_port\x18\x03 \x01(\r\x12\x38\n\tlocal_ips\x18\x04 \x03(\x0b\x32%.vpp.nat.DNat44.StaticMapping.LocalIP\x12*\n\x08protocol\x18\x05 \x01(\x0e\x32\x18.vpp.nat.DNat44.Protocol\x12=\n\ttwice_nat\x18\x06 \x01(\x0e\x32*.vpp.nat.DNat44.StaticMapping.TwiceNatMode\x12\x18\n\x10session_affinity\x18\x07 \x01(\r\x1aT\n\x07LocalIP\x12\x0e\n\x06vrf_id\x18\x01 \x01(\r\x12\x10\n\x08local_ip\x18\x02 \x01(\t\x12\x12\n\nlocal_port\x18\x03 \x01(\r\x12\x13\n\x0bprobability\x18\x04 \x01(\r\"3\n\x0cTwiceNatMode\x12\x0c\n\x08\x44ISABLED\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x08\n\x04SELF\x10\x02\x1a\x82\x01\n\x0fIdentityMapping\x12\x0e\n\x06vrf_id\x18\x01 \x01(\r\x12\x11\n\tinterface\x18\x02 \x01(\t\x12\x12\n\nip_address\x18\x03 \x01(\t\x12\x0c\n\x04port\x18\x04 \x01(\r\x12*\n\x08protocol\x18\x05 \x01(\x0e\x32\x18.vpp.nat.DNat44.Protocol\"&\n\x08Protocol\x12\x07\n\x03TCP\x10\x00\x12\x07\n\x03UDP\x10\x01\x12\x08\n\x04ICMP\x10\x02\"m\n\x11VirtualReassembly\x12\x0f\n\x07timeout\x18\x01 \x01(\r\x12\x18\n\x10max_reassemblies\x18\x02 \x01(\r\x12\x15\n\rmax_fragments\x18\x03 \x01(\r\x12\x16\n\x0e\x64rop_fragments\x18\x04 \x01(\x08\x42<Z6github.com/ligato/vpp-agent/api/models/vpp/nat;vpp_nat\xc8\xe3\x1e\x01\x62\x06proto3')
  ,
  dependencies=[github_dot_com_dot_gogo_dot_protobuf_dot_gogoproto_dot_gogo__pb2.DESCRIPTOR,])



_DNAT44_STATICMAPPING_TWICENATMODE = _descriptor.EnumDescriptor(
  name='TwiceNatMode',
  full_name='vpp.nat.DNat44.StaticMapping.TwiceNatMode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DISABLED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ENABLED', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SELF', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=916,
  serialized_end=967,
)
_sym_db.RegisterEnumDescriptor(_DNAT44_STATICMAPPING_TWICENATMODE)

_DNAT44_PROTOCOL = _descriptor.EnumDescriptor(
  name='Protocol',
  full_name='vpp.nat.DNat44.Protocol',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TCP', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UDP', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ICMP', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1102,
  serialized_end=1140,
)
_sym_db.RegisterEnumDescriptor(_DNAT44_PROTOCOL)


_NAT44GLOBAL_INTERFACE = _descriptor.Descriptor(
  name='Interface',
  full_name='vpp.nat.Nat44Global.Interface',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='vpp.nat.Nat44Global.Interface.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_inside', full_name='vpp.nat.Nat44Global.Interface.is_inside', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output_feature', full_name='vpp.nat.Nat44Global.Interface.output_feature', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=284,
  serialized_end=352,
)

_NAT44GLOBAL_ADDRESS = _descriptor.Descriptor(
  name='Address',
  full_name='vpp.nat.Nat44Global.Address',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='vpp.nat.Nat44Global.Address.address', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='vrf_id', full_name='vpp.nat.Nat44Global.Address.vrf_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='twice_nat', full_name='vpp.nat.Nat44Global.Address.twice_nat', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=354,
  serialized_end=415,
)

_NAT44GLOBAL = _descriptor.Descriptor(
  name='Nat44Global',
  full_name='vpp.nat.Nat44Global',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='forwarding', full_name='vpp.nat.Nat44Global.forwarding', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='nat_interfaces', full_name='vpp.nat.Nat44Global.nat_interfaces', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address_pool', full_name='vpp.nat.Nat44Global.address_pool', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='virtual_reassembly', full_name='vpp.nat.Nat44Global.virtual_reassembly', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_NAT44GLOBAL_INTERFACE, _NAT44GLOBAL_ADDRESS, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=85,
  serialized_end=415,
)


_DNAT44_STATICMAPPING_LOCALIP = _descriptor.Descriptor(
  name='LocalIP',
  full_name='vpp.nat.DNat44.StaticMapping.LocalIP',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='vrf_id', full_name='vpp.nat.DNat44.StaticMapping.LocalIP.vrf_id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='local_ip', full_name='vpp.nat.DNat44.StaticMapping.LocalIP.local_ip', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='local_port', full_name='vpp.nat.DNat44.StaticMapping.LocalIP.local_port', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='probability', full_name='vpp.nat.DNat44.StaticMapping.LocalIP.probability', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=830,
  serialized_end=914,
)

_DNAT44_STATICMAPPING = _descriptor.Descriptor(
  name='StaticMapping',
  full_name='vpp.nat.DNat44.StaticMapping',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='external_interface', full_name='vpp.nat.DNat44.StaticMapping.external_interface', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='external_ip', full_name='vpp.nat.DNat44.StaticMapping.external_ip', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='external_port', full_name='vpp.nat.DNat44.StaticMapping.external_port', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='local_ips', full_name='vpp.nat.DNat44.StaticMapping.local_ips', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='protocol', full_name='vpp.nat.DNat44.StaticMapping.protocol', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='twice_nat', full_name='vpp.nat.DNat44.StaticMapping.twice_nat', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='session_affinity', full_name='vpp.nat.DNat44.StaticMapping.session_affinity', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_DNAT44_STATICMAPPING_LOCALIP, ],
  enum_types=[
    _DNAT44_STATICMAPPING_TWICENATMODE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=550,
  serialized_end=967,
)

_DNAT44_IDENTITYMAPPING = _descriptor.Descriptor(
  name='IdentityMapping',
  full_name='vpp.nat.DNat44.IdentityMapping',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='vrf_id', full_name='vpp.nat.DNat44.IdentityMapping.vrf_id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='interface', full_name='vpp.nat.DNat44.IdentityMapping.interface', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ip_address', full_name='vpp.nat.DNat44.IdentityMapping.ip_address', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='port', full_name='vpp.nat.DNat44.IdentityMapping.port', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='protocol', full_name='vpp.nat.DNat44.IdentityMapping.protocol', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=970,
  serialized_end=1100,
)

_DNAT44 = _descriptor.Descriptor(
  name='DNat44',
  full_name='vpp.nat.DNat44',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='label', full_name='vpp.nat.DNat44.label', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='st_mappings', full_name='vpp.nat.DNat44.st_mappings', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='id_mappings', full_name='vpp.nat.DNat44.id_mappings', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_DNAT44_STATICMAPPING, _DNAT44_IDENTITYMAPPING, ],
  enum_types=[
    _DNAT44_PROTOCOL,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=418,
  serialized_end=1140,
)


_VIRTUALREASSEMBLY = _descriptor.Descriptor(
  name='VirtualReassembly',
  full_name='vpp.nat.VirtualReassembly',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='timeout', full_name='vpp.nat.VirtualReassembly.timeout', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_reassemblies', full_name='vpp.nat.VirtualReassembly.max_reassemblies', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_fragments', full_name='vpp.nat.VirtualReassembly.max_fragments', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='drop_fragments', full_name='vpp.nat.VirtualReassembly.drop_fragments', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1142,
  serialized_end=1251,
)

_NAT44GLOBAL_INTERFACE.containing_type = _NAT44GLOBAL
_NAT44GLOBAL_ADDRESS.containing_type = _NAT44GLOBAL
_NAT44GLOBAL.fields_by_name['nat_interfaces'].message_type = _NAT44GLOBAL_INTERFACE
_NAT44GLOBAL.fields_by_name['address_pool'].message_type = _NAT44GLOBAL_ADDRESS
_NAT44GLOBAL.fields_by_name['virtual_reassembly'].message_type = _VIRTUALREASSEMBLY
_DNAT44_STATICMAPPING_LOCALIP.containing_type = _DNAT44_STATICMAPPING
_DNAT44_STATICMAPPING.fields_by_name['local_ips'].message_type = _DNAT44_STATICMAPPING_LOCALIP
_DNAT44_STATICMAPPING.fields_by_name['protocol'].enum_type = _DNAT44_PROTOCOL
_DNAT44_STATICMAPPING.fields_by_name['twice_nat'].enum_type = _DNAT44_STATICMAPPING_TWICENATMODE
_DNAT44_STATICMAPPING.containing_type = _DNAT44
_DNAT44_STATICMAPPING_TWICENATMODE.containing_type = _DNAT44_STATICMAPPING
_DNAT44_IDENTITYMAPPING.fields_by_name['protocol'].enum_type = _DNAT44_PROTOCOL
_DNAT44_IDENTITYMAPPING.containing_type = _DNAT44
_DNAT44.fields_by_name['st_mappings'].message_type = _DNAT44_STATICMAPPING
_DNAT44.fields_by_name['id_mappings'].message_type = _DNAT44_IDENTITYMAPPING
_DNAT44_PROTOCOL.containing_type = _DNAT44
DESCRIPTOR.message_types_by_name['Nat44Global'] = _NAT44GLOBAL
DESCRIPTOR.message_types_by_name['DNat44'] = _DNAT44
DESCRIPTOR.message_types_by_name['VirtualReassembly'] = _VIRTUALREASSEMBLY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Nat44Global = _reflection.GeneratedProtocolMessageType('Nat44Global', (_message.Message,), dict(

  Interface = _reflection.GeneratedProtocolMessageType('Interface', (_message.Message,), dict(
    DESCRIPTOR = _NAT44GLOBAL_INTERFACE,
    __module__ = 'models.vpp.nat.nat_pb2'
    # @@protoc_insertion_point(class_scope:vpp.nat.Nat44Global.Interface)
    ))
  ,

  Address = _reflection.GeneratedProtocolMessageType('Address', (_message.Message,), dict(
    DESCRIPTOR = _NAT44GLOBAL_ADDRESS,
    __module__ = 'models.vpp.nat.nat_pb2'
    # @@protoc_insertion_point(class_scope:vpp.nat.Nat44Global.Address)
    ))
  ,
  DESCRIPTOR = _NAT44GLOBAL,
  __module__ = 'models.vpp.nat.nat_pb2'
  # @@protoc_insertion_point(class_scope:vpp.nat.Nat44Global)
  ))
_sym_db.RegisterMessage(Nat44Global)
_sym_db.RegisterMessage(Nat44Global.Interface)
_sym_db.RegisterMessage(Nat44Global.Address)

DNat44 = _reflection.GeneratedProtocolMessageType('DNat44', (_message.Message,), dict(

  StaticMapping = _reflection.GeneratedProtocolMessageType('StaticMapping', (_message.Message,), dict(

    LocalIP = _reflection.GeneratedProtocolMessageType('LocalIP', (_message.Message,), dict(
      DESCRIPTOR = _DNAT44_STATICMAPPING_LOCALIP,
      __module__ = 'models.vpp.nat.nat_pb2'
      # @@protoc_insertion_point(class_scope:vpp.nat.DNat44.StaticMapping.LocalIP)
      ))
    ,
    DESCRIPTOR = _DNAT44_STATICMAPPING,
    __module__ = 'models.vpp.nat.nat_pb2'
    # @@protoc_insertion_point(class_scope:vpp.nat.DNat44.StaticMapping)
    ))
  ,

  IdentityMapping = _reflection.GeneratedProtocolMessageType('IdentityMapping', (_message.Message,), dict(
    DESCRIPTOR = _DNAT44_IDENTITYMAPPING,
    __module__ = 'models.vpp.nat.nat_pb2'
    # @@protoc_insertion_point(class_scope:vpp.nat.DNat44.IdentityMapping)
    ))
  ,
  DESCRIPTOR = _DNAT44,
  __module__ = 'models.vpp.nat.nat_pb2'
  # @@protoc_insertion_point(class_scope:vpp.nat.DNat44)
  ))
_sym_db.RegisterMessage(DNat44)
_sym_db.RegisterMessage(DNat44.StaticMapping)
_sym_db.RegisterMessage(DNat44.StaticMapping.LocalIP)
_sym_db.RegisterMessage(DNat44.IdentityMapping)

VirtualReassembly = _reflection.GeneratedProtocolMessageType('VirtualReassembly', (_message.Message,), dict(
  DESCRIPTOR = _VIRTUALREASSEMBLY,
  __module__ = 'models.vpp.nat.nat_pb2'
  # @@protoc_insertion_point(class_scope:vpp.nat.VirtualReassembly)
  ))
_sym_db.RegisterMessage(VirtualReassembly)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
