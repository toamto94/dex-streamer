package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/toamto94/dex-streamer.git/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"strconv"
)

var (
	endpoint = flag.String("e", "localhost", "API endpoint to connect with")
	port     = flag.Int64("p", 50051, "Endpoint port")
	infuraId = flag.String("i", "", "Infura Project ID")
)

func main() {
	flag.Parse()
	connectionString := *endpoint + ":" + strconv.FormatInt(*port, 10)
	fmt.Println(connectionString)
	conn, err := grpc.Dial(connectionString, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to establish connection to server - %v", err)
	} else {
		log.Printf("Connection established")
	}
	defer conn.Close()
	client := proto.NewDEXStreamerClient(conn)

	contract := proto.Contract{
		Endpoint:       "https://mainnet.infura.io/v3/" + *infuraId,
		Chain:          "ethereum",
		Dex:            "uniswapV3",
		Address:        "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640",
		ScrapeInterval: 100000,
	}

	stream, err := client.StreamContract(context.Background(), &contract)

	for {
		tick, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Stream interrupted - %v", err)
		}
		log.Println(tick)
	}
}
