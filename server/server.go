package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"
	"time"

	proto "example.com/auction/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ctx context.Context = context.Background()

/*
Args[1] = My web socket.

Args[2] (optional) = Leader to follow.
*/
func main() {
	if len(os.Args) == 2 {
		runAsLeader()
	} else if len(os.Args) == 3 {
		runAsFollower()
	}

	log.Fatalln("*** Incorrect args ***")
}

func runAsLeader() {
	log.Println("Starting a new auction...")

	auction := newAuctionService()
	auction.listener = listenOn(os.Args[1])
	auction.leadStatus = true
	auction.leader = auction.listener.Addr()
	auction.closing_time = time.Now().Add(time.Second * 120)
	auction.startService()
}

func runAsFollower() {
	log.Println("Following an existing auction...")

	auction := newAuctionService()
	auction.listener = listenOn(os.Args[1])
	auction.leadStatus = false
	auction.registerAsFollower(os.Args[2])
	auction.startService()
}

func (auction *AuctionService) registerAsFollower(leaderAddr string) {
	conn := getConnectionToServer(leaderAddr)
	defer conn.Close()
	client := proto.NewAuctionClient(conn)

	lot, err := client.Register(ctx, &proto.Node{
		Addr: auction.listener.Addr().String(),
	})
	if err != nil {
		log.Fatalf("Replication error: %v\n", err)
	}
	auction.name = lot.Name
	auction.asking_price = lot.AskingPrice
	auction.starting_bid = lot.StartingBid
	auction.leader, err = net.ResolveTCPAddr("tcp", leaderAddr)
	if err != nil {
		log.Fatalf("Replication error: %v\n", err)
	}
	auction.closing_time = lot.ClosingTime.AsTime().In(time.Local)
	log.Printf("Registered as follower with %s, awaiting first update...\n", leaderAddr)
}

// Establishes connection to a server.
func getConnectionToServer(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("getConnectionToServer error: %v", err)
	}
	return conn
}

/*
Auction nodes either conduct the auction as the leader,
or follow the auction for redundancy.

Auction nodes should be comparable by logical time.
If a new leader needs to be elected, the winner should be
(among) the most up-to-date.
*/
type AuctionService struct {
	proto.UnimplementedAuctionServer
	listener net.Listener

	known_nodes []string
	leadStatus  bool
	leader      net.Addr

	top_bid  *proto.Bid
	bid_time int64      // Lamport timestamp that increases with each top bid.
	bid_lock sync.Mutex // Controls access to bid.
	bidders  int64

	name         string
	asking_price int64
	starting_bid int64

	closing_time time.Time
}

func (auction *AuctionService) isLeader() bool {
	return auction.leadStatus
}

func (auction *AuctionService) isFollower() bool {
	return !auction.leadStatus
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

// Obtains a TCP listener on a given network address.
func listenOn(address string) net.Listener {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	return listener
}

// Connects and begins the auction service.
func (auction *AuctionService) startService() {
	auction.known_nodes = append(auction.known_nodes, auction.listener.Addr().String()) // Makes address available to GetDiscovery RPC.

	go func() {
		for time.Now().Before(auction.closing_time) {
			time.Sleep(2 * time.Second)
		}
		log.Printf("Auction was closed at %s\n", auction.closing_time.String())
		if auction.top_bid.BidderId != 0 {
			log.Printf("The item '%s' was sold to bidder #%d for %d,-\n", auction.name, auction.top_bid.BidderId, auction.top_bid.Amount)
		} else {
			log.Printf("There were no bidders for the item '%s'\n", auction.name)
		}

	}()

	grpcServer := grpc.NewServer()
	proto.RegisterAuctionServer(grpcServer, auction)
	log.Printf("Now ready at %s\n", auction.listener.Addr().String())
	err := grpcServer.Serve(auction.listener)
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
	var result proto.StatusValue
	if time.Now().After(auction.closing_time) {
		if auction.top_bid.BidderId != 0 {
			result = proto.StatusValue_SOLD
		} else {
			result = proto.StatusValue_CLOSED
		}
	} else {
		result = proto.StatusValue_IN_PROGRESS
	}

	answer := &proto.Ack{
		Amount:   lead.Amount,
		BidderId: lead.BidderId,
		Result:   result,
	}
	return answer, nil
}

// Discover front-end service nodes for keeping contact with the auction. Returns discovery of node IP addresses.
func (auction *AuctionService) GetDiscovery(ctx context.Context, msg *proto.Empty) (*proto.Discovery, error) {
	answer := &proto.Discovery{
		IpAddresses: auction.known_nodes,
		Leader:      auction.leader.String(),
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
		ClosingTime: timestamppb.New(auction.closing_time),
	}
	return answer, nil
}

/*
Forward a bid to the auction. Returns acknowledgement showing outcome of the bid.

The returned BidderId allows the service to assign IDs to clients.
*/
func (auction *AuctionService) PutBid(ctx context.Context, msg *proto.Bid) (*proto.Ack, error) {
	if time.Now().After(auction.closing_time) { // If the auction is over, we instead return the auction status.
		return auction.GetAuctionStatus(ctx, &proto.Empty{})
	}
	if auction.isFollower() {
		return auction.forwardBid(ctx, msg)
	}

	result := auction.processBid(msg)

	log.Printf("PutBid op#%d: Bid now at %d,- held by bidder #%d\n", auction.bid_time, auction.top_bid.Amount, auction.top_bid.BidderId)

	answer := &proto.Ack{
		Amount:   auction.top_bid.Amount,
		Result:   result,
		BidderId: msg.BidderId,
	}
	return answer, nil
}

func (auction *AuctionService) forwardBid(ctx context.Context, msg *proto.Bid) (*proto.Ack, error) {
	conn, err := grpc.NewClient(auction.leader.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Error forwarding the bid: %s", err.Error())
	}
	defer conn.Close()
	forwardingClient := proto.NewAuctionClient(conn)
	log.Printf("Forwarded bid from #%d\n", msg.BidderId)
	return forwardingClient.PutBid(ctx, msg)
}

/*
Process an incoming bid against the current state of the auction.

Returns 'true' if incoming bid is now the highest and 'false' otherwise.
*/
func (auction *AuctionService) processBid(bid *proto.Bid) proto.StatusValue {
	if bid.Amount <= auction.top_bid.Amount ||
		bid.Amount < auction.starting_bid { // Pre-lock check. Under-bids cannot obtain lock.
		return proto.StatusValue_UNDERBID
	}

	auction.bid_lock.Lock()
	defer auction.bid_lock.Unlock()
	auction.bid_time++

	if bid.Amount <= auction.top_bid.Amount { // Second check once lock is obtained.
		return proto.StatusValue_UNDERBID
	}
	if bid.BidderId == 0 { // New bidders get an ID assigned and returned.
		auction.bidders++
		bid.BidderId = auction.bidders
	}
	auction.top_bid = bid
	return proto.StatusValue_ACCEPTED
}

// Ping a node to see that it is still active.
func (auction *AuctionService) Ping(ctx context.Context, msg *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

// Register a follower node. If the leader is able to accept this, it will return a Lot message.
func (auction *AuctionService) Register(ctx context.Context, msg *proto.Node) (*proto.Lot, error) {
	if auction.isFollower() {
		return nil, status.Errorf(codes.PermissionDenied, "Followers must register with the leader: %s", auction.leader)
	}
	if time.Now().After(auction.closing_time) { // If the auction is over, new nodes cannot register.
		return nil, status.Errorf(codes.Unavailable, "This auction has closed.")
	}

	log.Printf("New follower node: %s\n", msg.Addr)
	go auction.firstUpdate(msg.Addr)
	return auction.GetLot(ctx, &proto.Empty{})
}

func (auction *AuctionService) firstUpdate(addr string) {
	var err error
	for i := 0; i < 3; i++ {
		err = auction.sendUpdateToFollower(addr)
		if err == nil {
			log.Printf("Successful first update on follower node: %s\n", addr)
			auction.known_nodes = append(auction.known_nodes, addr)
			return
		}
	}
	log.Printf("FAILED first update on follower node: %s\n", addr)
}

func (auction *AuctionService) sendUpdateToFollower(addr string) error {
	conn := getConnectionToServer(addr)
	defer conn.Close()
	client := proto.NewAuctionClient(conn)
	_, err := client.UpdateNode(ctx, &proto.NodeState{
		TopBid:  auction.top_bid,
		BidTime: auction.bid_time,
		Bidders: auction.bidders,
		Addr:    auction.listener.Addr().String(),
	})
	return err
}

// Updates a follower node. The follower returns a confirmation when this has happened.
func (auction *AuctionService) UpdateNode(ctx context.Context, msg *proto.NodeState) (*proto.Empty, error) {
	if auction.isLeader() {
		return nil, status.Errorf(codes.PermissionDenied, "This node is leading the auction")
	}

	auction.bid_lock.Lock()
	defer auction.bid_lock.Unlock()
	auction.top_bid = msg.TopBid
	auction.bid_time = msg.BidTime
	auction.bidders = msg.Bidders
	log.Printf("Received update from %s\n", msg.Addr)
	return &proto.Empty{}, nil
}
