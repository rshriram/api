# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: networking/v1alpha3/service_entry.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from networking.v1alpha3 import gateway_pb2 as networking_dot_v1alpha3_dot_gateway__pb2
from networking.v1alpha3 import network_scope_pb2 as networking_dot_v1alpha3_dot_network__scope__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='networking/v1alpha3/service_entry.proto',
  package='istio.networking.v1alpha3',
  syntax='proto3',
  serialized_pb=_b('\n\'networking/v1alpha3/service_entry.proto\x12\x19istio.networking.v1alpha3\x1a!networking/v1alpha3/gateway.proto\x1a\'networking/v1alpha3/network_scope.proto\"\x96\x06\n\x0cServiceEntry\x12\r\n\x05hosts\x18\x01 \x03(\t\x12\x11\n\taddresses\x18\x02 \x03(\t\x12.\n\x05ports\x18\x03 \x03(\x0b\x32\x1f.istio.networking.v1alpha3.Port\x12\x42\n\x08location\x18\x04 \x01(\x0e\x32\x30.istio.networking.v1alpha3.ServiceEntry.Location\x12\x46\n\nresolution\x18\x05 \x01(\x0e\x32\x32.istio.networking.v1alpha3.ServiceEntry.Resolution\x12\x43\n\tendpoints\x18\x06 \x03(\x0b\x32\x30.istio.networking.v1alpha3.ServiceEntry.Endpoint\x12<\n\x0c\x63onfig_scope\x18\x07 \x01(\x0e\x32&.istio.networking.v1alpha3.ConfigScope\x1a\xc5\x02\n\x08\x45ndpoint\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x12J\n\x05ports\x18\x02 \x03(\x0b\x32;.istio.networking.v1alpha3.ServiceEntry.Endpoint.PortsEntry\x12L\n\x06labels\x18\x03 \x03(\x0b\x32<.istio.networking.v1alpha3.ServiceEntry.Endpoint.LabelsEntry\x12\x0f\n\x07network\x18\x04 \x01(\t\x12\x10\n\x08locality\x18\x05 \x01(\t\x12\x0e\n\x06weight\x18\x06 \x01(\r\x1a,\n\nPortsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\r:\x02\x38\x01\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"0\n\x08Location\x12\x11\n\rMESH_EXTERNAL\x10\x00\x12\x11\n\rMESH_INTERNAL\x10\x01\"+\n\nResolution\x12\x08\n\x04NONE\x10\x00\x12\n\n\x06STATIC\x10\x01\x12\x07\n\x03\x44NS\x10\x02\x42\"Z istio.io/api/networking/v1alpha3b\x06proto3')
  ,
  dependencies=[networking_dot_v1alpha3_dot_gateway__pb2.DESCRIPTOR,networking_dot_v1alpha3_dot_network__scope__pb2.DESCRIPTOR,])



_SERVICEENTRY_LOCATION = _descriptor.EnumDescriptor(
  name='Location',
  full_name='istio.networking.v1alpha3.ServiceEntry.Location',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='MESH_EXTERNAL', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MESH_INTERNAL', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=844,
  serialized_end=892,
)
_sym_db.RegisterEnumDescriptor(_SERVICEENTRY_LOCATION)

_SERVICEENTRY_RESOLUTION = _descriptor.EnumDescriptor(
  name='Resolution',
  full_name='istio.networking.v1alpha3.ServiceEntry.Resolution',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STATIC', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DNS', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=894,
  serialized_end=937,
)
_sym_db.RegisterEnumDescriptor(_SERVICEENTRY_RESOLUTION)


_SERVICEENTRY_ENDPOINT_PORTSENTRY = _descriptor.Descriptor(
  name='PortsEntry',
  full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.PortsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.PortsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.PortsEntry.value', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=751,
  serialized_end=795,
)

_SERVICEENTRY_ENDPOINT_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.LabelsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=797,
  serialized_end=842,
)

_SERVICEENTRY_ENDPOINT = _descriptor.Descriptor(
  name='Endpoint',
  full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.address', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ports', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.ports', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='labels', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.labels', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='network', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.network', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='locality', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.locality', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='weight', full_name='istio.networking.v1alpha3.ServiceEntry.Endpoint.weight', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SERVICEENTRY_ENDPOINT_PORTSENTRY, _SERVICEENTRY_ENDPOINT_LABELSENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=517,
  serialized_end=842,
)

_SERVICEENTRY = _descriptor.Descriptor(
  name='ServiceEntry',
  full_name='istio.networking.v1alpha3.ServiceEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hosts', full_name='istio.networking.v1alpha3.ServiceEntry.hosts', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addresses', full_name='istio.networking.v1alpha3.ServiceEntry.addresses', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ports', full_name='istio.networking.v1alpha3.ServiceEntry.ports', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location', full_name='istio.networking.v1alpha3.ServiceEntry.location', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resolution', full_name='istio.networking.v1alpha3.ServiceEntry.resolution', index=4,
      number=5, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='endpoints', full_name='istio.networking.v1alpha3.ServiceEntry.endpoints', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='config_scope', full_name='istio.networking.v1alpha3.ServiceEntry.config_scope', index=6,
      number=7, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_SERVICEENTRY_ENDPOINT, ],
  enum_types=[
    _SERVICEENTRY_LOCATION,
    _SERVICEENTRY_RESOLUTION,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=147,
  serialized_end=937,
)

_SERVICEENTRY_ENDPOINT_PORTSENTRY.containing_type = _SERVICEENTRY_ENDPOINT
_SERVICEENTRY_ENDPOINT_LABELSENTRY.containing_type = _SERVICEENTRY_ENDPOINT
_SERVICEENTRY_ENDPOINT.fields_by_name['ports'].message_type = _SERVICEENTRY_ENDPOINT_PORTSENTRY
_SERVICEENTRY_ENDPOINT.fields_by_name['labels'].message_type = _SERVICEENTRY_ENDPOINT_LABELSENTRY
_SERVICEENTRY_ENDPOINT.containing_type = _SERVICEENTRY
_SERVICEENTRY.fields_by_name['ports'].message_type = networking_dot_v1alpha3_dot_gateway__pb2._PORT
_SERVICEENTRY.fields_by_name['location'].enum_type = _SERVICEENTRY_LOCATION
_SERVICEENTRY.fields_by_name['resolution'].enum_type = _SERVICEENTRY_RESOLUTION
_SERVICEENTRY.fields_by_name['endpoints'].message_type = _SERVICEENTRY_ENDPOINT
_SERVICEENTRY.fields_by_name['config_scope'].enum_type = networking_dot_v1alpha3_dot_network__scope__pb2._CONFIGSCOPE
_SERVICEENTRY_LOCATION.containing_type = _SERVICEENTRY
_SERVICEENTRY_RESOLUTION.containing_type = _SERVICEENTRY
DESCRIPTOR.message_types_by_name['ServiceEntry'] = _SERVICEENTRY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ServiceEntry = _reflection.GeneratedProtocolMessageType('ServiceEntry', (_message.Message,), dict(

  Endpoint = _reflection.GeneratedProtocolMessageType('Endpoint', (_message.Message,), dict(

    PortsEntry = _reflection.GeneratedProtocolMessageType('PortsEntry', (_message.Message,), dict(
      DESCRIPTOR = _SERVICEENTRY_ENDPOINT_PORTSENTRY,
      __module__ = 'networking.v1alpha3.service_entry_pb2'
      # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.ServiceEntry.Endpoint.PortsEntry)
      ))
    ,

    LabelsEntry = _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), dict(
      DESCRIPTOR = _SERVICEENTRY_ENDPOINT_LABELSENTRY,
      __module__ = 'networking.v1alpha3.service_entry_pb2'
      # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.ServiceEntry.Endpoint.LabelsEntry)
      ))
    ,
    DESCRIPTOR = _SERVICEENTRY_ENDPOINT,
    __module__ = 'networking.v1alpha3.service_entry_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.ServiceEntry.Endpoint)
    ))
  ,
  DESCRIPTOR = _SERVICEENTRY,
  __module__ = 'networking.v1alpha3.service_entry_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1alpha3.ServiceEntry)
  ))
_sym_db.RegisterMessage(ServiceEntry)
_sym_db.RegisterMessage(ServiceEntry.Endpoint)
_sym_db.RegisterMessage(ServiceEntry.Endpoint.PortsEntry)
_sym_db.RegisterMessage(ServiceEntry.Endpoint.LabelsEntry)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z istio.io/api/networking/v1alpha3'))
_SERVICEENTRY_ENDPOINT_PORTSENTRY.has_options = True
_SERVICEENTRY_ENDPOINT_PORTSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
_SERVICEENTRY_ENDPOINT_LABELSENTRY.has_options = True
_SERVICEENTRY_ENDPOINT_LABELSENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)
