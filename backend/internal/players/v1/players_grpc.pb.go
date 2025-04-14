// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: players/v1/players.proto

package players

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
	Players_CreatePlayer_FullMethodName = "/Players/CreatePlayer"
)

// PlayersClient is the client API for Players service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayersClient interface {
	CreatePlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerReply, error)
}

type playersClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayersClient(cc grpc.ClientConnInterface) PlayersClient {
	return &playersClient{cc}
}

func (c *playersClient) CreatePlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePlayerReply)
	err := c.cc.Invoke(ctx, Players_CreatePlayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayersServer is the server API for Players service.
// All implementations must embed UnimplementedPlayersServer
// for forward compatibility.
type PlayersServer interface {
	CreatePlayer(context.Context, *CreatePlayerRequest) (*CreatePlayerReply, error)
	mustEmbedUnimplementedPlayersServer()
}

// UnimplementedPlayersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPlayersServer struct{}

func (UnimplementedPlayersServer) CreatePlayer(context.Context, *CreatePlayerRequest) (*CreatePlayerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayer not implemented")
}
func (UnimplementedPlayersServer) mustEmbedUnimplementedPlayersServer() {}
func (UnimplementedPlayersServer) testEmbeddedByValue()                 {}

// UnsafePlayersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayersServer will
// result in compilation errors.
type UnsafePlayersServer interface {
	mustEmbedUnimplementedPlayersServer()
}

func RegisterPlayersServer(s grpc.ServiceRegistrar, srv PlayersServer) {
	// If the following call pancis, it indicates UnimplementedPlayersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Players_ServiceDesc, srv)
}

func _Players_CreatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayersServer).CreatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Players_CreatePlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayersServer).CreatePlayer(ctx, req.(*CreatePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Players_ServiceDesc is the grpc.ServiceDesc for Players service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Players_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Players",
	HandlerType: (*PlayersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlayer",
			Handler:    _Players_CreatePlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "players/v1/players.proto",
}
