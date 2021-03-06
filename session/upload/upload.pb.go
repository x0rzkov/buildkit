// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upload.proto

package upload

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// BytesMessage contains a chunk of byte data
type BytesMessage struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BytesMessage) Reset()      { *m = BytesMessage{} }
func (*BytesMessage) ProtoMessage() {}
func (*BytesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b94b655bd2a7e5, []int{0}
}
func (m *BytesMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BytesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BytesMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BytesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesMessage.Merge(m, src)
}
func (m *BytesMessage) XXX_Size() int {
	return m.Size()
}
func (m *BytesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BytesMessage proto.InternalMessageInfo

func (m *BytesMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BytesMessage)(nil), "moby.upload.v1.BytesMessage")
}

func init() { proto.RegisterFile("upload.proto", fileDescriptor_91b94b655bd2a7e5) }

var fileDescriptor_91b94b655bd2a7e5 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2d, 0xc8, 0xc9,
	0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xcd, 0x4f, 0xaa, 0xd4, 0x83,
	0x0a, 0x95, 0x19, 0x2a, 0x29, 0x71, 0xf1, 0x38, 0x55, 0x96, 0xa4, 0x16, 0xfb, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x04, 0x81, 0xd9, 0x46, 0x01, 0x5c, 0x6c, 0xa1, 0x60, 0x0d, 0x42, 0x6e, 0x5c, 0x2c, 0x01,
	0xa5, 0x39, 0x39, 0x42, 0x32, 0x7a, 0xa8, 0xc6, 0xe8, 0x21, 0x9b, 0x21, 0x85, 0x57, 0x56, 0x83,
	0xd1, 0x80, 0xd1, 0xc9, 0xe6, 0xc2, 0x43, 0x39, 0x86, 0x1b, 0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94,
	0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xf8, 0xe2, 0x91, 0x1c, 0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x14, 0x1b, 0xc4, 0xc4,
	0x24, 0x36, 0xb0, 0x57, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12, 0xf2, 0xfc, 0xb4, 0xda,
	0x00, 0x00, 0x00,
}

func (this *BytesMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BytesMessage)
	if !ok {
		that2, ok := that.(BytesMessage)
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
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	return true
}
func (this *BytesMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&upload.BytesMessage{")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringUpload(v interface{}, typ string) string {
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

// UploadClient is the client API for Upload service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UploadClient interface {
	Pull(ctx context.Context, opts ...grpc.CallOption) (Upload_PullClient, error)
}

type uploadClient struct {
	cc *grpc.ClientConn
}

func NewUploadClient(cc *grpc.ClientConn) UploadClient {
	return &uploadClient{cc}
}

func (c *uploadClient) Pull(ctx context.Context, opts ...grpc.CallOption) (Upload_PullClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Upload_serviceDesc.Streams[0], "/moby.upload.v1.Upload/Pull", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadPullClient{stream}
	return x, nil
}

type Upload_PullClient interface {
	Send(*BytesMessage) error
	Recv() (*BytesMessage, error)
	grpc.ClientStream
}

type uploadPullClient struct {
	grpc.ClientStream
}

func (x *uploadPullClient) Send(m *BytesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadPullClient) Recv() (*BytesMessage, error) {
	m := new(BytesMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServer is the server API for Upload service.
type UploadServer interface {
	Pull(Upload_PullServer) error
}

func RegisterUploadServer(s *grpc.Server, srv UploadServer) {
	s.RegisterService(&_Upload_serviceDesc, srv)
}

func _Upload_Pull_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServer).Pull(&uploadPullServer{stream})
}

type Upload_PullServer interface {
	Send(*BytesMessage) error
	Recv() (*BytesMessage, error)
	grpc.ServerStream
}

type uploadPullServer struct {
	grpc.ServerStream
}

func (x *uploadPullServer) Send(m *BytesMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadPullServer) Recv() (*BytesMessage, error) {
	m := new(BytesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Upload_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moby.upload.v1.Upload",
	HandlerType: (*UploadServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pull",
			Handler:       _Upload_Pull_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "upload.proto",
}

func (m *BytesMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BytesMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUpload(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func encodeVarintUpload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BytesMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovUpload(uint64(l))
	}
	return n
}

func sovUpload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUpload(x uint64) (n int) {
	return sovUpload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *BytesMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BytesMessage{`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringUpload(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *BytesMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUpload
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
			return fmt.Errorf("proto: BytesMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BytesMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUpload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthUpload
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthUpload
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUpload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUpload
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthUpload
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
func skipUpload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUpload
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
					return 0, ErrIntOverflowUpload
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
					return 0, ErrIntOverflowUpload
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
				return 0, ErrInvalidLengthUpload
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthUpload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUpload
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
				next, err := skipUpload(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthUpload
				}
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
	ErrInvalidLengthUpload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUpload   = fmt.Errorf("proto: integer overflow")
)
