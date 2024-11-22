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

	client := proto.NewAuctionFrontEndClient(conn)

	lot = getLotInfo(client)
	fmt.Println(lot)
	// Client should then obtain auction details and place bids.
}

func getLotInfo(client proto.AuctionFrontEndClient) *proto.Lot {
	lot, err := client.GetLot(ctx, &proto.Empty{})
	if err != nil {
		log.Fatalf("client.GetLot error: %v", err)
	}
	return lot
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
