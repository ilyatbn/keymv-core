package grpc_client

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/ilyatbn/keymv-proto/authengine"
	"os"
)

func Authenticate(server string, token string, refId string) (string){
	//figure out a way to use a global logger instead of creating new ones each time
	logger := log.New(os.Stdout, refId+" ", log.LstdFlags|log.Lmsgprefix)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := auth.NewAuthEngineClient(conn)
	response, err := c.Auth(context.Background(), &auth.Request{RequestId: refId, AuthToken: token})
	if err != nil {
		logger.Printf("Error requesting authentication:%s", err)
		return ""
	}
	return response.Orgid
}
