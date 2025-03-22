// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: search.proto

package protos

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	SearchUsers(ctx context.Context, in *SearchUsersRequest, opts ...grpc.CallOption) (*SearchUsersResponse, error)
	SearchUsersByDate(ctx context.Context, in *SearchUsersByDateRequest, opts ...grpc.CallOption) (*SearchUsersByDateResponse, error)
	SearchPosts(ctx context.Context, in *SearchPostsRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error)
	SearchPostsByDate(ctx context.Context, in *SearchPostsByDateRequest, opts ...grpc.CallOption) (*SearchPostsByDateResponse, error)
	SearchReports(ctx context.Context, in *SearchReportsRequest, opts ...grpc.CallOption) (*SearchReportsResponse, error)
	SearchReportsByDate(ctx context.Context, in *SearchReportsByDateRequest, opts ...grpc.CallOption) (*SearchReportsByDateResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) SearchUsers(ctx context.Context, in *SearchUsersRequest, opts ...grpc.CallOption) (*SearchUsersResponse, error) {
	out := new(SearchUsersResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/SearchUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchUsersByDate(ctx context.Context, in *SearchUsersByDateRequest, opts ...grpc.CallOption) (*SearchUsersByDateResponse, error) {
	out := new(SearchUsersByDateResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/SearchUsersByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchPosts(ctx context.Context, in *SearchPostsRequest, opts ...grpc.CallOption) (*SearchPostsResponse, error) {
	out := new(SearchPostsResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/SearchPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchPostsByDate(ctx context.Context, in *SearchPostsByDateRequest, opts ...grpc.CallOption) (*SearchPostsByDateResponse, error) {
	out := new(SearchPostsByDateResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/SearchPostsByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchReports(ctx context.Context, in *SearchReportsRequest, opts ...grpc.CallOption) (*SearchReportsResponse, error) {
	out := new(SearchReportsResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/SearchReports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchReportsByDate(ctx context.Context, in *SearchReportsByDateRequest, opts ...grpc.CallOption) (*SearchReportsByDateResponse, error) {
	out := new(SearchReportsByDateResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/SearchReportsByDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	SearchUsers(context.Context, *SearchUsersRequest) (*SearchUsersResponse, error)
	SearchUsersByDate(context.Context, *SearchUsersByDateRequest) (*SearchUsersByDateResponse, error)
	SearchPosts(context.Context, *SearchPostsRequest) (*SearchPostsResponse, error)
	SearchPostsByDate(context.Context, *SearchPostsByDateRequest) (*SearchPostsByDateResponse, error)
	SearchReports(context.Context, *SearchReportsRequest) (*SearchReportsResponse, error)
	SearchReportsByDate(context.Context, *SearchReportsByDateRequest) (*SearchReportsByDateResponse, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) SearchUsers(context.Context, *SearchUsersRequest) (*SearchUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedSearchServiceServer) SearchUsersByDate(context.Context, *SearchUsersByDateRequest) (*SearchUsersByDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsersByDate not implemented")
}
func (UnimplementedSearchServiceServer) SearchPosts(context.Context, *SearchPostsRequest) (*SearchPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPosts not implemented")
}
func (UnimplementedSearchServiceServer) SearchPostsByDate(context.Context, *SearchPostsByDateRequest) (*SearchPostsByDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPostsByDate not implemented")
}
func (UnimplementedSearchServiceServer) SearchReports(context.Context, *SearchReportsRequest) (*SearchReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchReports not implemented")
}
func (UnimplementedSearchServiceServer) SearchReportsByDate(context.Context, *SearchReportsByDateRequest) (*SearchReportsByDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchReportsByDate not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_SearchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/SearchUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchUsers(ctx, req.(*SearchUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchUsersByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUsersByDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchUsersByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/SearchUsersByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchUsersByDate(ctx, req.(*SearchUsersByDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/SearchPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchPosts(ctx, req.(*SearchPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchPostsByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostsByDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchPostsByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/SearchPostsByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchPostsByDate(ctx, req.(*SearchPostsByDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchReports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchReports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/SearchReports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchReports(ctx, req.(*SearchReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchReportsByDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReportsByDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchReportsByDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/SearchReportsByDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchReportsByDate(ctx, req.(*SearchReportsByDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUsers",
			Handler:    _SearchService_SearchUsers_Handler,
		},
		{
			MethodName: "SearchUsersByDate",
			Handler:    _SearchService_SearchUsersByDate_Handler,
		},
		{
			MethodName: "SearchPosts",
			Handler:    _SearchService_SearchPosts_Handler,
		},
		{
			MethodName: "SearchPostsByDate",
			Handler:    _SearchService_SearchPostsByDate_Handler,
		},
		{
			MethodName: "SearchReports",
			Handler:    _SearchService_SearchReports_Handler,
		},
		{
			MethodName: "SearchReportsByDate",
			Handler:    _SearchService_SearchReportsByDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}
