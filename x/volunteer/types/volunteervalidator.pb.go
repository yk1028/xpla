// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: xpla/volunteer/v1beta1/volunteervalidator.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// VolunteerValidator required for validator set update logic.
type VolunteerValidator struct {
	// address is the address of the validator.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// power defines the power of the validator.
	Power int64 `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *VolunteerValidator) Reset()         { *m = VolunteerValidator{} }
func (m *VolunteerValidator) String() string { return proto.CompactTextString(m) }
func (*VolunteerValidator) ProtoMessage()    {}
func (*VolunteerValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_29985b0ee34b89e7, []int{0}
}
func (m *VolunteerValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VolunteerValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VolunteerValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VolunteerValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolunteerValidator.Merge(m, src)
}
func (m *VolunteerValidator) XXX_Size() int {
	return m.Size()
}
func (m *VolunteerValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_VolunteerValidator.DiscardUnknown(m)
}

var xxx_messageInfo_VolunteerValidator proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VolunteerValidator)(nil), "xpla.volunteer.v1beta1.VolunteerValidator")
}

func init() {
	proto.RegisterFile("xpla/volunteer/v1beta1/volunteervalidator.proto", fileDescriptor_29985b0ee34b89e7)
}

var fileDescriptor_29985b0ee34b89e7 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xaf, 0x28, 0xc8, 0x49,
	0xd4, 0x2f, 0xcb, 0xcf, 0x29, 0xcd, 0x2b, 0x49, 0x4d, 0x2d, 0xd2, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0x44, 0x88, 0x94, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xe4, 0x17, 0xe9, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0x89, 0x81, 0x34, 0xe8, 0xc1, 0xa5, 0xf5, 0xa0, 0x1a, 0xa4, 0x44, 0xd2,
	0xf3, 0xd3, 0xf3, 0xc1, 0x4a, 0xf4, 0x41, 0x2c, 0x88, 0x6a, 0x25, 0x3f, 0x2e, 0xa1, 0x30, 0x98,
	0xd2, 0x30, 0x98, 0x49, 0x42, 0x12, 0x5c, 0xec, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x41, 0x7e, 0x79, 0x6a,
	0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x84, 0x63, 0xc5, 0xd1, 0xb1, 0x40, 0x9e, 0xe1,
	0xc5, 0x02, 0x79, 0x06, 0x27, 0xe7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0,
	0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88,
	0xd2, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x05, 0xfb, 0x29, 0x25, 0xb5,
	0x0c, 0xe2, 0xb7, 0x0a, 0x24, 0xdf, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xdd, 0x66,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xaa, 0xd8, 0xf0, 0xfc, 0x00, 0x00, 0x00,
}

func (m *VolunteerValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VolunteerValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VolunteerValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintVolunteervalidator(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintVolunteervalidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVolunteervalidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovVolunteervalidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VolunteerValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovVolunteervalidator(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovVolunteervalidator(uint64(m.Power))
	}
	return n
}

func sovVolunteervalidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVolunteervalidator(x uint64) (n int) {
	return sovVolunteervalidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VolunteerValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVolunteervalidator
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
			return fmt.Errorf("proto: VolunteerValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VolunteerValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolunteervalidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVolunteervalidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVolunteervalidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVolunteervalidator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVolunteervalidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVolunteervalidator
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
func skipVolunteervalidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVolunteervalidator
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
					return 0, ErrIntOverflowVolunteervalidator
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
					return 0, ErrIntOverflowVolunteervalidator
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
				return 0, ErrInvalidLengthVolunteervalidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVolunteervalidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVolunteervalidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVolunteervalidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVolunteervalidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVolunteervalidator = fmt.Errorf("proto: unexpected end of group")
)
