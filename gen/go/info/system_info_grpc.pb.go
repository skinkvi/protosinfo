// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: info/system_info.proto

package infov1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SystemInfoService_SendCPU_FullMethodName              = "/systeminfo.SystemInfoService/SendCPU"
	SystemInfoService_SendCurrentProcesses_FullMethodName = "/systeminfo.SystemInfoService/SendCurrentProcesses"
)

// SystemInfoServiceClient is the client API for SystemInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemInfoServiceClient interface {
	SendCPU(ctx context.Context, in *CPURequest, opts ...grpc.CallOption) (*CPUResponse, error)
	SendCurrentProcesses(ctx context.Context, in *ProcessesRequest, opts ...grpc.CallOption) (*ProcessesResponse, error)
}

type systemInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemInfoServiceClient(cc grpc.ClientConnInterface) SystemInfoServiceClient {
	return &systemInfoServiceClient{cc}
}

func (c *systemInfoServiceClient) SendCPU(ctx context.Context, in *CPURequest, opts ...grpc.CallOption) (*CPUResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CPUResponse)
	err := c.cc.Invoke(ctx, SystemInfoService_SendCPU_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemInfoServiceClient) SendCurrentProcesses(ctx context.Context, in *ProcessesRequest, opts ...grpc.CallOption) (*ProcessesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProcessesResponse)
	err := c.cc.Invoke(ctx, SystemInfoService_SendCurrentProcesses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemInfoServiceServer is the server API for SystemInfoService service.
// All implementations must embed UnimplementedSystemInfoServiceServer
// for forward compatibility.
type SystemInfoServiceServer interface {
	SendCPU(context.Context, *CPURequest) (*CPUResponse, error)
	SendCurrentProcesses(context.Context, *ProcessesRequest) (*ProcessesResponse, error)
	mustEmbedUnimplementedSystemInfoServiceServer()
}

// UnimplementedSystemInfoServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSystemInfoServiceServer struct{}

func (UnimplementedSystemInfoServiceServer) SendCPU(context.Context, *CPURequest) (*CPUResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCPU not implemented")
}
func (UnimplementedSystemInfoServiceServer) SendCurrentProcesses(context.Context, *ProcessesRequest) (*ProcessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCurrentProcesses not implemented")
}
func (UnimplementedSystemInfoServiceServer) mustEmbedUnimplementedSystemInfoServiceServer() {}
func (UnimplementedSystemInfoServiceServer) testEmbeddedByValue()                           {}

// UnsafeSystemInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemInfoServiceServer will
// result in compilation errors.
type UnsafeSystemInfoServiceServer interface {
	mustEmbedUnimplementedSystemInfoServiceServer()
}

func RegisterSystemInfoServiceServer(s grpc.ServiceRegistrar, srv SystemInfoServiceServer) {
	// If the following call pancis, it indicates UnimplementedSystemInfoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SystemInfoService_ServiceDesc, srv)
}

func _SystemInfoService_SendCPU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CPURequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemInfoServiceServer).SendCPU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemInfoService_SendCPU_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemInfoServiceServer).SendCPU(ctx, req.(*CPURequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemInfoService_SendCurrentProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemInfoServiceServer).SendCurrentProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemInfoService_SendCurrentProcesses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemInfoServiceServer).SendCurrentProcesses(ctx, req.(*ProcessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemInfoService_ServiceDesc is the grpc.ServiceDesc for SystemInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "systeminfo.SystemInfoService",
	HandlerType: (*SystemInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCPU",
			Handler:    _SystemInfoService_SendCPU_Handler,
		},
		{
			MethodName: "SendCurrentProcesses",
			Handler:    _SystemInfoService_SendCurrentProcesses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "info/system_info.proto",
}
