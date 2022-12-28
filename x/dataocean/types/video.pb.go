// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dataocean/dataocean/video.proto

package types

import (
	fmt "fmt"
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

type Video struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator     string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CoverLink   string `protobuf:"bytes,5,opt,name=coverLink,proto3" json:"coverLink,omitempty"`
	VideoLink   string `protobuf:"bytes,6,opt,name=videoLink,proto3" json:"videoLink,omitempty"`
	PriceMB     uint64 `protobuf:"varint,7,opt,name=priceMB,proto3" json:"priceMB,omitempty"`
	CreatedAt   uint64 `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (m *Video) Reset()         { *m = Video{} }
func (m *Video) String() string { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()    {}
func (*Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_c249b91701fc129e, []int{0}
}
func (m *Video) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Video.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Video.Merge(m, src)
}
func (m *Video) XXX_Size() int {
	return m.Size()
}
func (m *Video) XXX_DiscardUnknown() {
	xxx_messageInfo_Video.DiscardUnknown(m)
}

var xxx_messageInfo_Video proto.InternalMessageInfo

func (m *Video) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Video) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Video) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Video) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Video) GetCoverLink() string {
	if m != nil {
		return m.CoverLink
	}
	return ""
}

func (m *Video) GetVideoLink() string {
	if m != nil {
		return m.VideoLink
	}
	return ""
}

func (m *Video) GetPriceMB() uint64 {
	if m != nil {
		return m.PriceMB
	}
	return 0
}

func (m *Video) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*Video)(nil), "dataocean.dataocean.Video")
}

func init() { proto.RegisterFile("dataocean/dataocean/video.proto", fileDescriptor_c249b91701fc129e) }

var fileDescriptor_c249b91701fc129e = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0x49, 0x2c, 0x49,
	0xcc, 0x4f, 0x4e, 0x4d, 0xcc, 0xd3, 0x47, 0xb0, 0xca, 0x32, 0x53, 0x52, 0xf3, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0xc2, 0x7a, 0x70, 0x96, 0xd2, 0x5d, 0x46, 0x2e, 0xd6, 0x30,
	0x90, 0x22, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6,
	0xcc, 0x14, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48, 0x84, 0x8b, 0xb5, 0x24, 0xb3, 0x24, 0x27, 0x55, 0x82,
	0x19, 0x2c, 0x0e, 0xe1, 0x08, 0x29, 0x70, 0x71, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16, 0x94,
	0x64, 0xe6, 0xe7, 0x49, 0xb0, 0x80, 0xe5, 0x90, 0x85, 0x84, 0x64, 0xb8, 0x38, 0x93, 0xf3, 0xcb,
	0x52, 0x8b, 0x7c, 0x32, 0xf3, 0xb2, 0x25, 0x58, 0xc1, 0xf2, 0x08, 0x01, 0x90, 0x2c, 0xd8, 0xb5,
	0x60, 0x59, 0x36, 0x88, 0x2c, 0x5c, 0x00, 0xe4, 0x9a, 0x82, 0xa2, 0xcc, 0xe4, 0x54, 0x5f, 0x27,
	0x09, 0x76, 0xb0, 0x13, 0x61, 0x5c, 0xb0, 0xa9, 0x20, 0x87, 0xa5, 0xa6, 0x38, 0x96, 0x48, 0x70,
	0x80, 0xe5, 0x10, 0x02, 0x4e, 0xa6, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10,
	0x25, 0x8d, 0x08, 0xa5, 0x0a, 0xa4, 0x10, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x07,
	0x99, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xc1, 0xb0, 0xa0, 0x55, 0x01, 0x00, 0x00,
}

func (m *Video) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Video) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Video) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintVideo(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x40
	}
	if m.PriceMB != 0 {
		i = encodeVarintVideo(dAtA, i, uint64(m.PriceMB))
		i--
		dAtA[i] = 0x38
	}
	if len(m.VideoLink) > 0 {
		i -= len(m.VideoLink)
		copy(dAtA[i:], m.VideoLink)
		i = encodeVarintVideo(dAtA, i, uint64(len(m.VideoLink)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CoverLink) > 0 {
		i -= len(m.CoverLink)
		copy(dAtA[i:], m.CoverLink)
		i = encodeVarintVideo(dAtA, i, uint64(len(m.CoverLink)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintVideo(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintVideo(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVideo(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintVideo(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVideo(dAtA []byte, offset int, v uint64) int {
	offset -= sovVideo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Video) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovVideo(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVideo(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovVideo(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovVideo(uint64(l))
	}
	l = len(m.CoverLink)
	if l > 0 {
		n += 1 + l + sovVideo(uint64(l))
	}
	l = len(m.VideoLink)
	if l > 0 {
		n += 1 + l + sovVideo(uint64(l))
	}
	if m.PriceMB != 0 {
		n += 1 + sovVideo(uint64(m.PriceMB))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovVideo(uint64(m.CreatedAt))
	}
	return n
}

func sovVideo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVideo(x uint64) (n int) {
	return sovVideo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Video) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVideo
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
			return fmt.Errorf("proto: Video: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Video: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVideo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVideo
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
				return ErrInvalidLengthVideo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVideo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVideo
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
				return ErrInvalidLengthVideo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVideo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVideo
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
				return ErrInvalidLengthVideo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVideo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoverLink", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVideo
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
				return ErrInvalidLengthVideo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVideo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoverLink = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoLink", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVideo
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
				return ErrInvalidLengthVideo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVideo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VideoLink = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceMB", wireType)
			}
			m.PriceMB = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVideo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceMB |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVideo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVideo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVideo
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
func skipVideo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVideo
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
					return 0, ErrIntOverflowVideo
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
					return 0, ErrIntOverflowVideo
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
				return 0, ErrInvalidLengthVideo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVideo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVideo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVideo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVideo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVideo = fmt.Errorf("proto: unexpected end of group")
)
