// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: grpc/pb.proto

package proto

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
	Auction_PutBid_FullMethodName           = "/Auction/PutBid"
	Auction_GetAuctionStatus_FullMethodName = "/Auction/GetAuctionStatus"
	Auction_GetLot_FullMethodName           = "/Auction/GetLot"
	Auction_GetDiscovery_FullMethodName     = "/Auction/GetDiscovery"
)

// AuctionClient is the client API for Auction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// API available at a front-end providing node and between nodes. Should have discovery capabilities.
type AuctionClient interface {
	// Forward a bid to the auction. Returns acknowledgement showing outcome of the bid.
	PutBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*Ack, error)
	// Get the status of the auction. Returns acknowledgement showing the current auction state.
	GetAuctionStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Ack, error)
	// Get details on the item up for auction. Returns lot message with details.
	GetLot(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Lot, error)
	// Discover front-end service nodes for keeping contact with the auction. Returns discovery of node IP addresses.
	GetDiscovery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Discovery, error)
}

type auctionClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionClient(cc grpc.ClientConnInterface) AuctionClient {
	return &auctionClient{cc}
}

func (c *auctionClient) PutBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, Auction_PutBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) GetAuctionStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Ack, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ack)
	err := c.cc.Invoke(ctx, Auction_GetAuctionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) GetLot(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Lot, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Lot)
	err := c.cc.Invoke(ctx, Auction_GetLot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) GetDiscovery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Discovery, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Discovery)
	err := c.cc.Invoke(ctx, Auction_GetDiscovery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServer is the server API for Auction service.
// All implementations must embed UnimplementedAuctionServer
// for forward compatibility.
//
// API available at a front-end providing node and between nodes. Should have discovery capabilities.
type AuctionServer interface {
	// Forward a bid to the auction. Returns acknowledgement showing outcome of the bid.
	PutBid(context.Context, *Bid) (*Ack, error)
	// Get the status of the auction. Returns acknowledgement showing the current auction state.
	GetAuctionStatus(context.Context, *Empty) (*Ack, error)
	// Get details on the item up for auction. Returns lot message with details.
	GetLot(context.Context, *Empty) (*Lot, error)
	// Discover front-end service nodes for keeping contact with the auction. Returns discovery of node IP addresses.
	GetDiscovery(context.Context, *Empty) (*Discovery, error)
	mustEmbedUnimplementedAuctionServer()
}

// UnimplementedAuctionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuctionServer struct{}

func (UnimplementedAuctionServer) PutBid(context.Context, *Bid) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutBid not implemented")
}
func (UnimplementedAuctionServer) GetAuctionStatus(context.Context, *Empty) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuctionStatus not implemented")
}
func (UnimplementedAuctionServer) GetLot(context.Context, *Empty) (*Lot, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLot not implemented")
}
func (UnimplementedAuctionServer) GetDiscovery(context.Context, *Empty) (*Discovery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscovery not implemented")
}
func (UnimplementedAuctionServer) mustEmbedUnimplementedAuctionServer() {}
func (UnimplementedAuctionServer) testEmbeddedByValue()                 {}

// UnsafeAuctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServer will
// result in compilation errors.
type UnsafeAuctionServer interface {
	mustEmbedUnimplementedAuctionServer()
}

func RegisterAuctionServer(s grpc.ServiceRegistrar, srv AuctionServer) {
	// If the following call pancis, it indicates UnimplementedAuctionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auction_ServiceDesc, srv)
}

func _Auction_PutBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).PutBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_PutBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).PutBid(ctx, req.(*Bid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_GetAuctionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).GetAuctionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_GetAuctionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).GetAuctionStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_GetLot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).GetLot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_GetLot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).GetLot(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_GetDiscovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).GetDiscovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auction_GetDiscovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).GetDiscovery(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Auction_ServiceDesc is the grpc.ServiceDesc for Auction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auction",
	HandlerType: (*AuctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutBid",
			Handler:    _Auction_PutBid_Handler,
		},
		{
			MethodName: "GetAuctionStatus",
			Handler:    _Auction_GetAuctionStatus_Handler,
		},
		{
			MethodName: "GetLot",
			Handler:    _Auction_GetLot_Handler,
		},
		{
			MethodName: "GetDiscovery",
			Handler:    _Auction_GetDiscovery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/pb.proto",
}
