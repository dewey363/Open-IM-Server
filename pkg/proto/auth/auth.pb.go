// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/auth.proto

package pbAuth // import "./auth"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type UserRegisterReq struct {
	UID                  string   `protobuf:"bytes,1,opt,name=UID" json:"UID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=Icon" json:"Icon,omitempty"`
	Gender               int32    `protobuf:"varint,4,opt,name=Gender" json:"Gender,omitempty"`
	Mobile               string   `protobuf:"bytes,5,opt,name=Mobile" json:"Mobile,omitempty"`
	Birth                string   `protobuf:"bytes,6,opt,name=Birth" json:"Birth,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=Email" json:"Email,omitempty"`
	Ex                   string   `protobuf:"bytes,8,opt,name=Ex" json:"Ex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterReq) Reset()         { *m = UserRegisterReq{} }
func (m *UserRegisterReq) String() string { return proto.CompactTextString(m) }
func (*UserRegisterReq) ProtoMessage()    {}
func (*UserRegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d2199f7b1388fd2f, []int{0}
}
func (m *UserRegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterReq.Unmarshal(m, b)
}
func (m *UserRegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterReq.Marshal(b, m, deterministic)
}
func (dst *UserRegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterReq.Merge(dst, src)
}
func (m *UserRegisterReq) XXX_Size() int {
	return xxx_messageInfo_UserRegisterReq.Size(m)
}
func (m *UserRegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterReq proto.InternalMessageInfo

func (m *UserRegisterReq) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *UserRegisterReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserRegisterReq) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *UserRegisterReq) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserRegisterReq) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *UserRegisterReq) GetBirth() string {
	if m != nil {
		return m.Birth
	}
	return ""
}

func (m *UserRegisterReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserRegisterReq) GetEx() string {
	if m != nil {
		return m.Ex
	}
	return ""
}

type UserRegisterResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegisterResp) Reset()         { *m = UserRegisterResp{} }
func (m *UserRegisterResp) String() string { return proto.CompactTextString(m) }
func (*UserRegisterResp) ProtoMessage()    {}
func (*UserRegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d2199f7b1388fd2f, []int{1}
}
func (m *UserRegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegisterResp.Unmarshal(m, b)
}
func (m *UserRegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegisterResp.Marshal(b, m, deterministic)
}
func (dst *UserRegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegisterResp.Merge(dst, src)
}
func (m *UserRegisterResp) XXX_Size() int {
	return xxx_messageInfo_UserRegisterResp.Size(m)
}
func (m *UserRegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegisterResp proto.InternalMessageInfo

func (m *UserRegisterResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type UserTokenReq struct {
	Platform             int32    `protobuf:"varint,1,opt,name=Platform" json:"Platform,omitempty"`
	UID                  string   `protobuf:"bytes,2,opt,name=UID" json:"UID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTokenReq) Reset()         { *m = UserTokenReq{} }
func (m *UserTokenReq) String() string { return proto.CompactTextString(m) }
func (*UserTokenReq) ProtoMessage()    {}
func (*UserTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d2199f7b1388fd2f, []int{2}
}
func (m *UserTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTokenReq.Unmarshal(m, b)
}
func (m *UserTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTokenReq.Marshal(b, m, deterministic)
}
func (dst *UserTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTokenReq.Merge(dst, src)
}
func (m *UserTokenReq) XXX_Size() int {
	return xxx_messageInfo_UserTokenReq.Size(m)
}
func (m *UserTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserTokenReq proto.InternalMessageInfo

func (m *UserTokenReq) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *UserTokenReq) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

type UserTokenResp struct {
	ErrCode              int32    `protobuf:"varint,1,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=Token" json:"Token,omitempty"`
	ExpiredTime          int64    `protobuf:"varint,4,opt,name=ExpiredTime" json:"ExpiredTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserTokenResp) Reset()         { *m = UserTokenResp{} }
func (m *UserTokenResp) String() string { return proto.CompactTextString(m) }
func (*UserTokenResp) ProtoMessage()    {}
func (*UserTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d2199f7b1388fd2f, []int{3}
}
func (m *UserTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserTokenResp.Unmarshal(m, b)
}
func (m *UserTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserTokenResp.Marshal(b, m, deterministic)
}
func (dst *UserTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserTokenResp.Merge(dst, src)
}
func (m *UserTokenResp) XXX_Size() int {
	return xxx_messageInfo_UserTokenResp.Size(m)
}
func (m *UserTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserTokenResp proto.InternalMessageInfo

func (m *UserTokenResp) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *UserTokenResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *UserTokenResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserTokenResp) GetExpiredTime() int64 {
	if m != nil {
		return m.ExpiredTime
	}
	return 0
}

func init() {
	proto.RegisterType((*UserRegisterReq)(nil), "pbAuth.UserRegisterReq")
	proto.RegisterType((*UserRegisterResp)(nil), "pbAuth.UserRegisterResp")
	proto.RegisterType((*UserTokenReq)(nil), "pbAuth.UserTokenReq")
	proto.RegisterType((*UserTokenResp)(nil), "pbAuth.UserTokenResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	out := new(UserRegisterResp)
	err := grpc.Invoke(ctx, "/pbAuth.Auth/UserRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error) {
	out := new(UserTokenResp)
	err := grpc.Invoke(ctx, "/pbAuth.Auth/UserToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error)
	UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbAuth.Auth/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbAuth.Auth/UserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserToken(ctx, req.(*UserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbAuth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _Auth_UserRegister_Handler,
		},
		{
			MethodName: "UserToken",
			Handler:    _Auth_UserToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

func init() { proto.RegisterFile("auth/auth.proto", fileDescriptor_auth_d2199f7b1388fd2f) }

var fileDescriptor_auth_d2199f7b1388fd2f = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0x86, 0xd3, 0x42, 0x0b, 0xcc, 0xf7, 0x21, 0x64, 0x82, 0xba, 0xe1, 0x44, 0x7a, 0xe2, 0x60,
	0x6a, 0xa2, 0x17, 0x13, 0xbd, 0x80, 0x36, 0x86, 0x03, 0xc6, 0x54, 0xb8, 0x78, 0x2b, 0xb0, 0x42,
	0x23, 0x65, 0xeb, 0x6e, 0x49, 0xf0, 0xec, 0x8f, 0xf2, 0xef, 0x99, 0x9d, 0xdd, 0x12, 0x34, 0x5c,
	0xda, 0x79, 0x9f, 0x9d, 0xc9, 0xee, 0x3b, 0x33, 0xd0, 0x4a, 0xb6, 0xc5, 0xea, 0x52, 0x7f, 0xc2,
	0x5c, 0x8a, 0x42, 0xa0, 0x9f, 0xcf, 0x06, 0xdb, 0x62, 0x15, 0x7c, 0x3b, 0xd0, 0x9a, 0x2a, 0x2e,
	0x63, 0xbe, 0x4c, 0x55, 0xa1, 0xff, 0x1f, 0xd8, 0x86, 0xca, 0x74, 0xf4, 0xc0, 0x9c, 0x9e, 0xd3,
	0x6f, 0xc4, 0x3a, 0x44, 0x84, 0xea, 0x53, 0x92, 0x71, 0xe6, 0x12, 0xa2, 0x58, 0xb3, 0xd1, 0x5c,
	0x6c, 0x58, 0xc5, 0x30, 0x1d, 0xe3, 0x19, 0xf8, 0x8f, 0x7c, 0xb3, 0xe0, 0x92, 0x55, 0x7b, 0x4e,
	0xdf, 0x8b, 0xad, 0xd2, 0x7c, 0x2c, 0x66, 0xe9, 0x9a, 0x33, 0x8f, 0xb2, 0xad, 0xc2, 0x0e, 0x78,
	0xc3, 0x54, 0x16, 0x2b, 0xe6, 0x13, 0x36, 0x42, 0xd3, 0x28, 0x4b, 0xd2, 0x35, 0xab, 0x19, 0x4a,
	0x02, 0x4f, 0xc0, 0x8d, 0x76, 0xac, 0x4e, 0xc8, 0x8d, 0x76, 0xc1, 0x05, 0xb4, 0x7f, 0x3f, 0x5c,
	0xe5, 0xc8, 0xa0, 0xf6, 0xb2, 0x9d, 0xcf, 0xb9, 0x52, 0xf4, 0xfa, 0x7a, 0x5c, 0xca, 0xe0, 0x0e,
	0xfe, 0xeb, 0xec, 0x89, 0x78, 0xe7, 0x1b, 0xed, 0xb1, 0x0b, 0xf5, 0xe7, 0x75, 0x52, 0xbc, 0x09,
	0x99, 0x51, 0xaa, 0x17, 0xef, 0x75, 0xe9, 0xdf, 0xdd, 0xfb, 0x0f, 0x3e, 0xa1, 0x79, 0x50, 0x6d,
	0x2e, 0x8a, 0xa4, 0xbc, 0x17, 0x0b, 0x6e, 0xab, 0x4b, 0xa9, 0xad, 0x46, 0x52, 0x8e, 0xd5, 0xd2,
	0xd6, 0x5b, 0xa5, 0x4d, 0x51, 0xb9, 0xed, 0x97, 0x11, 0xd8, 0x83, 0x7f, 0xd1, 0x2e, 0x4f, 0x25,
	0x5f, 0x4c, 0xd2, 0x8c, 0x53, 0xd7, 0x2a, 0xf1, 0x21, 0xba, 0xfa, 0x72, 0xa0, 0xaa, 0x27, 0x85,
	0x03, 0xe3, 0xa0, 0xf4, 0x8b, 0xe7, 0xa1, 0x19, 0x61, 0xf8, 0x67, 0x7c, 0x5d, 0x76, 0xfc, 0x40,
	0xe5, 0x78, 0x03, 0x8d, 0xbd, 0x0d, 0xec, 0x1c, 0xa6, 0x95, 0x7d, 0xe9, 0x9e, 0x1e, 0xa1, 0x2a,
	0x1f, 0xb6, 0x5e, 0x9b, 0x21, 0xad, 0xcf, 0xad, 0x39, 0x9e, 0xf9, 0xb4, 0x46, 0xd7, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x15, 0x5d, 0xc8, 0xb6, 0x59, 0x02, 0x00, 0x00,
}
