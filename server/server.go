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

// Sets the closing time of the auction from time of creation.
const AUCTION_DURATION = time.Second * 120

/*
Args[1] = My endpoint to serve on.

Args[2] (optional) = Network endpoint of leader.
*/
func main() {
	if len(os.Args) == 2 {
		runAsLeader(os.Args[1])
	} else if len(os.Args) == 3 {
		runAsFollower(os.Args[1], os.Args[2])
	}

	log.Fatalln("*** Incorrect args ***")
}

/*
Start this node as a new auction.
*/
func runAsLeader(addr string) {
	log.Println("Starting a new auction...")

	auction := newAuctionService()
	auction.listener = listenOn(addr)
	auction.lead_status = true
	auction.leader = auction.listener.Addr()
	auction.closing_time = time.Now().Add(AUCTION_DURATION)
	auction.startService()
}

/*
Connect this node to an existing auction.
*/
func runAsFollower(addr, leaderAddr string) {
	log.Println("Connecting to existing auction...")

	auction := newAuctionService()
	auction.listener = listenOn(addr)
	auction.lead_status = false
	auction.registerAsFollower(leaderAddr)
	auction.startService()
}

// Establishes a client connection to a server.
func getConnectionToServer(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to service: %v", err)
	}
	return conn
}

// Obtains a TCP listener on a given network address.
func listenOn(address string) net.Listener {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to obtain TCP listener: %v\n", err)
	}
	return listener
}

/*
Contains the data necessary to conduct an auction.

An instance can lead an auction, or follow an auction.
*/
type AuctionService struct {
	proto.UnimplementedAuctionServer

	listener    net.Listener    // Network listener for this instance.
	known_nodes map[string]bool // Set of known IP addresses to the auction.
	lead_status bool            // Set to true when instance is leading.
	leader      net.Addr        // Leader network end-point for bid forwarding.

	bid_lock sync.Mutex // The state of bidding is protected by a mutex.
	top_bid  *proto.Bid
	bid_time int64 // Lamport timestamp that increases with each top bid.
	bidders  int64 // Count of bidders for ID assignment.

	item_name    string
	asking_price int64
	starting_bid int64

	closing_time time.Time
}

/*
Creates and returns an auction struct.

Hardcoded for selling a 'course book'. :)
*/
func newAuctionService() *AuctionService {
	return &AuctionService{
		known_nodes: make(map[string]bool),
		top_bid: &proto.Bid{
			Amount:   0,
			BidderId: 0,
		},
		bid_time: 0,
		bidders:  0,

		item_name:    "Course Book",
		asking_price: 80,
		starting_bid: 40,
	}
}

/*
True when receiver is a leader.
*/
func (auction *AuctionService) isLeader() bool {
	return auction.lead_status
}

/*
True when receiver is a follower.
*/
func (auction *AuctionService) isFollower() bool {
	return !auction.lead_status
}

/*
Starts the AuctionService server. Always called by node starting functions.
*/
func (auction *AuctionService) startService() {
	auction.known_nodes[auction.listener.Addr().String()] = true // Makes address available to GetDiscovery RPC.

	go auction.closingCallRoutine()

	grpcServer := grpc.NewServer()
	proto.RegisterAuctionServer(grpcServer, auction)
	log.Printf("Now ready at %s\n", auction.listener.Addr().String())
	err := grpcServer.Serve(auction.listener)
	if err != nil {
		log.Fatalf("Service failure: %v\n", err)
	}
}

/*
Registers receiver as a follower with the existing auction.

Obtains the information necessary to replicate the auction
and signals to the auction leader, that it should regularly
update this follower.
*/
func (auction *AuctionService) registerAsFollower(leaderAddr string) {
	conn := getConnectionToServer(leaderAddr)
	defer conn.Close()
	client := proto.NewAuctionClient(conn)

	lot, err := client.Register(ctx, &proto.Node{
		Addr: auction.listener.Addr().String(),
	})
	if err != nil {
		log.Fatalf("Unable to register: %v\n", err)
	}
	auction.item_name = lot.Name
	auction.asking_price = lot.AskingPrice
	auction.starting_bid = lot.StartingBid
	auction.leader, err = net.ResolveTCPAddr("tcp", leaderAddr)
	if err != nil {
		log.Fatalf("Unable to resolve address: %v\n", err)
	}
	auction.closing_time = lot.ClosingTime.AsTime().In(time.Local)
	log.Printf("Registered as follower with %s. Now awaiting first update...\n", leaderAddr)
}

/*
Call as goroutine. Logs the auction outcome after closing time.
*/
func (auction *AuctionService) closingCallRoutine() {
	for time.Now().Before(auction.closing_time) {
		time.Sleep(2 * time.Second)
	}
	log.Printf("Auction was closed at %s\n", auction.closing_time.String())
	if auction.top_bid.BidderId != 0 {
		log.Printf("The item '%s' was sold to bidder #%d for %d,-\n", auction.item_name, auction.top_bid.BidderId, auction.top_bid.Amount)
	} else {
		log.Printf("There were no bidders for the item '%s'\n", auction.item_name)
	}
}

/*
Followers use this method to place a bid with the leader.
Uses a temporary client to forward the bid.
*/
func (auction *AuctionService) forwardBid(ctx context.Context, msg *proto.Bid) (*proto.Ack, error) {
	response, err := auction.forwardAttempt(ctx, msg)
	if err != nil {
		log.Printf("Leader unavailable, proposing election: %v", err)
		auction.holdElection()
		time.Sleep(time.Millisecond * 200)
		response, err = auction.forwardAttempt(ctx, msg)
	}
	return response, err
}

func (auction *AuctionService) forwardAttempt(ctx context.Context, msg *proto.Bid) (*proto.Ack, error) {
	conn, err := grpc.NewClient(auction.leader.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "grpc.NewClient: %s", err.Error())
	}
	defer conn.Close()
	forwardingClient := proto.NewAuctionClient(conn)
	log.Printf("Forwarding bid from #%d\n", msg.BidderId)
	response, err := forwardingClient.PutBid(ctx, msg)
	return response, err
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
	auction.updateAllFollowers()
	return proto.StatusValue_ACCEPTED
}

/*
Method called by the leader only. Updates all followers.

Called after a bid has been placed.
*/
func (auction *AuctionService) updateAllFollowers() {
	var group sync.WaitGroup
	group_tasks := make([]func(), 0)
	for addr := range auction.known_nodes {
		if addr == auction.listener.Addr().String() {
			continue
		}
		group.Add(1)
		group_tasks = append(group_tasks, func() {
			defer group.Done()
			auction.sendUpdateToFollower(addr)
		})
	}
	for _, call := range group_tasks {
		go call()
	}
	group.Wait()
}

/*
Called by leader to update a newly registered follower node.

If successful, all nodes will be updated to discover the new follower.
On failure, the follower will not be added to the auction.
This outcome is logged.
*/
func (auction *AuctionService) firstUpdate(addr string) {
	var err error
	for i := 0; i < 3; i++ {
		err = auction.sendUpdateToFollower(addr)
		if err == nil {
			log.Printf("Successful first update on follower node: %s\n", addr)
			auction.known_nodes[addr] = true
			auction.updateAllFollowers()
			return
		}
	}
	log.Printf("FAILED first update on follower node: %s\n", addr)
}

/*
Updates a single follower node using a temporary client instance.
*/
func (auction *AuctionService) sendUpdateToFollower(addr string) error {
	conn := getConnectionToServer(addr)
	defer conn.Close()
	client := proto.NewAuctionClient(conn)
	_, err := client.UpdateNode(ctx, &proto.NodeState{
		TopBid:  auction.top_bid,
		BidTime: auction.bid_time,
		Bidders: auction.bidders,
		Addr:    auction.listener.Addr().String(),
		Discovery: &proto.Discovery{
			Others: keys(auction.known_nodes),
			Leader: auction.leader.String(),
		},
	})
	return err
}

func keys(m map[string]bool) []string {
	s := make([]string, 0, len(m))
	for key, ok := range m {
		if ok {
			s = append(s, key)
		}
	}

	return s
}

/*
Hold an election among follower nodes.

Called when a follower determines that the leader is unreachable.
*/
func (auction *AuctionService) holdElection() {
	log.Printf("Election triggered...")

	var group sync.WaitGroup
	negativeAnswer := false
	group_tasks := make([]func(), 0)
	for addr := range auction.known_nodes {
		if addr == auction.listener.Addr().String() {
			continue
		}
		group.Add(1)
		group_tasks = append(group_tasks, func() {
			defer group.Done()
			answer, err := auction.sendElectionToFollower(addr)
			if err != nil {
				log.Printf("Error from addr %s: %s", addr, err.Error())
				return
			}
			if answer.Result == proto.StatusValue_REJECTED {
				negativeAnswer = true
			}
		})
	}
	for _, call := range group_tasks {
		go call()
	}
	group.Wait()
	if negativeAnswer {
		log.Printf("Election lost to another node.")
	}

	log.Printf("Won the election!")
	auction.lead_status = true
	auction.leader = auction.listener.Addr()
	auction.announceCoordinator()
	auction.updateAllFollowers()
}

/*
Sends an Election call to a single node.
*/
func (auction *AuctionService) sendElectionToFollower(addr string) (*proto.ElectionAnswer, error) {
	conn := getConnectionToServer(addr)
	defer conn.Close()
	client := proto.NewAuctionClient(conn)
	limitedContext, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	answer, err := client.Election(limitedContext, &proto.ElectionBallot{
		BidTime: auction.bid_time,
		Addr:    auction.listener.Addr().String(),
	})
	return answer, err
}

/*
Comparison method for the incoming election ballot vs the receiver.
*/
func (auction *AuctionService) ballotIsBetterCandidate(ballot *proto.ElectionBallot) bool {
	if auction.bid_time != ballot.BidTime {
		return auction.bid_time < ballot.BidTime
	} else {
		return auction.listener.Addr().String() > ballot.Addr
	}
}

/*
Declares election victory to all nodes.
*/
func (auction *AuctionService) announceCoordinator() {
	var group sync.WaitGroup
	group_tasks := make([]func(), 0)
	for addr := range auction.known_nodes {
		if addr == auction.listener.Addr().String() {
			continue
		}
		group.Add(1)
		group_tasks = append(group_tasks, func() {
			defer group.Done()
			err := auction.sendCoordinatorToFollower(addr)
			if err != nil {
				log.Printf("Error from addr %s: %s", addr, err.Error())
				return
			}
		})
	}
	for _, call := range group_tasks {
		go call()
	}
	group.Wait()
}

/*
Sends a Coordinator call to a single node.
*/
func (auction *AuctionService) sendCoordinatorToFollower(addr string) error {
	conn := getConnectionToServer(addr)
	defer conn.Close()
	client := proto.NewAuctionClient(conn)
	limitedContext, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	_, err := client.Coordinator(limitedContext, &proto.ElectionBallot{
		BidTime: auction.bid_time,
		Addr:    auction.listener.Addr().String(),
	})
	return err
}

//// RPC SERVICE METHODS ////

/*
gRPC method.

Forward a bid to the auction. Returns acknowledgement showing outcome of the bid.
Calls with bid ID '0' will register the bidder and return a new ID.

Followers will forward the bid to the leader. This may fail and trigger an election.
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

/*
gRPC method.

Get the status of the auction. The returned Ack message shows who currently holds the top bid.
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

/*
gRPC method.

Returns a message with generic auction details.
*/
func (auction *AuctionService) GetLot(ctx context.Context, msg *proto.Empty) (*proto.Lot, error) {
	answer := &proto.Lot{
		Id:          1,
		Name:        auction.item_name,
		AskingPrice: auction.asking_price,
		StartingBid: auction.starting_bid,
		CurrentBid:  auction.top_bid.Amount,
		ClosingTime: timestamppb.New(auction.closing_time),
	}
	return answer, nil
}

/*
gRPC method.

Returns the leader IP address "Leader", and a list of known endpoints "Others".
*/
func (auction *AuctionService) GetDiscovery(ctx context.Context, msg *proto.Empty) (*proto.Discovery, error) {
	answer := &proto.Discovery{
		Others: keys(auction.known_nodes),
		Leader: auction.leader.String(),
	}
	return answer, nil
}

/*
gRPC method.

Updates the follower node with new auction data. Returns an error if the node is a leader.
*/
func (auction *AuctionService) UpdateNode(ctx context.Context, msg *proto.NodeState) (*proto.Empty, error) {
	if auction.isLeader() {
		return nil, status.Errorf(codes.PermissionDenied, "This node is leading the auction")
	}

	auction.bid_lock.Lock()
	defer auction.bid_lock.Unlock()
	auction.top_bid = msg.TopBid
	auction.bid_time = msg.BidTime
	auction.bidders = msg.Bidders
	var err error
	auction.leader, err = net.ResolveTCPAddr("tcp", msg.Discovery.Leader)
	if err != nil {
		log.Printf("Unable to resolve leader address %v\n", err)
	}
	auction.known_nodes = make(map[string]bool)
	for _, addr := range msg.Discovery.Others {
		auction.known_nodes[addr] = true
	}

	log.Printf("Received update from %s\n", msg.Addr)
	log.Printf("PutBid op#%d: Bid now at %d,- held by bidder #%d\n", auction.bid_time, auction.top_bid.Amount, auction.top_bid.BidderId)
	return &proto.Empty{}, nil
}

/*
gRPC method.

Register a follower node. If the leader is able to accept this, it will return a Lot message.
*/
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

/*
gRPC method.

Ping a node to see that it is still active.
*/
func (auction *AuctionService) Ping(ctx context.Context, msg *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

/*
gRPC method.

Elect a new leader.

- Followers LESS qualified than the node on the ballot may answer with ACCEPT. An error is treated as an ACCEPT.

- Followers MORE qualified than the node on the ballot MUST answer with REJECT. The follower must then hold its own election.
*/
func (auction *AuctionService) Election(ctx context.Context, ballot *proto.ElectionBallot) (*proto.ElectionAnswer, error) {
	if auction.ballotIsBetterCandidate(ballot) {
		return &proto.ElectionAnswer{
			Result: proto.StatusValue_ACCEPTED,
		}, nil
	} else {
		defer auction.holdElection()
		return &proto.ElectionAnswer{
			Result: proto.StatusValue_REJECTED,
		}, nil
	}
}

/*
gRPC method.

Called by the winner of an election. All nodes resolve the new leader address. Leaders receiving this call will step down.
*/
func (auction *AuctionService) Coordinator(ctx context.Context, ballot *proto.ElectionBallot) (*proto.Empty, error) {
	if !auction.ballotIsBetterCandidate(ballot) {
		return nil, status.Errorf(codes.PermissionDenied, "This node is better than the coordinator")
	}
	auction.lead_status = false
	var err error
	auction.leader, err = net.ResolveTCPAddr("tcp", ballot.Addr)
	if err != nil {
		log.Println(err.Error())
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	auction.lead_status = false

	return &proto.Empty{}, nil
}
