package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	proto "example.com/auction/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var stdIn = setScanner()
var ctx context.Context = context.Background()

var my_id int64 = 0
var lot *proto.Lot
var discovery *proto.Discovery

var conn *grpc.ClientConn
var client proto.AuctionClient

func main() {
	fmt.Print("Enter auction IP-address:port and press ENTER: ")
	addr := nextLine()

	conn = getConnectionToServer(addr)
	defer conn.Close()
	client = proto.NewAuctionClient(conn)

	getAndShowAuctionDetails(client)

	interactionLoop()
}

/*
User-interaction loop. Parses the amount to bid and sends it to the auction.
*/
func interactionLoop() {
	for {
		var err error
		discovery, err = client.GetDiscovery(ctx, &proto.Empty{})
		if err != nil {
			log.Println(err.Error())
		}

		fmt.Print("To bid, enter an amount and press ENTER: ")
		amount, err := strconv.ParseInt(nextLine(), 10, 0)
		if err != nil {
			log.Println("*** Unable to parse amount. ***")
			continue
		}

		ack, err := client.PutBid(ctx, &proto.Bid{
			Amount:   amount,
			BidderId: my_id,
		})

		if err != nil {
			log.Println(err.Error())
			findAlternativeClient()
			time.Sleep(time.Millisecond * 200)
			ack, err = client.PutBid(ctx, &proto.Bid{
				Amount:   amount,
				BidderId: my_id,
			})
		}

		if err != nil {
			continue
		}

		if my_id == 0 && (ack.Result == proto.StatusValue_ACCEPTED || ack.Result == proto.StatusValue_UNDERBID) {
			my_id = ack.BidderId
			fmt.Printf("You are bidder #%d\n", my_id)
		}

		displayAcknowledge(ack)
	}
}

func findAlternativeClient() {
	for _, addr := range discovery.Others {
		conn.Close()
		conn = getConnectionToServer(addr)
		client = proto.NewAuctionClient(conn)
		_, err := client.Ping(ctx, &proto.Empty{})
		if err == nil {
			return
		}
	}

	panic("No alternative connections found!")
}

/*
Info gathering and display method for when first connecting to the auction.
*/
func getAndShowAuctionDetails(client proto.AuctionClient) {
	var err error
	lot, err = client.GetLot(ctx, &proto.Empty{})
	if err != nil {
		log.Fatalf("client.GetLot error: %v", err)
	}
	var ack *proto.Ack
	ack, err = client.GetAuctionStatus(ctx, &proto.Empty{})
	if err != nil {
		log.Fatalf("client.GetAuctionStatus error: %v", err)
	}
	fmt.Printf("The item up for auction is '%s'. \nPrice currently %d,- from bidder #%d\n", lot.Name, ack.Amount, ack.BidderId)
	fmt.Printf("The auction closes at %s\n", lot.ClosingTime.AsTime().In(time.Local).String())
	if ack.Amount < lot.StartingBid {
		fmt.Printf("The starting bid is %d,-\n", lot.StartingBid)
	}
	displayAcknowledge(ack)
}

/*
Display method for AuctionService ack messages.
*/
func displayAcknowledge(ack *proto.Ack) {
	switch ack.Result {
	case proto.StatusValue_FAULT:
		fmt.Printf("An auction service ERROR occurred.\n")
	case proto.StatusValue_ACCEPTED:
		fmt.Printf("Your bid has been accepted. The price is now %d,-\n", ack.Amount)
	case proto.StatusValue_UNDERBID:
		fmt.Printf("Your bid has been REJECTED as too low. The current top bid is %d,-\n", ack.Amount)
	case proto.StatusValue_NOT_STARTED:
		fmt.Printf("The auction has not started yet.\n")
	case proto.StatusValue_IN_PROGRESS:
		fmt.Printf("The auction is in progress.\n")
	case proto.StatusValue_SOLD:
		fmt.Printf("The auction has closed and the item has been sold to bidder #%d for %d,-\n", ack.BidderId, ack.Amount)
		if ack.BidderId == my_id {
			fmt.Println("You won the auction!")
		}
	case proto.StatusValue_CLOSED:
		fmt.Printf("The auction has closed with no takers.\n")
	default:
		log.Fatalf("ERROR - Unknown proto.StatusValue\n")
	}
}

// Establishes connection to a server.
func getConnectionToServer(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("getConnectionToServer error: %v", err)
	}
	return conn
}

// Setup for stdIn (input from console). Any scanner settings go here.
func setScanner() *bufio.Scanner {
	var sc = bufio.NewScanner(os.Stdin)
	return sc
}

// Obtains next line from stdIn.
func nextLine() string {
	stdIn.Scan()
	return stdIn.Text()
}
