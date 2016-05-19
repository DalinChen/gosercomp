// Code generated by protoc-gen-gogo.
// source: mygogo.proto
// DO NOT EDIT!

/*
	Package gosercomp is a generated protocol buffer package.

	It is generated from these files:
		mygogo.proto

	It has these top-level messages:
		GogoProtoColorGroup
*/
package gosercomp

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

type GogoProtoColorGroup struct {
	Id     int32    `protobuf:"varint,1,req,name=id" json:"id"`
	Name   string   `protobuf:"bytes,2,req,name=name" json:"name"`
	Colors []string `protobuf:"bytes,3,rep,name=colors" json:"colors,omitempty"`
}

func (m *GogoProtoColorGroup) Reset()                    { *m = GogoProtoColorGroup{} }
func (m *GogoProtoColorGroup) String() string            { return proto.CompactTextString(m) }
func (*GogoProtoColorGroup) ProtoMessage()               {}
func (*GogoProtoColorGroup) Descriptor() ([]byte, []int) { return fileDescriptorMygogo, []int{0} }

func (m *GogoProtoColorGroup) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GogoProtoColorGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GogoProtoColorGroup) GetColors() []string {
	if m != nil {
		return m.Colors
	}
	return nil
}

func init() {
	proto.RegisterType((*GogoProtoColorGroup)(nil), "gosercomp.GogoProtoColorGroup")
}
func (m *GogoProtoColorGroup) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *GogoProtoColorGroup) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintMygogo(data, i, uint64(m.Id))
	data[i] = 0x12
	i++
	i = encodeVarintMygogo(data, i, uint64(len(m.Name)))
	i += copy(data[i:], m.Name)
	if len(m.Colors) > 0 {
		for _, s := range m.Colors {
			data[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Mygogo(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Mygogo(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMygogo(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *GogoProtoColorGroup) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovMygogo(uint64(m.Id))
	l = len(m.Name)
	n += 1 + l + sovMygogo(uint64(l))
	if len(m.Colors) > 0 {
		for _, s := range m.Colors {
			l = len(s)
			n += 1 + l + sovMygogo(uint64(l))
		}
	}
	return n
}

func sovMygogo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMygogo(x uint64) (n int) {
	return sovMygogo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GogoProtoColorGroup) Unmarshal(data []byte) error {
	var hasFields [1]uint64
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMygogo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GogoProtoColorGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GogoProtoColorGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMygogo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(data[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Colors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMygogo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Colors = append(m.Colors, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMygogo(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMygogo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("id")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMygogo(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMygogo
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowMygogo
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMygogo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMygogo
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipMygogo(data[start:])
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
	ErrInvalidLengthMygogo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMygogo   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorMygogo = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xad, 0x4c, 0xcf,
	0x4f, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0xcf, 0x2f, 0x4e, 0x2d, 0x4a,
	0xce, 0xcf, 0x2d, 0x50, 0xf2, 0xe6, 0x12, 0x76, 0x07, 0x4a, 0x04, 0x80, 0xc4, 0x9d, 0xf3, 0x73,
	0xf2, 0x8b, 0xdc, 0x8b, 0xf2, 0x4b, 0x0b, 0x84, 0x04, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15,
	0x98, 0x34, 0x58, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x10, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc,
	0x4d, 0x95, 0x60, 0x02, 0x8a, 0x71, 0x42, 0xc5, 0xf8, 0xb8, 0xd8, 0x92, 0x41, 0x7a, 0x8a, 0x25,
	0x98, 0x15, 0x98, 0x81, 0xa2, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x00, 0xe2, 0x07, 0x40, 0x3c,
	0xe1, 0xb1, 0x1c, 0x03, 0x20, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x57, 0x7a, 0x4c, 0x78, 0x00, 0x00,
	0x00,
}
