package main

import (
	"context"
	"log"
	"net"
	"sync"

	proto "example.com/auction/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
Auction nodes either conduct the auction as the leader,
or follow the auction for redundancy.

Auction nodes should be comparable by logical time.
If a new leader needs to be elected, the winner should be
(among) the most up-to-date.
*/
type AuctionService struct {
	proto.UnimplementedAuctionServer

	top_bid  *proto.Bid
	bid_time int64      // Lamport timestamp that increases with each top bid.
	bid_lock sync.Mutex // Controls access to bid.

	name         string
	asking_price int64
	starting_bid int64
}

/*
Creates and returns an auction struct.
Simply hardcoded to be auctioning a course book.
*/
func newAuctionService() *AuctionService {
	return &AuctionService{
		top_bid: &proto.Bid{
			Amount:   0,
			BidderId: 0,
		},
		bid_time: 0,

		name:         "Course Book",
		asking_price: 80,
		starting_bid: 40,
	}
}

func main() {
	service := newAuctionService()
	listener := listenOn("localhost:5050")

	service.startService(listener)
}

// Obtains a TCP listener on a given network address.
func listenOn(address string) net.Listener {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	return listener
}

// Begins serving ChittyChat gRPC service as a goroutine and returns.
func (auction *AuctionService) startService(listener net.Listener) {
	grpcServer := grpc.NewServer()
	proto.RegisterAuctionServer(grpcServer, auction)
	log.Printf("Auction on at %v\n", listener.Addr())
	err := grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Service failure: %v\n", err)
	}
}

func (auction *AuctionService) GetAuctionStatus(ctx context.Context, msg *proto.Empty) (*proto.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuctionStatus not implemented")
}

func (auction *AuctionService) GetDiscovery(ctx context.Context, msg *proto.Empty) (*proto.Discovery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiscovery not implemented")
}

func (auction *AuctionService) GetLot(ctx context.Context, msg *proto.Empty) (*proto.Lot, error) {
	answer := &proto.Lot{
		Id:          1,
		Name:        auction.name,
		AskingPrice: auction.asking_price,
		StartingBid: auction.starting_bid,
		CurrentBid:  auction.top_bid.Amount,
	}
	return answer, nil
}

func (auction *AuctionService) PutBid(ctx context.Context, msg *proto.Bid) (*proto.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutBid not implemented")
}

/*
Process an incoming bid against the current state of the auction.

Returns 'true' if incoming bid is now the highest and 'false' otherwise.
*/
func (auction *AuctionService) processBid(incoming *proto.Bid) bool {
	if incoming.Amount <= auction.top_bid.Amount ||
		incoming.Amount < auction.starting_bid { // Pre-lock check. Under-bids cannot obtain lock.
		return false
	}

	auction.bid_lock.Lock()
	defer auction.bid_lock.Unlock()

	if incoming.Amount <= auction.top_bid.Amount { // Second check once lock is obtained.
		return false
	} else {
		auction.top_bid = incoming
		auction.bid_time++
		return true
	}
}
