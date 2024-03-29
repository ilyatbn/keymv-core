package params

import (
	"log"
	"os"

	grpc_client "github.com/ilyatbn/keymv-core/client"
	params "github.com/ilyatbn/keymv-proto/core"
	"github.com/lithammer/shortuuid/v4"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type Server struct {
	params.UnimplementedParamReaderServer
}

//change that to get from someplace else
var authenticationServiceServer string = "localhost:49001"

// func ParseMetadata(ctx context.Context, logger *log.Logger)  {
// 	md, ok := metadata.FromIncomingContext(ctx)
// }

func (s *Server) GetParam(ctx context.Context, in *params.RequestParam) (*params.ResponseParam, error) {
	logger := log.New(os.Stdout, shortuuid.New()+" ", log.LstdFlags|log.Lmsgprefix)

	//Get metadata. this should contain any url queries and headers when receiving a request from the webservice proxy.
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error getting gRPC context data. Ref:"+logger.Prefix())
	} else {
		logger.Printf("%+v", md)
	}
	ip, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Error getting client IP from context. Ref:"+logger.Prefix())
	} else {
		ip := ip.Addr.String()
		logger.Printf("Source IP: %s", ip)
	}

	val, err := grpc_client.Validate(authenticationServiceServer, in.AuthToken, logger.Prefix())
	if err != nil {
		logger.Printf(err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	} else {
		if val.Valid == "true" {
			logger.Printf("Authenticated to orgId:%v", val.OrgId)
		} else {
			msg := "Error authenticating to server. Please check your token."
			return nil, status.Errorf(codes.Unauthenticated, msg)
		}
	}

	//policy := Send headers,ip and requestId to PolicyChecker? (Client Stream)

	//Send param_name and policy
	return &params.ResponseParam{RequestId: logger.Prefix(), ParamValue: in.Param, FromAppliedPolicyId: "md"}, nil
}

//THIS IS FOR GETALLPARAMETERS/GETPARAMETERSBYPATH
//FOREACH RETURNED POLICY START A GOROUTINE THAT FETCHES ALL PARAMETERS BY MATCHED POLICY from DataGenerator (Client Stream)
