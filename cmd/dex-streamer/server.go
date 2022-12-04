package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/toamto94/dex-streamer.git/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"math/big"
	"net"
	"time"
)

var (
	tls      = flag.Bool("tls", false, "Choose between TLS and pure TCP")
	certFile = flag.String("cert_file", "", "TLS cert file")
	keyFile  = flag.String("key_file", "", "TLS key file")
	port     = flag.Int("port", 50051, "Server Port")
)

type DEXStreamerServerImp struct {
	proto.UnimplementedDEXStreamerServer
}

func (server *DEXStreamerServerImp) StreamContract(contract *proto.Contract, stream proto.DEXStreamer_StreamContractServer) error {
	client, err := ethclient.Dial(contract.Endpoint)
	if err != nil {
		log.Fatalf("EVM endpoint could not be stablished - %v", err)
	} else {
		log.Printf("Connection to EVM endpoint established")
	}
	address := common.HexToAddress(contract.Address)
	sender := common.HexToAddress("0x0000000000000000000000000000000000000000")
	_ = client
	done := make(chan bool)
	ticker := time.NewTicker(time.Millisecond * time.Duration(contract.ScrapeInterval))
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Println("Done")
			return nil
		case <-ticker.C:
			//abi.New
			blocknumber, _ := client.BlockNumber(context.Background())
			//msgData :=
			msg := ethereum.CallMsg{
				From: sender,
				To:   &address,
				Gas:  0,
				Data: []byte("1a686502"),
			}

			reserve0, err := client.CallContract(context.Background(), msg, big.NewInt(int64(blocknumber)))
			fmt.Println(reserve0)
			if err != nil {
				log.Fatalf("Low level call failed - %v", err)
			} else {
				response := proto.Response{Tbd: string(reserve0)}
				stream.Send(&response)
			}
			//blocknumber, _ := client.BlockNumber(context.Background())
			//result, _ := client.BalanceAt(context.Background(), address, big.NewInt(int64(blocknumber)))

		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen to localhost:%d - %v", *port, err)
	} else {
		log.Printf("Listening to localhost:%d", *port)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = ("server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = ("server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterDEXStreamerServer(grpcServer, &DEXStreamerServerImp{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start server - %v", err)
	}
}
