package main

import (
	"sync"

	proto "example.com/auction/grpc"
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

func main() {

}

/*
Process an incoming bid against the current state of the auction.

Returns 'true' if incoming bid is now the highest and 'false' otherwise.
*/
func (auction *AuctionService) processBid(incoming *proto.Bid) bool {
	if incoming.Amount <= auction.top_bid.Amount { // Pre-lock check. Under-bids cannot obtain lock.
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
