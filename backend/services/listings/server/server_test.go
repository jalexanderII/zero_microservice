package server

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	config "github.com/jalexanderII/zero_microservice"
	fileServicePB "github.com/jalexanderII/zero_microservice/gen/file_service"
	"google.golang.org/grpc"
)

var L = hclog.Default()

func MockFileServiceClient() fileServicePB.FileServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", config.FIlESERVICESERVERPORT), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return fileServicePB.NewFileServiceClient(conn)
}
