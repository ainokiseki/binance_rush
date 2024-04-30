package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"ainokiseki/binance_rush/api"
	"ainokiseki/binance_rush/pkg/trade"
)

var Port = "127.0.0.1:35555"

func Run(c *trade.BinanceClient) {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Listening on " + Port)

	s := grpc.NewServer()
	h, err := newHandler(c)
	if err != nil {
		log.Fatal("create handler fail:", err)
	}
	defer h.close()

	api.RegisterBinanceServer(s, h)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
