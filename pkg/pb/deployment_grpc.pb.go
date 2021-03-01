// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// DeployClient is the client API for Deploy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeployClient interface {
	Deployments(ctx context.Context, in *GetDeploymentOpts, opts ...grpc.CallOption) (Deploy_DeploymentsClient, error)
	ReportStatus(ctx context.Context, in *DeploymentStatus, opts ...grpc.CallOption) (*ReportStatusOpts, error)
}

type deployClient struct {
	cc grpc.ClientConnInterface
}

func NewDeployClient(cc grpc.ClientConnInterface) DeployClient {
	return &deployClient{cc}
}

func (c *deployClient) Deployments(ctx context.Context, in *GetDeploymentOpts, opts ...grpc.CallOption) (Deploy_DeploymentsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Deploy_ServiceDesc.Streams[0], "/pb.Deploy/Deployments", opts...)
	if err != nil {
		return nil, err
	}
	x := &deployDeploymentsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Deploy_DeploymentsClient interface {
	Recv() (*DeploymentRequest, error)
	grpc.ClientStream
}

type deployDeploymentsClient struct {
	grpc.ClientStream
}

func (x *deployDeploymentsClient) Recv() (*DeploymentRequest, error) {
	m := new(DeploymentRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *deployClient) ReportStatus(ctx context.Context, in *DeploymentStatus, opts ...grpc.CallOption) (*ReportStatusOpts, error) {
	out := new(ReportStatusOpts)
	err := c.cc.Invoke(ctx, "/pb.Deploy/ReportStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeployServer is the server API for Deploy service.
// All implementations must embed UnimplementedDeployServer
// for forward compatibility
type DeployServer interface {
	Deployments(*GetDeploymentOpts, Deploy_DeploymentsServer) error
	ReportStatus(context.Context, *DeploymentStatus) (*ReportStatusOpts, error)
	mustEmbedUnimplementedDeployServer()
}

// UnimplementedDeployServer must be embedded to have forward compatible implementations.
type UnimplementedDeployServer struct {
}

func (UnimplementedDeployServer) Deployments(*GetDeploymentOpts, Deploy_DeploymentsServer) error {
	return status.Errorf(codes.Unimplemented, "method Deployments not implemented")
}
func (UnimplementedDeployServer) ReportStatus(context.Context, *DeploymentStatus) (*ReportStatusOpts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportStatus not implemented")
}
func (UnimplementedDeployServer) mustEmbedUnimplementedDeployServer() {}

// UnsafeDeployServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeployServer will
// result in compilation errors.
type UnsafeDeployServer interface {
	mustEmbedUnimplementedDeployServer()
}

func RegisterDeployServer(s grpc.ServiceRegistrar, srv DeployServer) {
	s.RegisterService(&Deploy_ServiceDesc, srv)
}

func _Deploy_Deployments_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDeploymentOpts)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DeployServer).Deployments(m, &deployDeploymentsServer{stream})
}

type Deploy_DeploymentsServer interface {
	Send(*DeploymentRequest) error
	grpc.ServerStream
}

type deployDeploymentsServer struct {
	grpc.ServerStream
}

func (x *deployDeploymentsServer) Send(m *DeploymentRequest) error {
	return x.ServerStream.SendMsg(m)
}

func _Deploy_ReportStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployServer).ReportStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Deploy/ReportStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployServer).ReportStatus(ctx, req.(*DeploymentStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// Deploy_ServiceDesc is the grpc.ServiceDesc for Deploy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Deploy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Deploy",
	HandlerType: (*DeployServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportStatus",
			Handler:    _Deploy_ReportStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Deployments",
			Handler:       _Deploy_Deployments_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/pb/deployment.proto",
}
