// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ClustersClient is the client API for Clusters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClustersClient interface {
	GetClusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersResponse, error)
	GetNamespaces(ctx context.Context, in *GetNamespacesRequest, opts ...grpc.CallOption) (*GetNamespacesResponse, error)
	GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error)
}

type clustersClient struct {
	cc grpc.ClientConnInterface
}

func NewClustersClient(cc grpc.ClientConnInterface) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) GetClusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersResponse, error) {
	out := new(GetClustersResponse)
	err := c.cc.Invoke(ctx, "/clusters.Clusters/GetClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) GetNamespaces(ctx context.Context, in *GetNamespacesRequest, opts ...grpc.CallOption) (*GetNamespacesResponse, error) {
	out := new(GetNamespacesResponse)
	err := c.cc.Invoke(ctx, "/clusters.Clusters/GetNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) GetResources(ctx context.Context, in *GetResourcesRequest, opts ...grpc.CallOption) (*GetResourcesResponse, error) {
	out := new(GetResourcesResponse)
	err := c.cc.Invoke(ctx, "/clusters.Clusters/GetResources", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClustersServer is the server API for Clusters service.
// All implementations must embed UnimplementedClustersServer
// for forward compatibility
type ClustersServer interface {
	GetClusters(context.Context, *GetClustersRequest) (*GetClustersResponse, error)
	GetNamespaces(context.Context, *GetNamespacesRequest) (*GetNamespacesResponse, error)
	GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error)
	mustEmbedUnimplementedClustersServer()
}

// UnimplementedClustersServer must be embedded to have forward compatible implementations.
type UnimplementedClustersServer struct {
}

func (UnimplementedClustersServer) GetClusters(context.Context, *GetClustersRequest) (*GetClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusters not implemented")
}
func (UnimplementedClustersServer) GetNamespaces(context.Context, *GetNamespacesRequest) (*GetNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespaces not implemented")
}
func (UnimplementedClustersServer) GetResources(context.Context, *GetResourcesRequest) (*GetResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResources not implemented")
}
func (UnimplementedClustersServer) mustEmbedUnimplementedClustersServer() {}

// UnsafeClustersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClustersServer will
// result in compilation errors.
type UnsafeClustersServer interface {
	mustEmbedUnimplementedClustersServer()
}

func RegisterClustersServer(s grpc.ServiceRegistrar, srv ClustersServer) {
	s.RegisterService(&Clusters_ServiceDesc, srv)
}

func _Clusters_GetClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).GetClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusters.Clusters/GetClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).GetClusters(ctx, req.(*GetClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_GetNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).GetNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusters.Clusters/GetNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).GetNamespaces(ctx, req.(*GetNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_GetResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).GetResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clusters.Clusters/GetResources",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).GetResources(ctx, req.(*GetResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Clusters_ServiceDesc is the grpc.ServiceDesc for Clusters service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Clusters_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clusters.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusters",
			Handler:    _Clusters_GetClusters_Handler,
		},
		{
			MethodName: "GetNamespaces",
			Handler:    _Clusters_GetNamespaces_Handler,
		},
		{
			MethodName: "GetResources",
			Handler:    _Clusters_GetResources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clusters.proto",
}
