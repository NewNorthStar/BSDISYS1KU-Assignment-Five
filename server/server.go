package main

import (
	"context"
	"log"
	"net"
	"sync"

	proto "example.com/auction/grpc"
	"google.golang.org/grpc"
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

	known_nodes []string

	top_bid  *proto.Bid
	bid_time int64      // Lamport timestamp that increases with each top bid.
	bid_lock sync.Mutex // Controls access to bid.
	bidders  int64

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
		known_nodes: make([]string, 0),
		top_bid: &proto.Bid{
			Amount:   0,
			BidderId: 0,
		},
		bid_time: 0,
		bidders:  0,

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

// Connects and begins the auction service.
func (auction *AuctionService) startService(listener net.Listener) {
	auction.known_nodes = append(auction.known_nodes, listener.Addr().String()) // Makes address available to GetDiscovery RPC.

	grpcServer := grpc.NewServer()
	proto.RegisterAuctionServer(grpcServer, auction)
	log.Printf("Auction on at %s\n", listener.Addr().String())
	err := grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Service failure: %v\n", err)
	}
}

/*
Get the status of the auction. Returns acknowledgement showing the current auction state.

The returned BidderId shows who currently holds the top bid.
*/
func (auction *AuctionService) GetAuctionStatus(ctx context.Context, msg *proto.Empty) (*proto.Ack, error) {
	lead := auction.top_bid
	answer := &proto.Ack{
		Amount:   lead.Amount,
		BidderId: lead.BidderId,
		Result:   proto.StatusValue_IN_PROGRESS,
	}
	return answer, nil
}

// Discover front-end service nodes for keeping contact with the auction. Returns discovery of node IP addresses.
func (auction *AuctionService) GetDiscovery(ctx context.Context, msg *proto.Empty) (*proto.Discovery, error) {
	answer := &proto.Discovery{
		IpAddresses: auction.known_nodes,
	}
	return answer, nil
}

// Get details on the item up for auction. Returns lot message with details.
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

/*
Forward a bid to the auction. Returns acknowledgement showing outcome of the bid.

The returned BidderId allows the service to assign IDs to clients.
*/
func (auction *AuctionService) PutBid(ctx context.Context, msg *proto.Bid) (*proto.Ack, error) {
	if msg.BidderId == 0 {
		auction.bidders++
		msg.BidderId = auction.bidders
	}
	result := auction.processBid(msg)
	answer := &proto.Ack{
		Amount:   auction.top_bid.Amount,
		Result:   result,
		BidderId: msg.BidderId,
	}
	return answer, nil
}

/*
Process an incoming bid against the current state of the auction.

Returns 'true' if incoming bid is now the highest and 'false' otherwise.
*/
func (auction *AuctionService) processBid(incoming *proto.Bid) proto.StatusValue {
	if incoming.Amount <= auction.top_bid.Amount ||
		incoming.Amount < auction.starting_bid { // Pre-lock check. Under-bids cannot obtain lock.
		return proto.StatusValue_UNDERBID
	}

	auction.bid_lock.Lock()
	defer auction.bid_lock.Unlock()

	if incoming.Amount <= auction.top_bid.Amount { // Second check once lock is obtained.
		return proto.StatusValue_UNDERBID
	} else {
		auction.top_bid = incoming
		auction.bid_time++
		return proto.StatusValue_ACCEPTED
	}
}
