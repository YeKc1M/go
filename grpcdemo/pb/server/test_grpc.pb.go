// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package template_engine

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

// TestEngineClient is the client API for TestEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestEngineClient interface {
	HelloWorld(ctx context.Context, in *BasicRes, opts ...grpc.CallOption) (*StrResp, error)
	Echo(ctx context.Context, in *StrRes, opts ...grpc.CallOption) (*StrResp, error)
}

type testEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewTestEngineClient(cc grpc.ClientConnInterface) TestEngineClient {
	return &testEngineClient{cc}
}

func (c *testEngineClient) HelloWorld(ctx context.Context, in *BasicRes, opts ...grpc.CallOption) (*StrResp, error) {
	out := new(StrResp)
	err := c.cc.Invoke(ctx, "/template_engine.TestEngine/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testEngineClient) Echo(ctx context.Context, in *StrRes, opts ...grpc.CallOption) (*StrResp, error) {
	out := new(StrResp)
	err := c.cc.Invoke(ctx, "/template_engine.TestEngine/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestEngineServer is the server API for TestEngine service.
// All implementations must embed UnimplementedTestEngineServer
// for forward compatibility
type TestEngineServer interface {
	HelloWorld(context.Context, *BasicRes) (*StrResp, error)
	Echo(context.Context, *StrRes) (*StrResp, error)
	mustEmbedUnimplementedTestEngineServer()
}

// UnimplementedTestEngineServer must be embedded to have forward compatible implementations.
type UnimplementedTestEngineServer struct {
}

func (UnimplementedTestEngineServer) HelloWorld(context.Context, *BasicRes) (*StrResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedTestEngineServer) Echo(context.Context, *StrRes) (*StrResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedTestEngineServer) mustEmbedUnimplementedTestEngineServer() {}

// UnsafeTestEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestEngineServer will
// result in compilation errors.
type UnsafeTestEngineServer interface {
	mustEmbedUnimplementedTestEngineServer()
}

func RegisterTestEngineServer(s grpc.ServiceRegistrar, srv TestEngineServer) {
	s.RegisterService(&TestEngine_ServiceDesc, srv)
}

func _TestEngine_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestEngineServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template_engine.TestEngine/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestEngineServer).HelloWorld(ctx, req.(*BasicRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestEngine_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestEngineServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template_engine.TestEngine/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestEngineServer).Echo(ctx, req.(*StrRes))
	}
	return interceptor(ctx, in, info, handler)
}

// TestEngine_ServiceDesc is the grpc.ServiceDesc for TestEngine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestEngine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template_engine.TestEngine",
	HandlerType: (*TestEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _TestEngine_HelloWorld_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _TestEngine_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
