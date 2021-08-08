// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: memhouse_grpc_v3.proto

package memhouse_grpc_v3

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	memhouse "github.com/msaf1980/memhouse/protocol/memhouse"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type StoreMetricsResponce struct {
	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (m *StoreMetricsResponce) Reset()      { *m = StoreMetricsResponce{} }
func (*StoreMetricsResponce) ProtoMessage() {}
func (*StoreMetricsResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_9adf16f739661b56, []int{0}
}
func (m *StoreMetricsResponce) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoreMetricsResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoreMetricsResponce.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoreMetricsResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreMetricsResponce.Merge(m, src)
}
func (m *StoreMetricsResponce) XXX_Size() int {
	return m.Size()
}
func (m *StoreMetricsResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreMetricsResponce.DiscardUnknown(m)
}

var xxx_messageInfo_StoreMetricsResponce proto.InternalMessageInfo

func (m *StoreMetricsResponce) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StoreMetricsResponce) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*StoreMetricsResponce)(nil), "memhouse_grpc_v3.StoreMetricsResponce")
}

func init() { proto.RegisterFile("memhouse_grpc_v3.proto", fileDescriptor_9adf16f739661b56) }

var fileDescriptor_9adf16f739661b56 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4d, 0xcd, 0xcd,
	0xc8, 0x2f, 0x2d, 0x4e, 0x8d, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x33, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x40, 0x17, 0x97, 0xb2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0x4e, 0x4c, 0x33, 0xb4, 0xb4, 0x30, 0xd0, 0x87, 0xa9, 0xd2, 0x07,
	0xeb, 0x4a, 0xce, 0xcf, 0x41, 0x88, 0xc0, 0x18, 0x10, 0x03, 0x95, 0x9c, 0xb8, 0x44, 0x82, 0x4b,
	0xf2, 0x8b, 0x52, 0x7d, 0x53, 0x4b, 0x8a, 0x32, 0x93, 0x8b, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3,
	0x92, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x83, 0xc0, 0x6c, 0x21, 0x31, 0x2e, 0xb6, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x09, 0x26, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x28, 0xcf, 0x68, 0x05, 0x23, 0x17, 0x87, 0x2f, 0xd4, 0x58, 0x21, 0x3f,
	0x2e, 0x1e, 0xb7, 0xd4, 0x92, 0xe4, 0x0c, 0xa8, 0x81, 0x42, 0xb2, 0x7a, 0x70, 0x1b, 0x91, 0xc5,
	0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0xb0, 0x4b, 0xc3, 0x9c, 0xa1, 0xc4, 0x20, 0xe4,
	0xc3, 0xc5, 0x83, 0xec, 0x40, 0x21, 0x01, 0x84, 0x06, 0x88, 0x90, 0x94, 0x9a, 0x1e, 0x46, 0x60,
	0x61, 0xf3, 0x92, 0x12, 0x83, 0x06, 0xa3, 0x93, 0xc9, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca,
	0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x24, 0x36, 0x70, 0x58, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1c, 0x4f, 0x32, 0x9d, 0x96,
	0x01, 0x00, 0x00,
}

func (this *StoreMetricsResponce) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StoreMetricsResponce)
	if !ok {
		that2, ok := that.(StoreMetricsResponce)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (this *StoreMetricsResponce) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&memhouse_grpc_v3.StoreMetricsResponce{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMemhouseGrpcV3(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MemhouseClient is the client API for Memhouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemhouseClient interface {
	FetchMetrics(ctx context.Context, in *memhouse.FetchMetricsRequest, opts ...grpc.CallOption) (*memhouse.FetchMetricResponce, error)
	StoreMetrics(ctx context.Context, opts ...grpc.CallOption) (Memhouse_StoreMetricsClient, error)
}

type memhouseClient struct {
	cc *grpc.ClientConn
}

func NewMemhouseClient(cc *grpc.ClientConn) MemhouseClient {
	return &memhouseClient{cc}
}

func (c *memhouseClient) FetchMetrics(ctx context.Context, in *memhouse.FetchMetricsRequest, opts ...grpc.CallOption) (*memhouse.FetchMetricResponce, error) {
	out := new(memhouse.FetchMetricResponce)
	err := c.cc.Invoke(ctx, "/memhouse_grpc_v3.Memhouse/FetchMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memhouseClient) StoreMetrics(ctx context.Context, opts ...grpc.CallOption) (Memhouse_StoreMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Memhouse_serviceDesc.Streams[0], "/memhouse_grpc_v3.Memhouse/StoreMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &memhouseStoreMetricsClient{stream}
	return x, nil
}

type Memhouse_StoreMetricsClient interface {
	Send(*memhouse.Metric) error
	CloseAndRecv() (*StoreMetricsResponce, error)
	grpc.ClientStream
}

type memhouseStoreMetricsClient struct {
	grpc.ClientStream
}

func (x *memhouseStoreMetricsClient) Send(m *memhouse.Metric) error {
	return x.ClientStream.SendMsg(m)
}

func (x *memhouseStoreMetricsClient) CloseAndRecv() (*StoreMetricsResponce, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StoreMetricsResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MemhouseServer is the server API for Memhouse service.
type MemhouseServer interface {
	FetchMetrics(context.Context, *memhouse.FetchMetricsRequest) (*memhouse.FetchMetricResponce, error)
	StoreMetrics(Memhouse_StoreMetricsServer) error
}

// UnimplementedMemhouseServer can be embedded to have forward compatible implementations.
type UnimplementedMemhouseServer struct {
}

func (*UnimplementedMemhouseServer) FetchMetrics(ctx context.Context, req *memhouse.FetchMetricsRequest) (*memhouse.FetchMetricResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMetrics not implemented")
}
func (*UnimplementedMemhouseServer) StoreMetrics(srv Memhouse_StoreMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreMetrics not implemented")
}

func RegisterMemhouseServer(s *grpc.Server, srv MemhouseServer) {
	s.RegisterService(&_Memhouse_serviceDesc, srv)
}

func _Memhouse_FetchMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(memhouse.FetchMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemhouseServer).FetchMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/memhouse_grpc_v3.Memhouse/FetchMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemhouseServer).FetchMetrics(ctx, req.(*memhouse.FetchMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memhouse_StoreMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MemhouseServer).StoreMetrics(&memhouseStoreMetricsServer{stream})
}

type Memhouse_StoreMetricsServer interface {
	SendAndClose(*StoreMetricsResponce) error
	Recv() (*memhouse.Metric, error)
	grpc.ServerStream
}

type memhouseStoreMetricsServer struct {
	grpc.ServerStream
}

func (x *memhouseStoreMetricsServer) SendAndClose(m *StoreMetricsResponce) error {
	return x.ServerStream.SendMsg(m)
}

func (x *memhouseStoreMetricsServer) Recv() (*memhouse.Metric, error) {
	m := new(memhouse.Metric)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Memhouse_serviceDesc = grpc.ServiceDesc{
	ServiceName: "memhouse_grpc_v3.Memhouse",
	HandlerType: (*MemhouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchMetrics",
			Handler:    _Memhouse_FetchMetrics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StoreMetrics",
			Handler:       _Memhouse_StoreMetrics_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "memhouse_grpc_v3.proto",
}

func (m *StoreMetricsResponce) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreMetricsResponce) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoreMetricsResponce) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintMemhouseGrpcV3(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintMemhouseGrpcV3(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMemhouseGrpcV3(dAtA []byte, offset int, v uint64) int {
	offset -= sovMemhouseGrpcV3(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoreMetricsResponce) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovMemhouseGrpcV3(uint64(m.Code))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovMemhouseGrpcV3(uint64(l))
	}
	return n
}

func sovMemhouseGrpcV3(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMemhouseGrpcV3(x uint64) (n int) {
	return sovMemhouseGrpcV3(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StoreMetricsResponce) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StoreMetricsResponce{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMemhouseGrpcV3(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StoreMetricsResponce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMemhouseGrpcV3
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
			return fmt.Errorf("proto: StoreMetricsResponce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreMetricsResponce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemhouseGrpcV3
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMemhouseGrpcV3
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
				return ErrInvalidLengthMemhouseGrpcV3
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMemhouseGrpcV3
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMemhouseGrpcV3(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMemhouseGrpcV3
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
func skipMemhouseGrpcV3(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMemhouseGrpcV3
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
					return 0, ErrIntOverflowMemhouseGrpcV3
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
					return 0, ErrIntOverflowMemhouseGrpcV3
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
				return 0, ErrInvalidLengthMemhouseGrpcV3
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMemhouseGrpcV3
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMemhouseGrpcV3
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMemhouseGrpcV3        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMemhouseGrpcV3          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMemhouseGrpcV3 = fmt.Errorf("proto: unexpected end of group")
)