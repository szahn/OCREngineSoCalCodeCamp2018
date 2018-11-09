// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuff/service.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Block struct {
	X1                   int32    `protobuf:"varint,1,opt,name=x1,proto3" json:"x1,omitempty"`
	Y1                   int32    `protobuf:"varint,2,opt,name=y1,proto3" json:"y1,omitempty"`
	X2                   int32    `protobuf:"varint,3,opt,name=x2,proto3" json:"x2,omitempty"`
	Y2                   int32    `protobuf:"varint,4,opt,name=y2,proto3" json:"y2,omitempty"`
	Page                 int32    `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Tag                  string   `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
	HasBox               bool     `protobuf:"varint,7,opt,name=hasBox,proto3" json:"hasBox,omitempty"`
	Text                 string   `protobuf:"bytes,8,opt,name=text,proto3" json:"text,omitempty"`
	Confidence           string   `protobuf:"bytes,9,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf4f5531b619200e, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetX1() int32 {
	if m != nil {
		return m.X1
	}
	return 0
}

func (m *Block) GetY1() int32 {
	if m != nil {
		return m.Y1
	}
	return 0
}

func (m *Block) GetX2() int32 {
	if m != nil {
		return m.X2
	}
	return 0
}

func (m *Block) GetY2() int32 {
	if m != nil {
		return m.Y2
	}
	return 0
}

func (m *Block) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Block) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Block) GetHasBox() bool {
	if m != nil {
		return m.HasBox
	}
	return false
}

func (m *Block) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Block) GetConfidence() string {
	if m != nil {
		return m.Confidence
	}
	return ""
}

type OCRRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OCRRequest) Reset()         { *m = OCRRequest{} }
func (m *OCRRequest) String() string { return proto.CompactTextString(m) }
func (*OCRRequest) ProtoMessage()    {}
func (*OCRRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf4f5531b619200e, []int{1}
}

func (m *OCRRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OCRRequest.Unmarshal(m, b)
}
func (m *OCRRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OCRRequest.Marshal(b, m, deterministic)
}
func (m *OCRRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OCRRequest.Merge(m, src)
}
func (m *OCRRequest) XXX_Size() int {
	return xxx_messageInfo_OCRRequest.Size(m)
}
func (m *OCRRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OCRRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OCRRequest proto.InternalMessageInfo

func (m *OCRRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type OCRResponse struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	ImageName            string   `protobuf:"bytes,2,opt,name=imageName,proto3" json:"imageName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OCRResponse) Reset()         { *m = OCRResponse{} }
func (m *OCRResponse) String() string { return proto.CompactTextString(m) }
func (*OCRResponse) ProtoMessage()    {}
func (*OCRResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf4f5531b619200e, []int{2}
}

func (m *OCRResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OCRResponse.Unmarshal(m, b)
}
func (m *OCRResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OCRResponse.Marshal(b, m, deterministic)
}
func (m *OCRResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OCRResponse.Merge(m, src)
}
func (m *OCRResponse) XXX_Size() int {
	return xxx_messageInfo_OCRResponse.Size(m)
}
func (m *OCRResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OCRResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OCRResponse proto.InternalMessageInfo

func (m *OCRResponse) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *OCRResponse) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func init() {
	proto.RegisterType((*Block)(nil), "service.Block")
	proto.RegisterType((*OCRRequest)(nil), "service.OCRRequest")
	proto.RegisterType((*OCRResponse)(nil), "service.OCRResponse")
}

func init() { proto.RegisterFile("protobuff/service.proto", fileDescriptor_bf4f5531b619200e) }

var fileDescriptor_bf4f5531b619200e = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xd1, 0x4e, 0xab, 0x40,
	0x10, 0x86, 0xcf, 0x42, 0x4b, 0xcb, 0x34, 0x69, 0x4e, 0x46, 0xa3, 0x9b, 0xc6, 0x18, 0xc2, 0x85,
	0xe1, 0xaa, 0x86, 0xf5, 0x05, 0x4c, 0x7b, 0x6f, 0x93, 0xed, 0x13, 0x00, 0x0e, 0x48, 0x6c, 0x59,
	0xec, 0x52, 0x43, 0x1f, 0xcd, 0xb7, 0x33, 0x0c, 0xd8, 0xea, 0xdd, 0x7c, 0xff, 0x3f, 0x33, 0x3b,
	0x33, 0x0b, 0xb7, 0xf5, 0xc1, 0x34, 0x26, 0x3d, 0xe6, 0xf9, 0xa3, 0xa5, 0xc3, 0x67, 0x99, 0xd1,
	0x92, 0x15, 0x9c, 0x0c, 0x18, 0x7e, 0x09, 0x18, 0xaf, 0x76, 0x26, 0x7b, 0xc7, 0x39, 0x38, 0x6d,
	0x2c, 0x45, 0x20, 0xa2, 0xb1, 0x76, 0xda, 0xb8, 0xe3, 0x53, 0x2c, 0x9d, 0x9e, 0x4f, 0xcc, 0xad,
	0x92, 0xee, 0xe0, 0x2b, 0xf6, 0x95, 0x1c, 0x0d, 0xbe, 0x42, 0x84, 0x51, 0x9d, 0x14, 0x24, 0xc7,
	0xac, 0x70, 0x8c, 0xff, 0xc1, 0x6d, 0x92, 0x42, 0x7a, 0x81, 0x88, 0x7c, 0xdd, 0x85, 0x78, 0x03,
	0xde, 0x5b, 0x62, 0x57, 0xa6, 0x95, 0x93, 0x40, 0x44, 0x53, 0x3d, 0x50, 0x57, 0xdd, 0x50, 0xdb,
	0xc8, 0x29, 0xa7, 0x72, 0x8c, 0xf7, 0x00, 0x99, 0xa9, 0xf2, 0xf2, 0x95, 0xaa, 0x8c, 0xa4, 0xcf,
	0xce, 0x2f, 0x25, 0x8c, 0x00, 0x36, 0x6b, 0xad, 0xe9, 0xe3, 0x48, 0xb6, 0xc1, 0x05, 0x4c, 0xf3,
	0x72, 0x47, 0x55, 0xb2, 0x27, 0xde, 0xc2, 0xd7, 0x67, 0x0e, 0xb7, 0x30, 0xe3, 0x4c, 0x5b, 0x9b,
	0xca, 0x12, 0x3e, 0x80, 0x97, 0x76, 0x3b, 0x5b, 0x29, 0x02, 0x37, 0x9a, 0xa9, 0xf9, 0xf2, 0xe7,
	0x3a, 0x7c, 0x0a, 0x3d, 0xb8, 0x78, 0x07, 0x7e, 0xb9, 0x4f, 0x0a, 0x7a, 0xe9, 0x7a, 0x3a, 0xdc,
	0xf3, 0x22, 0xa8, 0x67, 0x7e, 0x7e, 0xdb, 0x57, 0xa2, 0x02, 0x77, 0xb3, 0xd6, 0x78, 0x75, 0x6e,
	0x75, 0x19, 0x6d, 0x71, 0xfd, 0x57, 0xec, 0xa7, 0x08, 0xff, 0xa5, 0x1e, 0x7f, 0xc6, 0xd3, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x82, 0x19, 0x27, 0xa7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OCRServiceClient is the client API for OCRService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OCRServiceClient interface {
	OCR(ctx context.Context, in *OCRRequest, opts ...grpc.CallOption) (*OCRResponse, error)
}

type oCRServiceClient struct {
	cc *grpc.ClientConn
}

func NewOCRServiceClient(cc *grpc.ClientConn) OCRServiceClient {
	return &oCRServiceClient{cc}
}

func (c *oCRServiceClient) OCR(ctx context.Context, in *OCRRequest, opts ...grpc.CallOption) (*OCRResponse, error) {
	out := new(OCRResponse)
	err := c.cc.Invoke(ctx, "/service.OCRService/OCR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OCRServiceServer is the server API for OCRService service.
type OCRServiceServer interface {
	OCR(context.Context, *OCRRequest) (*OCRResponse, error)
}

func RegisterOCRServiceServer(s *grpc.Server, srv OCRServiceServer) {
	s.RegisterService(&_OCRService_serviceDesc, srv)
}

func _OCRService_OCR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OCRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCRServiceServer).OCR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.OCRService/OCR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCRServiceServer).OCR(ctx, req.(*OCRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OCRService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.OCRService",
	HandlerType: (*OCRServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OCR",
			Handler:    _OCRService_OCR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuff/service.proto",
}