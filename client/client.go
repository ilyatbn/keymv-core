package grpc_client

import (
	"fmt"
	"log"
	"os"
	auth "github.com/ilyatbn/keymv-proto/authengine"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func Validate(server string, token string, refId string) (*auth.ValidationDataRes, error){
	//figure out a way to use a global logger instead of creating new ones each time
	logger := log.New(os.Stdout, refId+" ", log.LstdFlags|log.Lmsgprefix)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := auth.NewAuthEngineClient(conn)
	response, err := c.Validate(context.Background(), &auth.ValidationDataReq{RequestId: refId, Token: token})
	if err != nil {
		logger.Printf("Error requesting authentication:%s", err)
		return nil, fmt.Errorf("error during authentication request:%v", err)
	}
	
	return response, nil
}
