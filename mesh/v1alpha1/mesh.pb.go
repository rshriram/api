// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh/v1alpha1/mesh.proto

package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Network provides information about the endpoints in a routable L3
// network. A single routable L3 network can have one or more service
// registries. Note that the network has no relation to the locality of the
// endpoint. The endpoint locality will be obtained from the service
// registry.
type Network struct {
	// A unique name assigned to the network.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of endpoints in the network (obtained through the constituent
	// service registries or from CIDR ranges). All endpoints in the network
	// are directly accessible to one another.
	Endpoints []*Network_NetworkEndpoints    `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
	Gateways  []*Network_IstioNetworkGateway `protobuf:"bytes,3,rep,name=gateways" json:"gateways,omitempty"`
}

func (m *Network) Reset()                    { *m = Network{} }
func (m *Network) String() string            { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()               {}
func (*Network) Descriptor() ([]byte, []int) { return fileDescriptorMesh, []int{0} }

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetEndpoints() []*Network_NetworkEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *Network) GetGateways() []*Network_IstioNetworkGateway {
	if m != nil {
		return m.Gateways
	}
	return nil
}

type Network_NetworkEndpoints struct {
	// Types that are valid to be assigned to Ne:
	//	*Network_NetworkEndpoints_FromCidr
	//	*Network_NetworkEndpoints_FromRegistry
	Ne isNetwork_NetworkEndpoints_Ne `protobuf_oneof:"ne"`
}

func (m *Network_NetworkEndpoints) Reset()                    { *m = Network_NetworkEndpoints{} }
func (m *Network_NetworkEndpoints) String() string            { return proto.CompactTextString(m) }
func (*Network_NetworkEndpoints) ProtoMessage()               {}
func (*Network_NetworkEndpoints) Descriptor() ([]byte, []int) { return fileDescriptorMesh, []int{0, 0} }

type isNetwork_NetworkEndpoints_Ne interface {
	isNetwork_NetworkEndpoints_Ne()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Network_NetworkEndpoints_FromCidr struct {
	FromCidr string `protobuf:"bytes,1,opt,name=from_cidr,json=fromCidr,proto3,oneof"`
}
type Network_NetworkEndpoints_FromRegistry struct {
	FromRegistry string `protobuf:"bytes,2,opt,name=from_registry,json=fromRegistry,proto3,oneof"`
}

func (*Network_NetworkEndpoints_FromCidr) isNetwork_NetworkEndpoints_Ne()     {}
func (*Network_NetworkEndpoints_FromRegistry) isNetwork_NetworkEndpoints_Ne() {}

func (m *Network_NetworkEndpoints) GetNe() isNetwork_NetworkEndpoints_Ne {
	if m != nil {
		return m.Ne
	}
	return nil
}

func (m *Network_NetworkEndpoints) GetFromCidr() string {
	if x, ok := m.GetNe().(*Network_NetworkEndpoints_FromCidr); ok {
		return x.FromCidr
	}
	return ""
}

func (m *Network_NetworkEndpoints) GetFromRegistry() string {
	if x, ok := m.GetNe().(*Network_NetworkEndpoints_FromRegistry); ok {
		return x.FromRegistry
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Network_NetworkEndpoints) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Network_NetworkEndpoints_OneofMarshaler, _Network_NetworkEndpoints_OneofUnmarshaler, _Network_NetworkEndpoints_OneofSizer, []interface{}{
		(*Network_NetworkEndpoints_FromCidr)(nil),
		(*Network_NetworkEndpoints_FromRegistry)(nil),
	}
}

func _Network_NetworkEndpoints_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Network_NetworkEndpoints)
	// ne
	switch x := m.Ne.(type) {
	case *Network_NetworkEndpoints_FromCidr:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.FromCidr)
	case *Network_NetworkEndpoints_FromRegistry:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.FromRegistry)
	case nil:
	default:
		return fmt.Errorf("Network_NetworkEndpoints.Ne has unexpected type %T", x)
	}
	return nil
}

func _Network_NetworkEndpoints_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Network_NetworkEndpoints)
	switch tag {
	case 1: // ne.from_cidr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Ne = &Network_NetworkEndpoints_FromCidr{x}
		return true, err
	case 2: // ne.from_registry
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Ne = &Network_NetworkEndpoints_FromRegistry{x}
		return true, err
	default:
		return false, nil
	}
}

func _Network_NetworkEndpoints_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Network_NetworkEndpoints)
	// ne
	switch x := m.Ne.(type) {
	case *Network_NetworkEndpoints_FromCidr:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FromCidr)))
		n += len(x.FromCidr)
	case *Network_NetworkEndpoints_FromRegistry:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.FromRegistry)))
		n += len(x.FromRegistry)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The gateway associated with this network. Traffic from remote networks
// will arrive at the specified gateway:port. All incoming traffic must
// use mTLS.
type Network_IstioNetworkGateway struct {
	// Types that are valid to be assigned to Gw:
	//	*Network_IstioNetworkGateway_RegistryServiceName
	//	*Network_IstioNetworkGateway_Address
	Gw isNetwork_IstioNetworkGateway_Gw `protobuf_oneof:"gw"`
	// The port associated with the gateway.
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// The locality associated with an explicitly specified gateway (i.e. ip)
	Locality string `protobuf:"bytes,4,opt,name=locality,proto3" json:"locality,omitempty"`
}

func (m *Network_IstioNetworkGateway) Reset()         { *m = Network_IstioNetworkGateway{} }
func (m *Network_IstioNetworkGateway) String() string { return proto.CompactTextString(m) }
func (*Network_IstioNetworkGateway) ProtoMessage()    {}
func (*Network_IstioNetworkGateway) Descriptor() ([]byte, []int) {
	return fileDescriptorMesh, []int{0, 1}
}

type isNetwork_IstioNetworkGateway_Gw interface {
	isNetwork_IstioNetworkGateway_Gw()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Network_IstioNetworkGateway_RegistryServiceName struct {
	RegistryServiceName string `protobuf:"bytes,1,opt,name=registry_service_name,json=registryServiceName,proto3,oneof"`
}
type Network_IstioNetworkGateway_Address struct {
	Address string `protobuf:"bytes,2,opt,name=address,proto3,oneof"`
}

func (*Network_IstioNetworkGateway_RegistryServiceName) isNetwork_IstioNetworkGateway_Gw() {}
func (*Network_IstioNetworkGateway_Address) isNetwork_IstioNetworkGateway_Gw()             {}

func (m *Network_IstioNetworkGateway) GetGw() isNetwork_IstioNetworkGateway_Gw {
	if m != nil {
		return m.Gw
	}
	return nil
}

func (m *Network_IstioNetworkGateway) GetRegistryServiceName() string {
	if x, ok := m.GetGw().(*Network_IstioNetworkGateway_RegistryServiceName); ok {
		return x.RegistryServiceName
	}
	return ""
}

func (m *Network_IstioNetworkGateway) GetAddress() string {
	if x, ok := m.GetGw().(*Network_IstioNetworkGateway_Address); ok {
		return x.Address
	}
	return ""
}

func (m *Network_IstioNetworkGateway) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Network_IstioNetworkGateway) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Network_IstioNetworkGateway) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Network_IstioNetworkGateway_OneofMarshaler, _Network_IstioNetworkGateway_OneofUnmarshaler, _Network_IstioNetworkGateway_OneofSizer, []interface{}{
		(*Network_IstioNetworkGateway_RegistryServiceName)(nil),
		(*Network_IstioNetworkGateway_Address)(nil),
	}
}

func _Network_IstioNetworkGateway_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Network_IstioNetworkGateway)
	// gw
	switch x := m.Gw.(type) {
	case *Network_IstioNetworkGateway_RegistryServiceName:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.RegistryServiceName)
	case *Network_IstioNetworkGateway_Address:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Address)
	case nil:
	default:
		return fmt.Errorf("Network_IstioNetworkGateway.Gw has unexpected type %T", x)
	}
	return nil
}

func _Network_IstioNetworkGateway_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Network_IstioNetworkGateway)
	switch tag {
	case 1: // gw.registry_service_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Gw = &Network_IstioNetworkGateway_RegistryServiceName{x}
		return true, err
	case 2: // gw.address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Gw = &Network_IstioNetworkGateway_Address{x}
		return true, err
	default:
		return false, nil
	}
}

func _Network_IstioNetworkGateway_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Network_IstioNetworkGateway)
	// gw
	switch x := m.Gw.(type) {
	case *Network_IstioNetworkGateway_RegistryServiceName:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RegistryServiceName)))
		n += len(x.RegistryServiceName)
	case *Network_IstioNetworkGateway_Address:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Address)))
		n += len(x.Address)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Network)(nil), "istio.mesh.v1alpha1.Network")
	proto.RegisterType((*Network_NetworkEndpoints)(nil), "istio.mesh.v1alpha1.Network.NetworkEndpoints")
	proto.RegisterType((*Network_IstioNetworkGateway)(nil), "istio.mesh.v1alpha1.Network.IstioNetworkGateway")
}
func (m *Network) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Network) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMesh(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Endpoints) > 0 {
		for _, msg := range m.Endpoints {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMesh(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Gateways) > 0 {
		for _, msg := range m.Gateways {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintMesh(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Network_NetworkEndpoints) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Network_NetworkEndpoints) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Ne != nil {
		nn1, err := m.Ne.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *Network_NetworkEndpoints_FromCidr) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintMesh(dAtA, i, uint64(len(m.FromCidr)))
	i += copy(dAtA[i:], m.FromCidr)
	return i, nil
}
func (m *Network_NetworkEndpoints_FromRegistry) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintMesh(dAtA, i, uint64(len(m.FromRegistry)))
	i += copy(dAtA[i:], m.FromRegistry)
	return i, nil
}
func (m *Network_IstioNetworkGateway) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Network_IstioNetworkGateway) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Gw != nil {
		nn2, err := m.Gw.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn2
	}
	if m.Port != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMesh(dAtA, i, uint64(m.Port))
	}
	if len(m.Locality) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMesh(dAtA, i, uint64(len(m.Locality)))
		i += copy(dAtA[i:], m.Locality)
	}
	return i, nil
}

func (m *Network_IstioNetworkGateway_RegistryServiceName) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintMesh(dAtA, i, uint64(len(m.RegistryServiceName)))
	i += copy(dAtA[i:], m.RegistryServiceName)
	return i, nil
}
func (m *Network_IstioNetworkGateway_Address) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintMesh(dAtA, i, uint64(len(m.Address)))
	i += copy(dAtA[i:], m.Address)
	return i, nil
}
func encodeVarintMesh(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Network) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMesh(uint64(l))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovMesh(uint64(l))
		}
	}
	if len(m.Gateways) > 0 {
		for _, e := range m.Gateways {
			l = e.Size()
			n += 1 + l + sovMesh(uint64(l))
		}
	}
	return n
}

func (m *Network_NetworkEndpoints) Size() (n int) {
	var l int
	_ = l
	if m.Ne != nil {
		n += m.Ne.Size()
	}
	return n
}

func (m *Network_NetworkEndpoints_FromCidr) Size() (n int) {
	var l int
	_ = l
	l = len(m.FromCidr)
	n += 1 + l + sovMesh(uint64(l))
	return n
}
func (m *Network_NetworkEndpoints_FromRegistry) Size() (n int) {
	var l int
	_ = l
	l = len(m.FromRegistry)
	n += 1 + l + sovMesh(uint64(l))
	return n
}
func (m *Network_IstioNetworkGateway) Size() (n int) {
	var l int
	_ = l
	if m.Gw != nil {
		n += m.Gw.Size()
	}
	if m.Port != 0 {
		n += 1 + sovMesh(uint64(m.Port))
	}
	l = len(m.Locality)
	if l > 0 {
		n += 1 + l + sovMesh(uint64(l))
	}
	return n
}

func (m *Network_IstioNetworkGateway_RegistryServiceName) Size() (n int) {
	var l int
	_ = l
	l = len(m.RegistryServiceName)
	n += 1 + l + sovMesh(uint64(l))
	return n
}
func (m *Network_IstioNetworkGateway_Address) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	n += 1 + l + sovMesh(uint64(l))
	return n
}

func sovMesh(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMesh(x uint64) (n int) {
	return sovMesh(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Network) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Network: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Network: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, &Network_NetworkEndpoints{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gateways", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gateways = append(m.Gateways, &Network_IstioNetworkGateway{})
			if err := m.Gateways[len(m.Gateways)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Network_NetworkEndpoints) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkEndpoints: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkEndpoints: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromCidr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ne = &Network_NetworkEndpoints_FromCidr{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromRegistry", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ne = &Network_NetworkEndpoints_FromRegistry{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Network_IstioNetworkGateway) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IstioNetworkGateway: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioNetworkGateway: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistryServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gw = &Network_IstioNetworkGateway_RegistryServiceName{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gw = &Network_IstioNetworkGateway_Address{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locality", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMesh
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locality = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMesh
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMesh(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMesh
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMesh
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMesh
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMesh
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMesh(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMesh = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMesh   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mesh/v1alpha1/mesh.proto", fileDescriptorMesh) }

var fileDescriptorMesh = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x5d, 0xd6, 0xe1, 0xb6, 0xab, 0x03, 0xc9, 0x10, 0x42, 0xc1, 0x31, 0x04, 0xa1, 0x2f, 0xb6,
	0x4e, 0xfd, 0x82, 0x89, 0x38, 0x51, 0xf6, 0x10, 0xdf, 0x7c, 0x70, 0xc4, 0x35, 0x6e, 0xc1, 0xad,
	0x29, 0x49, 0xd8, 0xd8, 0xc7, 0xf8, 0x1f, 0x7e, 0x82, 0x8f, 0x7e, 0x82, 0xec, 0x4b, 0x24, 0x69,
	0xb3, 0xa1, 0x0c, 0x9f, 0x7a, 0xef, 0x39, 0x3d, 0xe7, 0xde, 0x9e, 0x5e, 0x20, 0x73, 0xae, 0xa7,
	0xc9, 0xa2, 0xc7, 0x66, 0xf9, 0x94, 0xf5, 0x12, 0xdb, 0xc5, 0xb9, 0x92, 0x46, 0xe2, 0xb6, 0xd0,
	0x46, 0xc8, 0xd8, 0x21, 0x9e, 0x3f, 0xf9, 0x08, 0xa0, 0x3e, 0xe4, 0x66, 0x29, 0xd5, 0x1b, 0xc6,
	0x50, 0xcb, 0xd8, 0x9c, 0x13, 0xd4, 0x45, 0x51, 0x93, 0xba, 0x1a, 0xdf, 0x43, 0x93, 0x67, 0x69,
	0x2e, 0x45, 0x66, 0x34, 0xa9, 0x76, 0x83, 0x68, 0xff, 0xe2, 0x2c, 0xde, 0x61, 0x14, 0x97, 0x26,
	0xfe, 0x79, 0xe3, 0x45, 0x74, 0xab, 0xc7, 0x0f, 0xd0, 0x98, 0x30, 0xc3, 0x97, 0x6c, 0xa5, 0x49,
	0xe0, 0xbc, 0xce, 0xff, 0xf5, 0xba, 0xb3, 0x5c, 0xd9, 0xdc, 0x16, 0x42, 0xba, 0x71, 0x08, 0x9f,
	0xe1, 0xf0, 0xef, 0x30, 0x7c, 0x0c, 0xcd, 0x57, 0x25, 0xe7, 0xa3, 0xb1, 0x48, 0x55, 0xf1, 0x1d,
	0x83, 0x0a, 0x6d, 0x58, 0xe8, 0x5a, 0xa4, 0x0a, 0x9f, 0x42, 0xcb, 0xd1, 0x8a, 0x4f, 0x84, 0x36,
	0x6a, 0x45, 0xaa, 0xe5, 0x2b, 0x07, 0x16, 0xa6, 0x25, 0xda, 0xaf, 0x41, 0x35, 0xe3, 0xe1, 0x3b,
	0x82, 0xf6, 0x8e, 0x0d, 0xf0, 0x15, 0x1c, 0x79, 0xfd, 0x48, 0x73, 0xb5, 0x10, 0x63, 0x3e, 0xda,
	0xe6, 0x36, 0xa8, 0xd0, 0xb6, 0xa7, 0x1f, 0x0b, 0x76, 0x68, 0x83, 0x0c, 0xa1, 0xce, 0xd2, 0x54,
	0x71, 0xad, 0x37, 0x43, 0x3d, 0x60, 0x83, 0xcf, 0xa5, 0x32, 0x24, 0xe8, 0xa2, 0xa8, 0x45, 0x5d,
	0x8d, 0x43, 0x68, 0xcc, 0xe4, 0x98, 0xcd, 0x84, 0x59, 0x91, 0x9a, 0xfb, 0x21, 0x9b, 0xde, 0xee,
	0x37, 0x59, 0xf6, 0xa3, 0xcf, 0x75, 0x07, 0x7d, 0xad, 0x3b, 0xe8, 0x7b, 0xdd, 0x41, 0x4f, 0x61,
	0x11, 0xa4, 0x90, 0x09, 0xcb, 0x45, 0xf2, 0xeb, 0x08, 0x5e, 0xf6, 0xdc, 0x01, 0x5c, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x5a, 0x29, 0x8a, 0x44, 0x1c, 0x02, 0x00, 0x00,
}
