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
		Endpoint:       "https://mainnet.infura.io/v3/4821cb96059a4d9eb05435da1f54fdad",
		ChainId:        "b",
		Address:        "0x4585FE77225b41b697C938B018E2Ac67Ac5a20c0",
		Abi:            "d",
		ScrapeInterval: 900,
	}

	stream, err := client.StreamContract(context.Background(), &contract)

	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Stream interrupted - %v", err)
		}
		log.Println(feature)
	}
}
