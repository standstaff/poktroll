// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/supplier/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/pokt-network/poktroll/x/shared/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EventSupplierStaked is emitted with the commitment of the supplier stake message.
type EventSupplierStaked struct {
	Supplier *types.Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier"`
}

func (m *EventSupplierStaked) Reset()         { *m = EventSupplierStaked{} }
func (m *EventSupplierStaked) String() string { return proto.CompactTextString(m) }
func (*EventSupplierStaked) ProtoMessage()    {}
func (*EventSupplierStaked) Descriptor() ([]byte, []int) {
	return fileDescriptor_22d2d1a82853ce0a, []int{0}
}
func (m *EventSupplierStaked) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSupplierStaked) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EventSupplierStaked) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSupplierStaked.Merge(m, src)
}
func (m *EventSupplierStaked) XXX_Size() int {
	return m.Size()
}
func (m *EventSupplierStaked) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSupplierStaked.DiscardUnknown(m)
}

var xxx_messageInfo_EventSupplierStaked proto.InternalMessageInfo

func (m *EventSupplierStaked) GetSupplier() *types.Supplier {
	if m != nil {
		return m.Supplier
	}
	return nil
}

// EventSupplierUnstaked is emitted with the commitment of the supplier unstake message.
type EventSupplierUnbondingBegin struct {
	Supplier        *types.Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier"`
	UnbondingHeight int64           `protobuf:"varint,2,opt,name=unbonding_height,json=unbondingHeight,proto3" json:"unbonding_height"`
}

func (m *EventSupplierUnbondingBegin) Reset()         { *m = EventSupplierUnbondingBegin{} }
func (m *EventSupplierUnbondingBegin) String() string { return proto.CompactTextString(m) }
func (*EventSupplierUnbondingBegin) ProtoMessage()    {}
func (*EventSupplierUnbondingBegin) Descriptor() ([]byte, []int) {
	return fileDescriptor_22d2d1a82853ce0a, []int{1}
}
func (m *EventSupplierUnbondingBegin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSupplierUnbondingBegin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EventSupplierUnbondingBegin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSupplierUnbondingBegin.Merge(m, src)
}
func (m *EventSupplierUnbondingBegin) XXX_Size() int {
	return m.Size()
}
func (m *EventSupplierUnbondingBegin) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSupplierUnbondingBegin.DiscardUnknown(m)
}

var xxx_messageInfo_EventSupplierUnbondingBegin proto.InternalMessageInfo

func (m *EventSupplierUnbondingBegin) GetSupplier() *types.Supplier {
	if m != nil {
		return m.Supplier
	}
	return nil
}

func (m *EventSupplierUnbondingBegin) GetUnbondingHeight() int64 {
	if m != nil {
		return m.UnbondingHeight
	}
	return 0
}

// EventSupplierUnbondingEnd is emitted with the commitment of last block of the
// supplier unbonding period.
type EventSupplierUnbondingEnd struct {
	Supplier        *types.Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier"`
	UnbondingHeight int64           `protobuf:"varint,2,opt,name=unbonding_height,json=unbondingHeight,proto3" json:"unbonding_height"`
}

func (m *EventSupplierUnbondingEnd) Reset()         { *m = EventSupplierUnbondingEnd{} }
func (m *EventSupplierUnbondingEnd) String() string { return proto.CompactTextString(m) }
func (*EventSupplierUnbondingEnd) ProtoMessage()    {}
func (*EventSupplierUnbondingEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_22d2d1a82853ce0a, []int{2}
}
func (m *EventSupplierUnbondingEnd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSupplierUnbondingEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *EventSupplierUnbondingEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSupplierUnbondingEnd.Merge(m, src)
}
func (m *EventSupplierUnbondingEnd) XXX_Size() int {
	return m.Size()
}
func (m *EventSupplierUnbondingEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSupplierUnbondingEnd.DiscardUnknown(m)
}

var xxx_messageInfo_EventSupplierUnbondingEnd proto.InternalMessageInfo

func (m *EventSupplierUnbondingEnd) GetSupplier() *types.Supplier {
	if m != nil {
		return m.Supplier
	}
	return nil
}

func (m *EventSupplierUnbondingEnd) GetUnbondingHeight() int64 {
	if m != nil {
		return m.UnbondingHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*EventSupplierStaked)(nil), "poktroll.supplier.EventSupplierStaked")
	proto.RegisterType((*EventSupplierUnbondingBegin)(nil), "poktroll.supplier.EventSupplierUnbondingBegin")
	proto.RegisterType((*EventSupplierUnbondingEnd)(nil), "poktroll.supplier.EventSupplierUnbondingEnd")
}

func init() { proto.RegisterFile("poktroll/supplier/event.proto", fileDescriptor_22d2d1a82853ce0a) }

var fileDescriptor_22d2d1a82853ce0a = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xc8, 0xcf, 0x2e,
	0x29, 0xca, 0xcf, 0xc9, 0xd1, 0x2f, 0x2e, 0x2d, 0x28, 0xc8, 0xc9, 0x4c, 0x2d, 0xd2, 0x4f, 0x2d,
	0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x49, 0xeb, 0xc1, 0xa4,
	0xa5, 0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0xc1, 0x0a, 0xf4, 0x21, 0x1c, 0x88, 0x6a,
	0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x88, 0x38, 0x88, 0x05, 0x15, 0x95, 0x43, 0x58, 0x91, 0x91,
	0x58, 0x94, 0x9a, 0x02, 0xb7, 0x09, 0x22, 0xaf, 0x14, 0xc5, 0x25, 0xec, 0x0a, 0xb2, 0x32, 0x18,
	0x2a, 0x1c, 0x5c, 0x92, 0x98, 0x9d, 0x9a, 0x22, 0xe4, 0xcc, 0xc5, 0x01, 0x53, 0x28, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x87, 0x70, 0x0d, 0xd8, 0x24, 0x3d, 0x98, 0x16, 0x27, 0x9e,
	0x57, 0xf7, 0xe4, 0xe1, 0xca, 0x83, 0xe0, 0x2c, 0xa5, 0xc5, 0x8c, 0x5c, 0xd2, 0x28, 0x86, 0x87,
	0xe6, 0x25, 0xe5, 0xe7, 0xa5, 0x64, 0xe6, 0xa5, 0x3b, 0xa5, 0xa6, 0x67, 0xe6, 0x51, 0xc5, 0x12,
	0x21, 0x7b, 0x2e, 0x81, 0x52, 0x98, 0xb1, 0xf1, 0x19, 0xa9, 0x99, 0xe9, 0x19, 0x25, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0xcc, 0x4e, 0x22, 0xaf, 0xee, 0xc9, 0x63, 0xc8, 0x05, 0xf1, 0xc3, 0x45, 0x3c,
	0xc0, 0x02, 0x4a, 0x0b, 0x19, 0xb9, 0x24, 0xb1, 0xbb, 0xd2, 0x35, 0x2f, 0x65, 0x70, 0xb8, 0xd1,
	0xc9, 0xff, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x6f, 0x3c, 0x92, 0x63, 0x7c, 0xf0,
	0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88,
	0x32, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0xb9, 0x4d, 0x37,
	0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x5b, 0x1f, 0x1e, 0xf7, 0x15, 0x88, 0x04, 0x56, 0x52, 0x59,
	0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x8e, 0x7d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x42,
	0x78, 0xcc, 0x82, 0x02, 0x00, 0x00,
}

func (m *EventSupplierStaked) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSupplierStaked) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSupplierStaked) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Supplier != nil {
		{
			size, err := m.Supplier.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSupplierUnbondingBegin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSupplierUnbondingBegin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSupplierUnbondingBegin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnbondingHeight != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.UnbondingHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Supplier != nil {
		{
			size, err := m.Supplier.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSupplierUnbondingEnd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSupplierUnbondingEnd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSupplierUnbondingEnd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnbondingHeight != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.UnbondingHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Supplier != nil {
		{
			size, err := m.Supplier.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventSupplierStaked) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Supplier != nil {
		l = m.Supplier.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventSupplierUnbondingBegin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Supplier != nil {
		l = m.Supplier.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.UnbondingHeight != 0 {
		n += 1 + sovEvent(uint64(m.UnbondingHeight))
	}
	return n
}

func (m *EventSupplierUnbondingEnd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Supplier != nil {
		l = m.Supplier.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.UnbondingHeight != 0 {
		n += 1 + sovEvent(uint64(m.UnbondingHeight))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventSupplierStaked) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventSupplierStaked: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSupplierStaked: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supplier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Supplier == nil {
				m.Supplier = &types.Supplier{}
			}
			if err := m.Supplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventSupplierUnbondingBegin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventSupplierUnbondingBegin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSupplierUnbondingBegin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supplier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Supplier == nil {
				m.Supplier = &types.Supplier{}
			}
			if err := m.Supplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingHeight", wireType)
			}
			m.UnbondingHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventSupplierUnbondingEnd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventSupplierUnbondingEnd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSupplierUnbondingEnd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supplier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Supplier == nil {
				m.Supplier = &types.Supplier{}
			}
			if err := m.Supplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingHeight", wireType)
			}
			m.UnbondingHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
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
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)