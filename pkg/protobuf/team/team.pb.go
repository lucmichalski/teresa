// Code generated by protoc-gen-go.
// source: pkg/protobuf/team/team.proto
// DO NOT EDIT!

/*
Package team is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/team/team.proto

It has these top-level messages:
	CreateRequest
	AddUserRequest
	RemoveUserRequest
	ListResponse
	RenameRequest
	Empty
*/
package team

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

type CreateRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Url   string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type AddUserRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *AddUserRequest) Reset()                    { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()               {}
func (*AddUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddUserRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type RemoveUserRequest struct {
	Team string `protobuf:"bytes,1,opt,name=team" json:"team,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *RemoveUserRequest) Reset()                    { *m = RemoveUserRequest{} }
func (m *RemoveUserRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveUserRequest) ProtoMessage()               {}
func (*RemoveUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RemoveUserRequest) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *RemoveUserRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type ListResponse struct {
	Teams []*ListResponse_Team `protobuf:"bytes,1,rep,name=teams" json:"teams,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListResponse) GetTeams() []*ListResponse_Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type ListResponse_User struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *ListResponse_User) Reset()                    { *m = ListResponse_User{} }
func (m *ListResponse_User) String() string            { return proto.CompactTextString(m) }
func (*ListResponse_User) ProtoMessage()               {}
func (*ListResponse_User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *ListResponse_User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListResponse_User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ListResponse_Team struct {
	Name  string               `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email string               `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Url   string               `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Users []*ListResponse_User `protobuf:"bytes,4,rep,name=users" json:"users,omitempty"`
}

func (m *ListResponse_Team) Reset()                    { *m = ListResponse_Team{} }
func (m *ListResponse_Team) String() string            { return proto.CompactTextString(m) }
func (*ListResponse_Team) ProtoMessage()               {}
func (*ListResponse_Team) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 1} }

func (m *ListResponse_Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListResponse_Team) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ListResponse_Team) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ListResponse_Team) GetUsers() []*ListResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type RenameRequest struct {
	OldName string `protobuf:"bytes,1,opt,name=oldName" json:"oldName,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=newName" json:"newName,omitempty"`
}

func (m *RenameRequest) Reset()                    { *m = RenameRequest{} }
func (m *RenameRequest) String() string            { return proto.CompactTextString(m) }
func (*RenameRequest) ProtoMessage()               {}
func (*RenameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RenameRequest) GetOldName() string {
	if m != nil {
		return m.OldName
	}
	return ""
}

func (m *RenameRequest) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "team.CreateRequest")
	proto.RegisterType((*AddUserRequest)(nil), "team.AddUserRequest")
	proto.RegisterType((*RemoveUserRequest)(nil), "team.RemoveUserRequest")
	proto.RegisterType((*ListResponse)(nil), "team.ListResponse")
	proto.RegisterType((*ListResponse_User)(nil), "team.ListResponse.User")
	proto.RegisterType((*ListResponse_Team)(nil), "team.ListResponse.Team")
	proto.RegisterType((*RenameRequest)(nil), "team.RenameRequest")
	proto.RegisterType((*Empty)(nil), "team.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Team service

type TeamClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*Empty, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*Empty, error)
	Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*Empty, error)
}

type teamClient struct {
	cc *grpc.ClientConn
}

func NewTeamClient(cc *grpc.ClientConn) TeamClient {
	return &teamClient{cc}
}

func (c *teamClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/team.Team/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/team.Team/AddUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/team.Team/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/team.Team/RemoveUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamClient) Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/team.Team/Rename", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Team service

type TeamServer interface {
	Create(context.Context, *CreateRequest) (*Empty, error)
	AddUser(context.Context, *AddUserRequest) (*Empty, error)
	List(context.Context, *Empty) (*ListResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*Empty, error)
	Rename(context.Context, *RenameRequest) (*Empty, error)
}

func RegisterTeamServer(s *grpc.Server, srv TeamServer) {
	s.RegisterService(&_Team_serviceDesc, srv)
}

func _Team_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Team_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.Team/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServer).Rename(ctx, req.(*RenameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Team_serviceDesc = grpc.ServiceDesc{
	ServiceName: "team.Team",
	HandlerType: (*TeamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Team_Create_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _Team_AddUser_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Team_List_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _Team_RemoveUser_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _Team_Rename_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/team/team.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/team/team.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x25, 0x5f, 0xb7, 0x2d, 0xdf, 0xd4, 0x8a, 0x8e, 0x05, 0x97, 0xe0, 0xa1, 0xe4, 0x62, 0x29,
	0xda, 0x4a, 0xbd, 0x08, 0x9e, 0xa4, 0x78, 0x52, 0x3c, 0x2c, 0xf5, 0x07, 0xa4, 0x74, 0x94, 0x62,
	0x36, 0x49, 0xb3, 0x1b, 0xc5, 0xbf, 0xeb, 0xcd, 0x7f, 0x21, 0xbb, 0x9b, 0xd0, 0xae, 0x55, 0x11,
	0x2f, 0x61, 0xf6, 0xe5, 0xcd, 0xcc, 0x9b, 0x37, 0x03, 0x47, 0xf9, 0xd3, 0xe3, 0x38, 0x2f, 0x32,
	0x9d, 0xcd, 0xcb, 0x87, 0xb1, 0xa6, 0x58, 0xda, 0xcf, 0xc8, 0x42, 0xc8, 0x4c, 0x1c, 0xdd, 0x40,
	0x77, 0x5a, 0x50, 0xac, 0x49, 0xd0, 0xaa, 0x24, 0xa5, 0x11, 0x81, 0xa5, 0xb1, 0x24, 0x1e, 0xf4,
	0x83, 0xc1, 0x7f, 0x61, 0x63, 0xec, 0x41, 0x93, 0x64, 0xbc, 0x4c, 0xf8, 0x3f, 0x0b, 0xba, 0x07,
	0xee, 0x41, 0xa3, 0x2c, 0x12, 0xde, 0xb0, 0x98, 0x09, 0xa3, 0x0b, 0xd8, 0xbd, 0x5a, 0x2c, 0xee,
	0x15, 0x15, 0x3f, 0x55, 0x43, 0x60, 0xa5, 0xa2, 0xa2, 0x2a, 0x66, 0xe3, 0xe8, 0x12, 0xf6, 0x05,
	0xc9, 0xec, 0x99, 0x3e, 0x25, 0x1b, 0x8d, 0x75, 0xb2, 0x89, 0xbf, 0x4c, 0x7e, 0x0b, 0x60, 0xe7,
	0x76, 0xa9, 0xb4, 0x20, 0x95, 0x67, 0xa9, 0x22, 0x3c, 0x85, 0xa6, 0x21, 0x2b, 0x1e, 0xf4, 0x1b,
	0x83, 0xce, 0xe4, 0x70, 0x64, 0xc7, 0xde, 0xa4, 0x8c, 0x66, 0x14, 0x4b, 0xe1, 0x58, 0xe1, 0x19,
	0x30, 0xd3, 0xf6, 0xf7, 0xa3, 0x87, 0x2b, 0x60, 0xb3, 0x4a, 0xcd, 0x5f, 0xcd, 0x32, 0x22, 0x8d,
	0x7a, 0xc5, 0xd9, 0xb7, 0x22, 0xad, 0x19, 0x8e, 0x15, 0x4d, 0xa1, 0x2b, 0xc8, 0x34, 0xa8, 0xdd,
	0xe1, 0xd0, 0xce, 0x92, 0xc5, 0xdd, 0xba, 0x7d, 0xfd, 0x34, 0x7f, 0x52, 0x7a, 0xb1, 0x7f, 0x9c,
	0x86, 0xfa, 0x19, 0xb5, 0xa1, 0x79, 0x2d, 0x73, 0xfd, 0x3a, 0x79, 0x0f, 0xaa, 0x09, 0x86, 0xd0,
	0x72, 0xfb, 0xc7, 0x03, 0x27, 0xc0, 0xbb, 0x86, 0xb0, 0xe3, 0x40, 0x9b, 0x84, 0x27, 0xd0, 0xae,
	0xd6, 0x8b, 0x3d, 0x87, 0xfb, 0xdb, 0xf6, 0xd9, 0xc7, 0xc0, 0xcc, 0x30, 0xb8, 0x09, 0x86, 0xb8,
	0x3d, 0x25, 0x4e, 0x00, 0xd6, 0xbb, 0xc7, 0xca, 0x87, 0xad, 0x6b, 0xf0, 0x8b, 0x0f, 0xa1, 0xe5,
	0xdc, 0xa8, 0x65, 0x7b, 0xde, 0x78, 0xdc, 0x79, 0xcb, 0xde, 0xfb, 0xf9, 0x47, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1f, 0x1a, 0x40, 0x2f, 0x0f, 0x03, 0x00, 0x00,
}
