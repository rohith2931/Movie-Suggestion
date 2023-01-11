// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: msproto/ms.proto

package movieSuggestion

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

// MsDatabaseCrudClient is the client API for MsDatabaseCrud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsDatabaseCrudClient interface {
	CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error)
	GetAllMovies(ctx context.Context, in *EmptyMovie, opts ...grpc.CallOption) (*Movies, error)
	GetMovieByCategory(ctx context.Context, in *MovieCategory, opts ...grpc.CallOption) (*Movies, error)
	AddMovie(ctx context.Context, in *NewMovie, opts ...grpc.CallOption) (*Movie, error)
	DeleteMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Movie, error)
	AddMovieToWatchlist(ctx context.Context, in *AddMovieByUser, opts ...grpc.CallOption) (*Movie, error)
	GetAllWatchlistMovies(ctx context.Context, in *EmptyMovie, opts ...grpc.CallOption) (*Movies, error)
	DeleteMovieFromWatchlist(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Movie, error)
	CreateReview(ctx context.Context, in *NewReview, opts ...grpc.CallOption) (*Review, error)
	UpdateReview(ctx context.Context, in *Review, opts ...grpc.CallOption) (*Review, error)
}

type msDatabaseCrudClient struct {
	cc grpc.ClientConnInterface
}

func NewMsDatabaseCrudClient(cc grpc.ClientConnInterface) MsDatabaseCrudClient {
	return &msDatabaseCrudClient{cc}
}

func (c *msDatabaseCrudClient) CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) GetAllMovies(ctx context.Context, in *EmptyMovie, opts ...grpc.CallOption) (*Movies, error) {
	out := new(Movies)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/GetAllMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) GetMovieByCategory(ctx context.Context, in *MovieCategory, opts ...grpc.CallOption) (*Movies, error) {
	out := new(Movies)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/GetMovieByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) AddMovie(ctx context.Context, in *NewMovie, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/AddMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) DeleteMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/DeleteMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) AddMovieToWatchlist(ctx context.Context, in *AddMovieByUser, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/AddMovieToWatchlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) GetAllWatchlistMovies(ctx context.Context, in *EmptyMovie, opts ...grpc.CallOption) (*Movies, error) {
	out := new(Movies)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/GetAllWatchlistMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) DeleteMovieFromWatchlist(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/DeleteMovieFromWatchlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) CreateReview(ctx context.Context, in *NewReview, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/CreateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msDatabaseCrudClient) UpdateReview(ctx context.Context, in *Review, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/msproto.MsDatabaseCrud/UpdateReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsDatabaseCrudServer is the server API for MsDatabaseCrud service.
// All implementations must embed UnimplementedMsDatabaseCrudServer
// for forward compatibility
type MsDatabaseCrudServer interface {
	CreateUser(context.Context, *NewUser) (*User, error)
	GetAllMovies(context.Context, *EmptyMovie) (*Movies, error)
	GetMovieByCategory(context.Context, *MovieCategory) (*Movies, error)
	AddMovie(context.Context, *NewMovie) (*Movie, error)
	DeleteMovie(context.Context, *Movie) (*Movie, error)
	AddMovieToWatchlist(context.Context, *AddMovieByUser) (*Movie, error)
	GetAllWatchlistMovies(context.Context, *EmptyMovie) (*Movies, error)
	DeleteMovieFromWatchlist(context.Context, *Movie) (*Movie, error)
	CreateReview(context.Context, *NewReview) (*Review, error)
	UpdateReview(context.Context, *Review) (*Review, error)
	mustEmbedUnimplementedMsDatabaseCrudServer()
}

// UnimplementedMsDatabaseCrudServer must be embedded to have forward compatible implementations.
type UnimplementedMsDatabaseCrudServer struct {
}

func (UnimplementedMsDatabaseCrudServer) CreateUser(context.Context, *NewUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMsDatabaseCrudServer) GetAllMovies(context.Context, *EmptyMovie) (*Movies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMovies not implemented")
}
func (UnimplementedMsDatabaseCrudServer) GetMovieByCategory(context.Context, *MovieCategory) (*Movies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByCategory not implemented")
}
func (UnimplementedMsDatabaseCrudServer) AddMovie(context.Context, *NewMovie) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovie not implemented")
}
func (UnimplementedMsDatabaseCrudServer) DeleteMovie(context.Context, *Movie) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
}
func (UnimplementedMsDatabaseCrudServer) AddMovieToWatchlist(context.Context, *AddMovieByUser) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovieToWatchlist not implemented")
}
func (UnimplementedMsDatabaseCrudServer) GetAllWatchlistMovies(context.Context, *EmptyMovie) (*Movies, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllWatchlistMovies not implemented")
}
func (UnimplementedMsDatabaseCrudServer) DeleteMovieFromWatchlist(context.Context, *Movie) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMovieFromWatchlist not implemented")
}
func (UnimplementedMsDatabaseCrudServer) CreateReview(context.Context, *NewReview) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedMsDatabaseCrudServer) UpdateReview(context.Context, *Review) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedMsDatabaseCrudServer) mustEmbedUnimplementedMsDatabaseCrudServer() {}

// UnsafeMsDatabaseCrudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsDatabaseCrudServer will
// result in compilation errors.
type UnsafeMsDatabaseCrudServer interface {
	mustEmbedUnimplementedMsDatabaseCrudServer()
}

func RegisterMsDatabaseCrudServer(s grpc.ServiceRegistrar, srv MsDatabaseCrudServer) {
	s.RegisterService(&MsDatabaseCrud_ServiceDesc, srv)
}

func _MsDatabaseCrud_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).CreateUser(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_GetAllMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).GetAllMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/GetAllMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).GetAllMovies(ctx, req.(*EmptyMovie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_GetMovieByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).GetMovieByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/GetMovieByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).GetMovieByCategory(ctx, req.(*MovieCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_AddMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).AddMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/AddMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).AddMovie(ctx, req.(*NewMovie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_DeleteMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Movie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).DeleteMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/DeleteMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).DeleteMovie(ctx, req.(*Movie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_AddMovieToWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMovieByUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).AddMovieToWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/AddMovieToWatchlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).AddMovieToWatchlist(ctx, req.(*AddMovieByUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_GetAllWatchlistMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMovie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).GetAllWatchlistMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/GetAllWatchlistMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).GetAllWatchlistMovies(ctx, req.(*EmptyMovie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_DeleteMovieFromWatchlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Movie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).DeleteMovieFromWatchlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/DeleteMovieFromWatchlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).DeleteMovieFromWatchlist(ctx, req.(*Movie))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewReview)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/CreateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).CreateReview(ctx, req.(*NewReview))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsDatabaseCrud_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsDatabaseCrudServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msproto.MsDatabaseCrud/UpdateReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsDatabaseCrudServer).UpdateReview(ctx, req.(*Review))
	}
	return interceptor(ctx, in, info, handler)
}

// MsDatabaseCrud_ServiceDesc is the grpc.ServiceDesc for MsDatabaseCrud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsDatabaseCrud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msproto.MsDatabaseCrud",
	HandlerType: (*MsDatabaseCrudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MsDatabaseCrud_CreateUser_Handler,
		},
		{
			MethodName: "GetAllMovies",
			Handler:    _MsDatabaseCrud_GetAllMovies_Handler,
		},
		{
			MethodName: "GetMovieByCategory",
			Handler:    _MsDatabaseCrud_GetMovieByCategory_Handler,
		},
		{
			MethodName: "AddMovie",
			Handler:    _MsDatabaseCrud_AddMovie_Handler,
		},
		{
			MethodName: "DeleteMovie",
			Handler:    _MsDatabaseCrud_DeleteMovie_Handler,
		},
		{
			MethodName: "AddMovieToWatchlist",
			Handler:    _MsDatabaseCrud_AddMovieToWatchlist_Handler,
		},
		{
			MethodName: "GetAllWatchlistMovies",
			Handler:    _MsDatabaseCrud_GetAllWatchlistMovies_Handler,
		},
		{
			MethodName: "DeleteMovieFromWatchlist",
			Handler:    _MsDatabaseCrud_DeleteMovieFromWatchlist_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _MsDatabaseCrud_CreateReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _MsDatabaseCrud_UpdateReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msproto/ms.proto",
}
