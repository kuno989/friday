// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package clamav_api

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

// ClamAVScannerClient is the client API for ClamAVScanner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClamAVScannerClient interface {
	ScanFile(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error)
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type clamAVScannerClient struct {
	cc grpc.ClientConnInterface
}

func NewClamAVScannerClient(cc grpc.ClientConnInterface) ClamAVScannerClient {
	return &clamAVScannerClient{cc}
}

func (c *clamAVScannerClient) ScanFile(ctx context.Context, in *ScanFileRequest, opts ...grpc.CallOption) (*ScanResponse, error) {
	out := new(ScanResponse)
	err := c.cc.Invoke(ctx, "/clamav.ClamAVScanner/ScanFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clamAVScannerClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/clamav.ClamAVScanner/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClamAVScannerServer is the server API for ClamAVScanner service.
// All implementations must embed UnimplementedClamAVScannerServer
// for forward compatibility
type ClamAVScannerServer interface {
	ScanFile(context.Context, *ScanFileRequest) (*ScanResponse, error)
	GetVersion(context.Context, *VersionRequest) (*VersionResponse, error)
	mustEmbedUnimplementedClamAVScannerServer()
}

// UnimplementedClamAVScannerServer must be embedded to have forward compatible implementations.
type UnimplementedClamAVScannerServer struct {
}

func (UnimplementedClamAVScannerServer) ScanFile(context.Context, *ScanFileRequest) (*ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScanFile not implemented")
}
func (UnimplementedClamAVScannerServer) GetVersion(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedClamAVScannerServer) mustEmbedUnimplementedClamAVScannerServer() {}

// UnsafeClamAVScannerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClamAVScannerServer will
// result in compilation errors.
type UnsafeClamAVScannerServer interface {
	mustEmbedUnimplementedClamAVScannerServer()
}

func RegisterClamAVScannerServer(s grpc.ServiceRegistrar, srv ClamAVScannerServer) {
	s.RegisterService(&ClamAVScanner_ServiceDesc, srv)
}

func _ClamAVScanner_ScanFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClamAVScannerServer).ScanFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clamav.ClamAVScanner/ScanFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClamAVScannerServer).ScanFile(ctx, req.(*ScanFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClamAVScanner_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClamAVScannerServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clamav.ClamAVScanner/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClamAVScannerServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClamAVScanner_ServiceDesc is the grpc.ServiceDesc for ClamAVScanner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClamAVScanner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clamav.ClamAVScanner",
	HandlerType: (*ClamAVScannerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScanFile",
			Handler:    _ClamAVScanner_ScanFile_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ClamAVScanner_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clamav.proto",
}