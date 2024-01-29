// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: storage.proto

package landmark_storage

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

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageServiceClient interface {
	GetLandmark(ctx context.Context, in *GetLandmarkRequest, opts ...grpc.CallOption) (*GetLandmarkResponse, error)
	AddLandmark(ctx context.Context, in *AddLandmarkRequest, opts ...grpc.CallOption) (*AddLandmarkResponse, error)
	LikeLandmark(ctx context.Context, in *LikeLandmarkRequest, opts ...grpc.CallOption) (*LikeLandmarkResponse, error)
	DislikeLandmark(ctx context.Context, in *DislikeLandmarkRequest, opts ...grpc.CallOption) (*DislikeLandmarkResponse, error)
	GetLikes(ctx context.Context, in *GetLikesRequest, opts ...grpc.CallOption) (*GetLikesResponse, error)
	ViewLandmark(ctx context.Context, in *ViewLandmarkRequest, opts ...grpc.CallOption) (*ViewLandmarkResponse, error)
	RecommendLandmarks(ctx context.Context, in *RecommendLandmarksRequest, opts ...grpc.CallOption) (*RecommendLandmarksResponse, error)
	GetRandomFeed(ctx context.Context, in *GetRandomFeedRequest, opts ...grpc.CallOption) (*GetRandomFeedResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	GetProfileComments(ctx context.Context, in *GetProfileCommentsRequest, opts ...grpc.CallOption) (*GetProfileCommentsResponse, error)
}

type storageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageServiceClient(cc grpc.ClientConnInterface) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) GetLandmark(ctx context.Context, in *GetLandmarkRequest, opts ...grpc.CallOption) (*GetLandmarkResponse, error) {
	out := new(GetLandmarkResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/GetLandmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) AddLandmark(ctx context.Context, in *AddLandmarkRequest, opts ...grpc.CallOption) (*AddLandmarkResponse, error) {
	out := new(AddLandmarkResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/AddLandmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) LikeLandmark(ctx context.Context, in *LikeLandmarkRequest, opts ...grpc.CallOption) (*LikeLandmarkResponse, error) {
	out := new(LikeLandmarkResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/LikeLandmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) DislikeLandmark(ctx context.Context, in *DislikeLandmarkRequest, opts ...grpc.CallOption) (*DislikeLandmarkResponse, error) {
	out := new(DislikeLandmarkResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/DislikeLandmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetLikes(ctx context.Context, in *GetLikesRequest, opts ...grpc.CallOption) (*GetLikesResponse, error) {
	out := new(GetLikesResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/GetLikes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) ViewLandmark(ctx context.Context, in *ViewLandmarkRequest, opts ...grpc.CallOption) (*ViewLandmarkResponse, error) {
	out := new(ViewLandmarkResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/ViewLandmark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) RecommendLandmarks(ctx context.Context, in *RecommendLandmarksRequest, opts ...grpc.CallOption) (*RecommendLandmarksResponse, error) {
	out := new(RecommendLandmarksResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/RecommendLandmarks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetRandomFeed(ctx context.Context, in *GetRandomFeedRequest, opts ...grpc.CallOption) (*GetRandomFeedResponse, error) {
	out := new(GetRandomFeedResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/GetRandomFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/GetComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageServiceClient) GetProfileComments(ctx context.Context, in *GetProfileCommentsRequest, opts ...grpc.CallOption) (*GetProfileCommentsResponse, error) {
	out := new(GetProfileCommentsResponse)
	err := c.cc.Invoke(ctx, "/landmark.storage.StorageService/GetProfileComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
// All implementations must embed UnimplementedStorageServiceServer
// for forward compatibility
type StorageServiceServer interface {
	GetLandmark(context.Context, *GetLandmarkRequest) (*GetLandmarkResponse, error)
	AddLandmark(context.Context, *AddLandmarkRequest) (*AddLandmarkResponse, error)
	LikeLandmark(context.Context, *LikeLandmarkRequest) (*LikeLandmarkResponse, error)
	DislikeLandmark(context.Context, *DislikeLandmarkRequest) (*DislikeLandmarkResponse, error)
	GetLikes(context.Context, *GetLikesRequest) (*GetLikesResponse, error)
	ViewLandmark(context.Context, *ViewLandmarkRequest) (*ViewLandmarkResponse, error)
	RecommendLandmarks(context.Context, *RecommendLandmarksRequest) (*RecommendLandmarksResponse, error)
	GetRandomFeed(context.Context, *GetRandomFeedRequest) (*GetRandomFeedResponse, error)
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error)
	GetProfileComments(context.Context, *GetProfileCommentsRequest) (*GetProfileCommentsResponse, error)
	mustEmbedUnimplementedStorageServiceServer()
}

// UnimplementedStorageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServiceServer struct {
}

func (UnimplementedStorageServiceServer) GetLandmark(context.Context, *GetLandmarkRequest) (*GetLandmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLandmark not implemented")
}
func (UnimplementedStorageServiceServer) AddLandmark(context.Context, *AddLandmarkRequest) (*AddLandmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLandmark not implemented")
}
func (UnimplementedStorageServiceServer) LikeLandmark(context.Context, *LikeLandmarkRequest) (*LikeLandmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeLandmark not implemented")
}
func (UnimplementedStorageServiceServer) DislikeLandmark(context.Context, *DislikeLandmarkRequest) (*DislikeLandmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DislikeLandmark not implemented")
}
func (UnimplementedStorageServiceServer) GetLikes(context.Context, *GetLikesRequest) (*GetLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikes not implemented")
}
func (UnimplementedStorageServiceServer) ViewLandmark(context.Context, *ViewLandmarkRequest) (*ViewLandmarkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewLandmark not implemented")
}
func (UnimplementedStorageServiceServer) RecommendLandmarks(context.Context, *RecommendLandmarksRequest) (*RecommendLandmarksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecommendLandmarks not implemented")
}
func (UnimplementedStorageServiceServer) GetRandomFeed(context.Context, *GetRandomFeedRequest) (*GetRandomFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomFeed not implemented")
}
func (UnimplementedStorageServiceServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedStorageServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedStorageServiceServer) GetComments(context.Context, *GetCommentsRequest) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedStorageServiceServer) GetProfileComments(context.Context, *GetProfileCommentsRequest) (*GetProfileCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileComments not implemented")
}
func (UnimplementedStorageServiceServer) mustEmbedUnimplementedStorageServiceServer() {}

// UnsafeStorageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServiceServer will
// result in compilation errors.
type UnsafeStorageServiceServer interface {
	mustEmbedUnimplementedStorageServiceServer()
}

func RegisterStorageServiceServer(s grpc.ServiceRegistrar, srv StorageServiceServer) {
	s.RegisterService(&StorageService_ServiceDesc, srv)
}

func _StorageService_GetLandmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLandmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetLandmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/GetLandmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetLandmark(ctx, req.(*GetLandmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_AddLandmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLandmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).AddLandmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/AddLandmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).AddLandmark(ctx, req.(*AddLandmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_LikeLandmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeLandmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).LikeLandmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/LikeLandmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).LikeLandmark(ctx, req.(*LikeLandmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_DislikeLandmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DislikeLandmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).DislikeLandmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/DislikeLandmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).DislikeLandmark(ctx, req.(*DislikeLandmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/GetLikes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetLikes(ctx, req.(*GetLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_ViewLandmark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewLandmarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).ViewLandmark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/ViewLandmark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).ViewLandmark(ctx, req.(*ViewLandmarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_RecommendLandmarks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecommendLandmarksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).RecommendLandmarks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/RecommendLandmarks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).RecommendLandmarks(ctx, req.(*RecommendLandmarksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetRandomFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetRandomFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/GetRandomFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetRandomFeed(ctx, req.(*GetRandomFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetComments(ctx, req.(*GetCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageService_GetProfileComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).GetProfileComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/landmark.storage.StorageService/GetProfileComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).GetProfileComments(ctx, req.(*GetProfileCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageService_ServiceDesc is the grpc.ServiceDesc for StorageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "landmark.storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLandmark",
			Handler:    _StorageService_GetLandmark_Handler,
		},
		{
			MethodName: "AddLandmark",
			Handler:    _StorageService_AddLandmark_Handler,
		},
		{
			MethodName: "LikeLandmark",
			Handler:    _StorageService_LikeLandmark_Handler,
		},
		{
			MethodName: "DislikeLandmark",
			Handler:    _StorageService_DislikeLandmark_Handler,
		},
		{
			MethodName: "GetLikes",
			Handler:    _StorageService_GetLikes_Handler,
		},
		{
			MethodName: "ViewLandmark",
			Handler:    _StorageService_ViewLandmark_Handler,
		},
		{
			MethodName: "RecommendLandmarks",
			Handler:    _StorageService_RecommendLandmarks_Handler,
		},
		{
			MethodName: "GetRandomFeed",
			Handler:    _StorageService_GetRandomFeed_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _StorageService_AddUser_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _StorageService_CreateComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _StorageService_GetComments_Handler,
		},
		{
			MethodName: "GetProfileComments",
			Handler:    _StorageService_GetProfileComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}