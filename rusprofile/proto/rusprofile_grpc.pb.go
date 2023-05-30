// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/rusprofile.proto

package proto

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

// RusProfileServiceClient is the client API for RusProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RusProfileServiceClient interface {
	GetCompanyInfo(ctx context.Context, in *CompanyRequest, opts ...grpc.CallOption) (*CompanyResponse, error)
	GetSwaggerUI(ctx context.Context, in *SwaggerUIRequest, opts ...grpc.CallOption) (*SwaggerUIResponse, error)
}

type rusProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRusProfileServiceClient(cc grpc.ClientConnInterface) RusProfileServiceClient {
	return &rusProfileServiceClient{cc}
}

func (c *rusProfileServiceClient) GetCompanyInfo(ctx context.Context, in *CompanyRequest, opts ...grpc.CallOption) (*CompanyResponse, error) {
	out := new(CompanyResponse)
	err := c.cc.Invoke(ctx, "/rusprofile.RusProfileService/GetCompanyInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rusProfileServiceClient) GetSwaggerUI(ctx context.Context, in *SwaggerUIRequest, opts ...grpc.CallOption) (*SwaggerUIResponse, error) {
	out := new(SwaggerUIResponse)
	err := c.cc.Invoke(ctx, "/rusprofile.RusProfileService/GetSwaggerUI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RusProfileServiceServer is the server API for RusProfileService service.
// All implementations must embed UnimplementedRusProfileServiceServer
// for forward compatibility
type RusProfileServiceServer interface {
	GetCompanyInfo(context.Context, *CompanyRequest) (*CompanyResponse, error)
	GetSwaggerUI(context.Context, *SwaggerUIRequest) (*SwaggerUIResponse, error)
	mustEmbedUnimplementedRusProfileServiceServer()
}

// UnimplementedRusProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRusProfileServiceServer struct {
}

func (UnimplementedRusProfileServiceServer) GetCompanyInfo(context.Context, *CompanyRequest) (*CompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyInfo not implemented")
}
func (UnimplementedRusProfileServiceServer) GetSwaggerUI(context.Context, *SwaggerUIRequest) (*SwaggerUIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSwaggerUI not implemented")
}
func (UnimplementedRusProfileServiceServer) mustEmbedUnimplementedRusProfileServiceServer() {}

// UnsafeRusProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RusProfileServiceServer will
// result in compilation errors.
type UnsafeRusProfileServiceServer interface {
	mustEmbedUnimplementedRusProfileServiceServer()
}

func RegisterRusProfileServiceServer(s grpc.ServiceRegistrar, srv RusProfileServiceServer) {
	s.RegisterService(&RusProfileService_ServiceDesc, srv)
}

func _RusProfileService_GetCompanyInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RusProfileServiceServer).GetCompanyInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusprofile.RusProfileService/GetCompanyInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RusProfileServiceServer).GetCompanyInfo(ctx, req.(*CompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RusProfileService_GetSwaggerUI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwaggerUIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RusProfileServiceServer).GetSwaggerUI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusprofile.RusProfileService/GetSwaggerUI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RusProfileServiceServer).GetSwaggerUI(ctx, req.(*SwaggerUIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RusProfileService_ServiceDesc is the grpc.ServiceDesc for RusProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RusProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rusprofile.RusProfileService",
	HandlerType: (*RusProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanyInfo",
			Handler:    _RusProfileService_GetCompanyInfo_Handler,
		},
		{
			MethodName: "GetSwaggerUI",
			Handler:    _RusProfileService_GetSwaggerUI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rusprofile.proto",
}
