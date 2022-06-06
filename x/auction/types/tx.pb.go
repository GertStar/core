// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gert/auction/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgPlaceBid represents a message used by bidders to place bids on auctions
type MsgPlaceBid struct {
	AuctionId uint64     `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	Bidder    string     `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
	Amount    types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgPlaceBid) Reset()         { *m = MsgPlaceBid{} }
func (m *MsgPlaceBid) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBid) ProtoMessage()    {}
func (*MsgPlaceBid) Descriptor() ([]byte, []int) {
	return fileDescriptor_226282be4da73be5, []int{0}
}
func (m *MsgPlaceBid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBid.Merge(m, src)
}
func (m *MsgPlaceBid) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBid) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBid.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBid proto.InternalMessageInfo

// MsgPlaceBidResponse defines the Msg/PlaceBid response type.
type MsgPlaceBidResponse struct {
}

func (m *MsgPlaceBidResponse) Reset()         { *m = MsgPlaceBidResponse{} }
func (m *MsgPlaceBidResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBidResponse) ProtoMessage()    {}
func (*MsgPlaceBidResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_226282be4da73be5, []int{1}
}
func (m *MsgPlaceBidResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBidResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBidResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBidResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBidResponse.Merge(m, src)
}
func (m *MsgPlaceBidResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBidResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBidResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBidResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPlaceBid)(nil), "gert.auction.v1beta1.MsgPlaceBid")
	proto.RegisterType((*MsgPlaceBidResponse)(nil), "gert.auction.v1beta1.MsgPlaceBidResponse")
}

func init() { proto.RegisterFile("gert/auction/v1beta1/tx.proto", fileDescriptor_226282be4da73be5) }

var fileDescriptor_226282be4da73be5 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x4e, 0x2c, 0x4b,
	0xd4, 0x4f, 0x2c, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34,
	0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x01, 0x49, 0xeb, 0x41, 0xa5,
	0xf5, 0xa0, 0xd2, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x05, 0xfa, 0x20, 0x16, 0x44, 0xad,
	0x94, 0x5c, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0x2a, 0xdc, 0xa4, 0xe4,
	0xfc, 0xcc, 0x3c, 0x88, 0xbc, 0x52, 0x3b, 0x23, 0x17, 0xb7, 0x6f, 0x71, 0x7a, 0x40, 0x4e, 0x62,
	0x72, 0xaa, 0x53, 0x66, 0x8a, 0x90, 0x2c, 0x17, 0x17, 0xd4, 0xe0, 0xf8, 0xcc, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x96, 0x20, 0x4e, 0xa8, 0x88, 0x67, 0x8a, 0x90, 0x18, 0x17, 0x5b, 0x52, 0x66,
	0x4a, 0x4a, 0x6a, 0x91, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x94, 0x27, 0x64, 0xce, 0xc5,
	0x96, 0x98, 0x9b, 0x5f, 0x9a, 0x57, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x07,
	0xb1, 0x57, 0x0f, 0x64, 0x2f, 0xcc, 0x89, 0x7a, 0xce, 0xf9, 0x99, 0x79, 0x4e, 0x2c, 0x27, 0xee,
	0xc9, 0x33, 0x04, 0x41, 0x95, 0x5b, 0x71, 0x74, 0x2c, 0x90, 0x67, 0x78, 0xb1, 0x40, 0x9e, 0x41,
	0x49, 0x94, 0x4b, 0x18, 0xc9, 0x21, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x46, 0xf1,
	0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x11, 0x5c, 0x1c, 0x70, 0x37, 0x2a, 0xea, 0x61, 0x0b, 0x00,
	0x3d, 0x24, 0xdd, 0x52, 0x9a, 0x04, 0x95, 0xc0, 0x2c, 0x70, 0x72, 0x3e, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e,
	0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xcd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc,
	0x5c, 0x7d, 0x90, 0x71, 0xba, 0x39, 0x89, 0x49, 0xc5, 0x60, 0x96, 0x7e, 0x05, 0x3c, 0x76, 0x4a,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xa1, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x97,
	0xae, 0x6b, 0x8a, 0xba, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// PlaceBid message type used by bidders to place bids on auctions
	PlaceBid(ctx context.Context, in *MsgPlaceBid, opts ...grpc.CallOption) (*MsgPlaceBidResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PlaceBid(ctx context.Context, in *MsgPlaceBid, opts ...grpc.CallOption) (*MsgPlaceBidResponse, error) {
	out := new(MsgPlaceBidResponse)
	err := c.cc.Invoke(ctx, "/gert.auction.v1beta1.Msg/PlaceBid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PlaceBid message type used by bidders to place bids on auctions
	PlaceBid(context.Context, *MsgPlaceBid) (*MsgPlaceBidResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PlaceBid(ctx context.Context, req *MsgPlaceBid) (*MsgPlaceBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBid not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PlaceBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPlaceBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PlaceBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gert.auction.v1beta1.Msg/PlaceBid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PlaceBid(ctx, req.(*MsgPlaceBid))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gert.auction.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBid",
			Handler:    _Msg_PlaceBid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gert/auction/v1beta1/tx.proto",
}

func (m *MsgPlaceBid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.AuctionId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlaceBidResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBidResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBidResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPlaceBid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovTx(uint64(m.AuctionId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgPlaceBidResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPlaceBid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPlaceBid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPlaceBidResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPlaceBidResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBidResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
