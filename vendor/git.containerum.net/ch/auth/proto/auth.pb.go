// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

/*
Package authProto is a generated protocol buffer package.

It is generated from these files:
	auth.proto
	auth_types.proto
	uuid.proto

It has these top-level messages:
	CreateTokenRequest
	CreateTokenResponse
	CheckTokenRequest
	CheckTokenResponse
	ExtendTokenRequest
	ExtendTokenResponse
	UpdateAccessRequestElement
	UpdateAccessRequest
	GetUserTokensRequest
	GetUserTokensResponse
	DeleteTokenRequest
	DeleteUserTokensRequest
	StoredToken
	AccessObject
	ResourcesAccess
	StoredTokenForUser
	UUID
*/
package authProto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"

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

type CreateTokenRequest struct {
	UserAgent   string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
	UserId      *UUID  `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// @inject_tag: binding:"ip"
	UserIp      string           `protobuf:"bytes,4,opt,name=user_ip,json=userIp" json:"user_ip,omitempty" binding:"ip"`
	UserRole    string           `protobuf:"bytes,5,opt,name=user_role,json=userRole" json:"user_role,omitempty"`
	RwAccess    bool             `protobuf:"varint,6,opt,name=rw_access,json=rwAccess" json:"rw_access,omitempty"`
	Access      *ResourcesAccess `protobuf:"bytes,7,opt,name=access" json:"access,omitempty"`
	PartTokenId *UUID            `protobuf:"bytes,8,opt,name=part_token_id,json=partTokenId" json:"part_token_id,omitempty"`
}

func (m *CreateTokenRequest) Reset()                    { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()               {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateTokenRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *CreateTokenRequest) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *CreateTokenRequest) GetUserId() *UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *CreateTokenRequest) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

func (m *CreateTokenRequest) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

func (m *CreateTokenRequest) GetRwAccess() bool {
	if m != nil {
		return m.RwAccess
	}
	return false
}

func (m *CreateTokenRequest) GetAccess() *ResourcesAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *CreateTokenRequest) GetPartTokenId() *UUID {
	if m != nil {
		return m.PartTokenId
	}
	return nil
}

type CreateTokenResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *CreateTokenResponse) Reset()                    { *m = CreateTokenResponse{} }
func (m *CreateTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTokenResponse) ProtoMessage()               {}
func (*CreateTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CreateTokenResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type CheckTokenRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	UserAgent   string `protobuf:"bytes,2,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	FingerPrint string `protobuf:"bytes,3,opt,name=finger_print,json=fingerPrint" json:"finger_print,omitempty"`
	// @inject_tag: binding:"ip"
	UserIp string `protobuf:"bytes,4,opt,name=user_ip,json=userIp" json:"user_ip,omitempty" binding:"ip"`
}

func (m *CheckTokenRequest) Reset()                    { *m = CheckTokenRequest{} }
func (m *CheckTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckTokenRequest) ProtoMessage()               {}
func (*CheckTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CheckTokenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CheckTokenRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *CheckTokenRequest) GetFingerPrint() string {
	if m != nil {
		return m.FingerPrint
	}
	return ""
}

func (m *CheckTokenRequest) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

type CheckTokenResponse struct {
	Access      *ResourcesAccess `protobuf:"bytes,1,opt,name=access" json:"access,omitempty"`
	UserId      *UUID            `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserRole    string           `protobuf:"bytes,3,opt,name=user_role,json=userRole" json:"user_role,omitempty"`
	TokenId     *UUID            `protobuf:"bytes,4,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	PartTokenId *UUID            `protobuf:"bytes,5,opt,name=part_token_id,json=partTokenId" json:"part_token_id,omitempty"`
}

func (m *CheckTokenResponse) Reset()                    { *m = CheckTokenResponse{} }
func (m *CheckTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckTokenResponse) ProtoMessage()               {}
func (*CheckTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CheckTokenResponse) GetAccess() *ResourcesAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *CheckTokenResponse) GetUserId() *UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *CheckTokenResponse) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

func (m *CheckTokenResponse) GetTokenId() *UUID {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *CheckTokenResponse) GetPartTokenId() *UUID {
	if m != nil {
		return m.PartTokenId
	}
	return nil
}

type ExtendTokenRequest struct {
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	Fingerprint  string `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *ExtendTokenRequest) Reset()                    { *m = ExtendTokenRequest{} }
func (m *ExtendTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*ExtendTokenRequest) ProtoMessage()               {}
func (*ExtendTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExtendTokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *ExtendTokenRequest) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

type ExtendTokenResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *ExtendTokenResponse) Reset()                    { *m = ExtendTokenResponse{} }
func (m *ExtendTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*ExtendTokenResponse) ProtoMessage()               {}
func (*ExtendTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ExtendTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *ExtendTokenResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type UpdateAccessRequestElement struct {
	UserId *UUID            `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Access *ResourcesAccess `protobuf:"bytes,2,opt,name=access" json:"access,omitempty"`
}

func (m *UpdateAccessRequestElement) Reset()                    { *m = UpdateAccessRequestElement{} }
func (m *UpdateAccessRequestElement) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccessRequestElement) ProtoMessage()               {}
func (*UpdateAccessRequestElement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateAccessRequestElement) GetUserId() *UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *UpdateAccessRequestElement) GetAccess() *ResourcesAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

type UpdateAccessRequest struct {
	Users []*UpdateAccessRequestElement `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *UpdateAccessRequest) Reset()                    { *m = UpdateAccessRequest{} }
func (m *UpdateAccessRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccessRequest) ProtoMessage()               {}
func (*UpdateAccessRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateAccessRequest) GetUsers() []*UpdateAccessRequestElement {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUserTokensRequest struct {
	UserId *UUID `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetUserTokensRequest) Reset()                    { *m = GetUserTokensRequest{} }
func (m *GetUserTokensRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserTokensRequest) ProtoMessage()               {}
func (*GetUserTokensRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetUserTokensRequest) GetUserId() *UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

type GetUserTokensResponse struct {
	Tokens []*StoredTokenForUser `protobuf:"bytes,1,rep,name=tokens" json:"tokens,omitempty"`
}

func (m *GetUserTokensResponse) Reset()                    { *m = GetUserTokensResponse{} }
func (m *GetUserTokensResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserTokensResponse) ProtoMessage()               {}
func (*GetUserTokensResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetUserTokensResponse) GetTokens() []*StoredTokenForUser {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type DeleteTokenRequest struct {
	TokenId *UUID `protobuf:"bytes,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	UserId  *UUID `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *DeleteTokenRequest) Reset()                    { *m = DeleteTokenRequest{} }
func (m *DeleteTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTokenRequest) ProtoMessage()               {}
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteTokenRequest) GetTokenId() *UUID {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *DeleteTokenRequest) GetUserId() *UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

type DeleteUserTokensRequest struct {
	UserId *UUID `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *DeleteUserTokensRequest) Reset()                    { *m = DeleteUserTokensRequest{} }
func (m *DeleteUserTokensRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserTokensRequest) ProtoMessage()               {}
func (*DeleteUserTokensRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteUserTokensRequest) GetUserId() *UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTokenRequest)(nil), "CreateTokenRequest")
	proto.RegisterType((*CreateTokenResponse)(nil), "CreateTokenResponse")
	proto.RegisterType((*CheckTokenRequest)(nil), "CheckTokenRequest")
	proto.RegisterType((*CheckTokenResponse)(nil), "CheckTokenResponse")
	proto.RegisterType((*ExtendTokenRequest)(nil), "ExtendTokenRequest")
	proto.RegisterType((*ExtendTokenResponse)(nil), "ExtendTokenResponse")
	proto.RegisterType((*UpdateAccessRequestElement)(nil), "UpdateAccessRequestElement")
	proto.RegisterType((*UpdateAccessRequest)(nil), "UpdateAccessRequest")
	proto.RegisterType((*GetUserTokensRequest)(nil), "GetUserTokensRequest")
	proto.RegisterType((*GetUserTokensResponse)(nil), "GetUserTokensResponse")
	proto.RegisterType((*DeleteTokenRequest)(nil), "DeleteTokenRequest")
	proto.RegisterType((*DeleteUserTokensRequest)(nil), "DeleteUserTokensRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error)
	ExtendToken(ctx context.Context, in *ExtendTokenRequest, opts ...grpc.CallOption) (*ExtendTokenResponse, error)
	UpdateAccess(ctx context.Context, in *UpdateAccessRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	GetUserTokens(ctx context.Context, in *GetUserTokensRequest, opts ...grpc.CallOption) (*GetUserTokensResponse, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	DeleteUserTokens(ctx context.Context, in *DeleteUserTokensRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/CreateToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {
	out := new(CheckTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/CheckToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ExtendToken(ctx context.Context, in *ExtendTokenRequest, opts ...grpc.CallOption) (*ExtendTokenResponse, error) {
	out := new(ExtendTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/ExtendToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateAccess(ctx context.Context, in *UpdateAccessRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/Auth/UpdateAccess", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserTokens(ctx context.Context, in *GetUserTokensRequest, opts ...grpc.CallOption) (*GetUserTokensResponse, error) {
	out := new(GetUserTokensResponse)
	err := grpc.Invoke(ctx, "/Auth/GetUserTokens", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/Auth/DeleteToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteUserTokens(ctx context.Context, in *DeleteUserTokensRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/Auth/DeleteUserTokens", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	CheckToken(context.Context, *CheckTokenRequest) (*CheckTokenResponse, error)
	ExtendToken(context.Context, *ExtendTokenRequest) (*ExtendTokenResponse, error)
	UpdateAccess(context.Context, *UpdateAccessRequest) (*google_protobuf2.Empty, error)
	GetUserTokens(context.Context, *GetUserTokensRequest) (*GetUserTokensResponse, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*google_protobuf2.Empty, error)
	DeleteUserTokens(context.Context, *DeleteUserTokensRequest) (*google_protobuf2.Empty, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckToken(ctx, req.(*CheckTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ExtendToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ExtendToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/ExtendToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ExtendToken(ctx, req.(*ExtendTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/UpdateAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateAccess(ctx, req.(*UpdateAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/GetUserTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserTokens(ctx, req.(*GetUserTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteUserTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteUserTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/DeleteUserTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteUserTokens(ctx, req.(*DeleteUserTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Auth_CreateToken_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _Auth_CheckToken_Handler,
		},
		{
			MethodName: "ExtendToken",
			Handler:    _Auth_ExtendToken_Handler,
		},
		{
			MethodName: "UpdateAccess",
			Handler:    _Auth_UpdateAccess_Handler,
		},
		{
			MethodName: "GetUserTokens",
			Handler:    _Auth_GetUserTokens_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Auth_DeleteToken_Handler,
		},
		{
			MethodName: "DeleteUserTokens",
			Handler:    _Auth_DeleteUserTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0xa4, 0x49, 0xd3, 0x71, 0x2a, 0xf5, 0x1b, 0xa7, 0xad, 0xe5, 0xea, 0x43, 0xe9,
	0x72, 0x13, 0x84, 0xb4, 0x15, 0x45, 0x02, 0x81, 0x10, 0xa2, 0xf4, 0x00, 0xbd, 0xab, 0x0c, 0xe1,
	0x02, 0x84, 0xa2, 0x90, 0x4c, 0x0e, 0x6a, 0x1a, 0x9b, 0xf5, 0x5a, 0xa5, 0x8f, 0xc1, 0x03, 0xf0,
	0x2c, 0xdc, 0xf2, 0x58, 0x68, 0x77, 0xdd, 0xd6, 0x8e, 0x9d, 0x06, 0x24, 0xee, 0xec, 0x99, 0xd9,
	0xd9, 0xd9, 0xff, 0x6f, 0x76, 0x16, 0xa0, 0x17, 0xcb, 0x31, 0x0f, 0x45, 0x20, 0x03, 0x0f, 0xe2,
	0x78, 0x32, 0x48, 0xbe, 0x37, 0x94, 0xbd, 0x2b, 0xaf, 0x42, 0x8a, 0x12, 0xcb, 0xce, 0x28, 0x08,
	0x46, 0x53, 0xda, 0xd3, 0x7f, 0x5f, 0xe2, 0xe1, 0x1e, 0x5d, 0x84, 0xf2, 0xca, 0x38, 0xd9, 0x8f,
	0x32, 0xe0, 0xa1, 0xa0, 0x9e, 0xa4, 0xf7, 0xc1, 0x39, 0xcd, 0x7c, 0xfa, 0x1a, 0x53, 0x24, 0xf1,
	0x7f, 0x80, 0x38, 0x22, 0xd1, 0xed, 0x8d, 0x68, 0x26, 0x5d, 0xab, 0x65, 0xb5, 0xd7, 0xfc, 0x35,
	0x65, 0x39, 0x50, 0x06, 0x6c, 0x81, 0x3d, 0x9c, 0xcc, 0x46, 0x24, 0x42, 0x31, 0x99, 0x49, 0xb7,
	0xac, 0xfd, 0x69, 0x13, 0xde, 0x83, 0x55, 0x9d, 0x60, 0x32, 0x70, 0x2b, 0x2d, 0xab, 0x6d, 0xef,
	0x57, 0x79, 0xa7, 0x73, 0x7a, 0xe4, 0xd7, 0x94, 0xf5, 0x74, 0x80, 0xdb, 0xd7, 0xfe, 0xd0, 0x5d,
	0xd1, 0xab, 0x8d, 0x23, 0xc4, 0x1d, 0xd0, 0xfb, 0x74, 0x45, 0x30, 0x25, 0xb7, 0xaa, 0x5d, 0x75,
	0x65, 0xf0, 0x83, 0x29, 0x29, 0xa7, 0xb8, 0xec, 0xf6, 0xfa, 0x7d, 0x8a, 0x22, 0xb7, 0xd6, 0xb2,
	0xda, 0x75, 0xbf, 0x2e, 0x2e, 0x0f, 0xf4, 0x3f, 0xb6, 0xa1, 0x96, 0x78, 0x56, 0xf5, 0x8e, 0x1b,
	0xdc, 0xa7, 0x28, 0x88, 0x45, 0x9f, 0x22, 0x13, 0xe1, 0x27, 0x7e, 0x7c, 0x00, 0xeb, 0x61, 0x4f,
	0xc8, 0xae, 0x54, 0x47, 0x56, 0x25, 0xd6, 0xd3, 0x25, 0xda, 0xca, 0xa7, 0xd5, 0x38, 0x1d, 0xb0,
	0xcf, 0xe0, 0x64, 0xe4, 0x89, 0xc2, 0x60, 0x16, 0x11, 0xee, 0x42, 0xc3, 0xe4, 0x32, 0x39, 0x12,
	0x85, 0x6c, 0x63, 0xd3, 0xa1, 0x78, 0x1f, 0xd6, 0x05, 0x0d, 0x05, 0x45, 0xe3, 0x24, 0xc6, 0xa8,
	0xd4, 0x48, 0x8c, 0x3a, 0x88, 0x7d, 0xb7, 0xe0, 0xbf, 0xc3, 0x31, 0xf5, 0xcf, 0x33, 0xea, 0xff,
	0x41, 0xf6, 0x2c, 0xa0, 0xf2, 0x3c, 0xa0, 0x5d, 0x68, 0x18, 0x1a, 0x5d, 0x43, 0xa8, 0x92, 0x26,
	0x74, 0xa6, 0x09, 0x2d, 0x22, 0xc0, 0x7e, 0x59, 0x80, 0xe9, 0x9a, 0x92, 0x23, 0xdf, 0xca, 0x6b,
	0x2d, 0x91, 0x37, 0xc5, 0xbe, 0x5c, 0xc4, 0x3e, 0x83, 0xb8, 0x32, 0x87, 0xb8, 0x05, 0xf5, 0x1b,
	0x2c, 0x2b, 0xe9, 0xd5, 0xab, 0xd2, 0x20, 0xc9, 0xd3, 0xab, 0x2e, 0xa4, 0xf7, 0x09, 0xf0, 0xf8,
	0x9b, 0xa4, 0xd9, 0x20, 0x23, 0x6f, 0x8e, 0x8c, 0x95, 0x27, 0xb3, 0xbc, 0xc5, 0x55, 0x6b, 0x64,
	0x92, 0xff, 0xe3, 0xd6, 0x18, 0x82, 0xd7, 0x09, 0x07, 0x3d, 0x49, 0x89, 0xba, 0xa6, 0xf8, 0xe3,
	0x29, 0x5d, 0x50, 0xf6, 0x7e, 0x59, 0x45, 0x1a, 0xdf, 0xd2, 0x2a, 0xdf, 0x4d, 0x8b, 0xbd, 0x05,
	0xa7, 0x60, 0x1f, 0x7c, 0x04, 0x55, 0x95, 0x4a, 0xd1, 0xae, 0xb4, 0xed, 0xfd, 0x1d, 0xbe, 0xb8,
	0x18, 0xdf, 0x44, 0xb2, 0x27, 0xd0, 0x7c, 0x43, 0xb2, 0x13, 0x91, 0xd0, 0x27, 0xb8, 0x49, 0xb5,
	0xa4, 0x56, 0x76, 0x04, 0x9b, 0x73, 0xeb, 0x12, 0x29, 0x1f, 0x42, 0x4d, 0xeb, 0x73, 0x5d, 0x84,
	0xc3, 0xdf, 0xc9, 0x40, 0x90, 0x11, 0xfc, 0x24, 0x10, 0x6a, 0x89, 0x9f, 0x84, 0xb0, 0x0f, 0x80,
	0x47, 0x34, 0xa5, 0xb9, 0x41, 0x96, 0x6e, 0x27, 0xab, 0xb0, 0x9d, 0x96, 0x74, 0x2b, 0x7b, 0x06,
	0xdb, 0x26, 0xef, 0x5f, 0x1f, 0x6c, 0xff, 0x67, 0x05, 0x56, 0x0e, 0x62, 0x39, 0xc6, 0xe7, 0x60,
	0xa7, 0xa6, 0x08, 0x3a, 0x3c, 0x3f, 0x72, 0xbd, 0x26, 0x2f, 0x18, 0x34, 0xac, 0x84, 0x4f, 0x01,
	0x6e, 0x6f, 0x23, 0x22, 0xcf, 0x8d, 0x0b, 0xcf, 0xe1, 0xf9, 0xeb, 0xca, 0x4a, 0x6a, 0xd3, 0x54,
	0x7f, 0xa2, 0xc3, 0xf3, 0x57, 0xc1, 0x6b, 0xf2, 0x82, 0x16, 0x66, 0x25, 0x7c, 0x09, 0x8d, 0x34,
	0x6f, 0x6c, 0x16, 0xe1, 0xf7, 0xb6, 0xb8, 0x79, 0x5a, 0xf8, 0xf5, 0xd3, 0xc2, 0x8f, 0xd5, 0xd3,
	0xc2, 0x4a, 0xf8, 0x0a, 0xd6, 0x33, 0x48, 0x71, 0x93, 0x17, 0xb5, 0x86, 0xb7, 0xc5, 0x0b, 0xc9,
	0xb3, 0x12, 0xbe, 0x00, 0x3b, 0x85, 0x13, 0x1d, 0x9e, 0x87, 0x7b, 0xc7, 0xfe, 0x27, 0xb0, 0x31,
	0x0f, 0x0d, 0x5d, 0xbe, 0x80, 0xe3, 0xe2, 0x3c, 0xaf, 0xed, 0x8f, 0x6b, 0xea, 0x3d, 0x3d, 0xd3,
	0xf6, 0x9a, 0x76, 0x3f, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x03, 0xda, 0xca, 0x7b, 0x07,
	0x00, 0x00,
}
