package server

import proto "example.com/auction/grpc"

/*
Front-end nodes handle the contact to clients.

Incoming requests from clients should be
forwarded to the current leader of the auction.
*/
type FrontEndNode struct {
	proto.UnimplementedAuctionFrontEndServer
}
