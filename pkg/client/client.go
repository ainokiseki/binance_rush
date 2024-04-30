package client

import (
	"log"

	"google.golang.org/grpc"

	"ainokiseki/binance_rush/api"
	"ainokiseki/binance_rush/pkg/server"
)

func NewGRPCClient() api.BinanceClient {
	conn, err := grpc.Dial(server.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := api.NewBinanceClient(conn)
	return c
}
