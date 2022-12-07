// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: protos/author.proto

package blogpost

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorServiceClient interface {
	// Ping rpc...
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error)
	CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	DeleteAuthor(ctx context.Context, in *DeleteAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	GetAuthorList(ctx context.Context, in *GetAuthorListRequest, opts ...grpc.CallOption) (*GetAuthorListResponse, error)
	GetAuthorById(ctx context.Context, in *GetAuthorByIDRequest, opts ...grpc.CallOption) (*GetAuthorByIDResponse, error)
}

type authorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorServiceClient(cc grpc.ClientConnInterface) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/AuthorService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/AuthorService/CreateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/AuthorService/UpdateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) DeleteAuthor(ctx context.Context, in *DeleteAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/AuthorService/DeleteAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAuthorList(ctx context.Context, in *GetAuthorListRequest, opts ...grpc.CallOption) (*GetAuthorListResponse, error) {
	out := new(GetAuthorListResponse)
	err := c.cc.Invoke(ctx, "/AuthorService/GetAuthorList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAuthorById(ctx context.Context, in *GetAuthorByIDRequest, opts ...grpc.CallOption) (*GetAuthorByIDResponse, error) {
	out := new(GetAuthorByIDResponse)
	err := c.cc.Invoke(ctx, "/AuthorService/GetAuthorById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
// All implementations must embed UnimplementedAuthorServiceServer
// for forward compatibility
type AuthorServiceServer interface {
	// Ping rpc...
	Ping(context.Context, *Empty) (*Pong, error)
	CreateAuthor(context.Context, *CreateAuthorRequest) (*Author, error)
	UpdateAuthor(context.Context, *UpdateAuthorRequest) (*Author, error)
	DeleteAuthor(context.Context, *DeleteAuthorRequest) (*Author, error)
	GetAuthorList(context.Context, *GetAuthorListRequest) (*GetAuthorListResponse, error)
	GetAuthorById(context.Context, *GetAuthorByIDRequest) (*GetAuthorByIDResponse, error)
	mustEmbedUnimplementedAuthorServiceServer()
}

// UnimplementedAuthorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (UnimplementedAuthorServiceServer) Ping(context.Context, *Empty) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedAuthorServiceServer) CreateAuthor(context.Context, *CreateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) UpdateAuthor(context.Context, *UpdateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) DeleteAuthor(context.Context, *DeleteAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) GetAuthorList(context.Context, *GetAuthorListRequest) (*GetAuthorListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorList not implemented")
}
func (UnimplementedAuthorServiceServer) GetAuthorById(context.Context, *GetAuthorByIDRequest) (*GetAuthorByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorById not implemented")
}
func (UnimplementedAuthorServiceServer) mustEmbedUnimplementedAuthorServiceServer() {}

// UnsafeAuthorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServiceServer will
// result in compilation errors.
type UnsafeAuthorServiceServer interface {
	mustEmbedUnimplementedAuthorServiceServer()
}

func RegisterAuthorServiceServer(s grpc.ServiceRegistrar, srv AuthorServiceServer) {
	s.RegisterService(&AuthorService_ServiceDesc, srv)
}

func _AuthorService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthorService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthorService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, req.(*CreateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_UpdateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).UpdateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthorService/UpdateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).UpdateAuthor(ctx, req.(*UpdateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_DeleteAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).DeleteAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthorService/DeleteAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).DeleteAuthor(ctx, req.(*DeleteAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetAuthorList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthorList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthorService/GetAuthorList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthorList(ctx, req.(*GetAuthorListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetAuthorById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthorById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthorService/GetAuthorById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthorById(ctx, req.(*GetAuthorByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorService_ServiceDesc is the grpc.ServiceDesc for AuthorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _AuthorService_Ping_Handler,
		},
		{
			MethodName: "CreateAuthor",
			Handler:    _AuthorService_CreateAuthor_Handler,
		},
		{
			MethodName: "UpdateAuthor",
			Handler:    _AuthorService_UpdateAuthor_Handler,
		},
		{
			MethodName: "DeleteAuthor",
			Handler:    _AuthorService_DeleteAuthor_Handler,
		},
		{
			MethodName: "GetAuthorList",
			Handler:    _AuthorService_GetAuthorList_Handler,
		},
		{
			MethodName: "GetAuthorById",
			Handler:    _AuthorService_GetAuthorById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/author.proto",
}
