package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	erc20 "github.com/toamto94/dex-streamer.git/pkg/abigen/erc20"
	uniswapV3Pair "github.com/toamto94/dex-streamer.git/pkg/abigen/uniswapV3Pair"
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
	done := make(chan bool)
	ticker := time.NewTicker(time.Millisecond * time.Duration(contract.ScrapeInterval))
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Println("Done")
			return nil
		case <-ticker.C:
			blocknumber, _ := client.BlockNumber(context.TODO())

			pairInstance, err := uniswapV3Pair.NewUniswapV3PairAbigen(address, client)

			callOpts := bind.CallOpts{
				Pending:     false,
				BlockNumber: big.NewInt(int64(blocknumber)),
				Context:     context.Background(),
			}

			token0, err := pairInstance.Token0(&callOpts)
			token1, err := pairInstance.Token1(&callOpts)

			token0Instance, err := erc20.NewErc20Abigen(token0, client)
			token1Instance, err := erc20.NewErc20Abigen(token1, client)

			fmt.Println(token0Instance.BalanceOf(&callOpts, address))
			fmt.Println(token1Instance.BalanceOf(&callOpts, address))

			if err != nil {
				log.Fatalf("API call failed - %v", err)
			} else {
				response := proto.Response{Tbd: "a"}
				stream.Send(&response)
			}

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
