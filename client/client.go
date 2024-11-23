package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	proto "example.com/auction/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var stdIn = setScanner()
var ctx context.Context = context.Background()

var name string
var id int64
var lot *proto.Lot

func main() {
	fmt.Print("Enter your name and press ENTER: ")
	name = nextLine()

	fmt.Print("Enter auction IP-address:port and press ENTER: ")
	addr := nextLine()

	conn := getConnectionToServer(addr)
	defer conn.Close()

	client := proto.NewAuctionClient(conn)

	lot := getLotInfo(client)
	status := getAuctionStatus(client)
	fmt.Printf("The item up for auction is '%s'. \nPrice currently %d,- from bidder '%d'\n", lot.Name, status.Amount, status.BidderId)
	fmt.Println(statusValueToString(status.Result))
	// Client should then obtain auction details and place bids.
}

func statusValueToString(value proto.StatusValue) string {
	switch value {
	case proto.StatusValue_FAULT:
		return "An auction service ERROR occurred."
	case proto.StatusValue_ACCEPTED:
		return "Your bid has been accepted."
	case proto.StatusValue_UNDERBID:
		return "Your bid was REJECTED as too low."
	case proto.StatusValue_NOT_STARTED:
		return "The auction has not started yet."
	case proto.StatusValue_IN_PROGRESS:
		return "The auction is in progress."
	case proto.StatusValue_SOLD:
		return "The auction has closed and the item has been sold."
	case proto.StatusValue_CLOSED:
		return "The auction has closed with no takers."
	default:
		log.Fatalln("ERROR - Unknown proto.StatusValue")
		return "ERROR - Unknown proto.StatusValue"
	}
}

func getLotInfo(client proto.AuctionClient) *proto.Lot {
	lot, err := client.GetLot(ctx, &proto.Empty{})
	if err != nil {
		log.Fatalf("client.GetLot error: %v", err)
	}
	fmt.Println(lot)
	return lot
}

func getAuctionStatus(client proto.AuctionClient) *proto.Ack {
	status, err := client.GetAuctionStatus(ctx, &proto.Empty{})
	if err != nil {
		log.Fatalf("client.GetLot error: %v", err)
	}
	fmt.Println(status)
	return status
}

// Establishes connection to the server.
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
