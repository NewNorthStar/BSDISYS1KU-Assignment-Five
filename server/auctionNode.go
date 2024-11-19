package server

import proto "example.com/auction/grpc"

/*
Auction nodes either conduct the auction as the leader,
or follow the auction for redundancy.

Auction nodes should be comparable by logical time.
If a new leader needs to be elected, the winner should be
(among) the most up-to-date.
*/
type AuctionNode struct {
	proto.UnimplementedAuctionServer
}
