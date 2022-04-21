package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/ilyatbn/keymv-core/params"
	paramsengine "github.com/ilyatbn/keymv-proto/core"
)

func main() {
	listenPort := ":49000"
	lis, err := net.Listen("tcp4", listenPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	paramServer := params.Server{}
	grpcServer := grpc.NewServer()
	paramsengine.RegisterParamReaderServer(grpcServer, &paramServer)
	log.Printf("Listening on 0.0.0.0"+listenPort)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}