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
	"math"
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

func computePrice(sqrtPriceX96 *big.Int, denominator *big.Int, decimals0 uint8, decimals1 uint8) float32 {
	decimalRatio := big.NewInt(int64(math.Pow10(int(decimals1 - decimals0))))
	sqrtPriceX96.Mul(sqrtPriceX96, sqrtPriceX96)
	sqrtPriceX96.Div(sqrtPriceX96, denominator)
	price := float64(sqrtPriceX96.Int64()) / float64(decimalRatio.Int64())
	return float32(price)
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

	blocknumber, err := client.BlockNumber(context.TODO())
	if err != nil {
		log.Fatalf("Blocknumber could not be fetched - %v", err)
	} else {
		log.Printf("Current block number: %v", blocknumber)
	}

	pairInstance, err := uniswapV3Pair.NewUniswapV3PairAbigen(address, client)
	if err != nil {
		log.Fatalf("Pair instance could not be fetched - %v", err)
	}

	callOpts := bind.CallOpts{
		Pending:     false,
		BlockNumber: big.NewInt(int64(blocknumber)),
		Context:     context.Background(),
	}

	token0, err := pairInstance.Token0(&callOpts)
	if err != nil {
		log.Fatalf("Token0 instance could not be fetched - %v", err)
	}

	token1, err := pairInstance.Token1(&callOpts)
	if err != nil {
		log.Fatalf("Token1 instance could not be fetched - %v", err)
	}

	token0Instance, err := erc20.NewErc20Abigen(token0, client)
	token1Instance, err := erc20.NewErc20Abigen(token1, client)

	token0Name, err := token0Instance.Name(&callOpts)
	token1Name, err := token1Instance.Name(&callOpts)

	decimals0, err := token0Instance.Decimals(&callOpts)
	if err != nil {
		log.Fatalf("Token0 decimals could not be fetched - %v", err)
	}

	decimals1, err := token1Instance.Decimals(&callOpts)
	if err != nil {
		log.Fatalf("Token1 decimals could not be fetched - %v", err)
	}

	denominator := big.NewInt(1)
	x := big.NewInt(2)
	for i := 0; i < 192; i++ {
		denominator.Mul(denominator, x)
	}

	var currentSpotPrice float32
	currentSpotPrice = -1

	for {
		select {
		case <-done:
			fmt.Println("Done")
			return nil
		case <-ticker.C:
			blocknumber, _ := client.BlockNumber(context.TODO())
			callOpts := bind.CallOpts{
				Pending:     false,
				BlockNumber: big.NewInt(int64(blocknumber)),
				Context:     context.Background(),
			}
			slot0, err := pairInstance.Slot0(&callOpts)
			if err != nil {
				log.Fatalf("slot() could not be fetched - %v", err)
			}
			sqrtPriceX96 := slot0.SqrtPriceX96
			spotPrice := computePrice(sqrtPriceX96, denominator, decimals0, decimals1)
			timeStamp := time.Now().String()

			if spotPrice != currentSpotPrice {
				currentSpotPrice = spotPrice
				response := proto.Response{Token0: token0Name, Token1: token1Name, SpotPrice: spotPrice,
					Blocknumber: int32(blocknumber), TimeStamp: timeStamp}
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
