package main

import (
	"log"
	"net"
	"github.com/ilyatbn/keymv-proto/core"
	"google.golang.org/grpc"
)

func main() {
	listenPort := ":49000"
	lis, err := net.Listen("tcp4", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	paramServer := params.Server{}
	grpcServer := grpc.NewServer()
	params.RegisterParamReaderServer(grpcServer, &paramServer)
	log.Printf("Listening on 0.0.0.0"+listenPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}